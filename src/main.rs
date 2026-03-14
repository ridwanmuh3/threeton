use std::{fs, path::Path, process::Command};

use clap::{Parser, Subcommand};
use include_dir::{Dir, include_dir};

static TEMPLATE_DIR: Dir = include_dir!("$CARGO_MANIFEST_DIR/src/templates");

#[derive(Parser)]
#[command(
    name = "threeton",
    version,
    about = "CLI tool for initializing Go project templates"
)]
struct Cli {
    #[command(subcommand)]
    command: Commands,
}

#[derive(Subcommand)]
enum Commands {
    /// Initialize a new Go project from template
    Init {
        /// Project directory name
        #[arg(short, long)]
        name: String,

        /// Go module path (e.g., github.com/user/project)
        #[arg(short, long)]
        module: String,
    },
}

fn main() {
    let cli = Cli::parse();

    match cli.command {
        Commands::Init { name, module } => {
            init_project(&name, &module);
        }
    }
}

fn init_project(name: &str, module: &str) {
    let project_path = Path::new(name);

    if project_path.exists() {
        eprintln!("error: directory '{}' already exists", name);
        std::process::exit(1);
    }

    println!("Creating project '{}'...", name);

    fs::create_dir_all(project_path).unwrap_or_else(|e| {
        eprintln!("error: failed to create project directory: {}", e);
        std::process::exit(1);
    });

    write_dir(&TEMPLATE_DIR, project_path, module);

    // Create empty placeholder directories
    let empty_dirs = ["api", "migrations", "pkg/util"];
    for dir in &empty_dirs {
        let dir_path = project_path.join(dir);
        fs::create_dir_all(&dir_path).unwrap_or_else(|e| {
            eprintln!("warning: failed to create directory '{}': {}", dir, e);
        });
    }

    println!("Running go mod tidy...");
    let status = Command::new("go")
        .args(["mod", "tidy"])
        .current_dir(project_path)
        .status();

    match status {
        Ok(s) if s.success() => {
            println!("Project '{}' initialized successfully!", name);
        }
        Ok(_) => {
            eprintln!("warning: go mod tidy exited with non-zero status");
            println!(
                "Project '{}' created, but go mod tidy failed. Run it manually.",
                name
            );
        }
        Err(e) => {
            eprintln!("warning: failed to run go mod tidy: {}", e);
            println!(
                "Project '{}' created, but go mod tidy could not be run.",
                name
            );
        }
    }
}

fn write_dir(dir: &Dir, base_path: &Path, module: &str) {
    for file in dir.files() {
        let file_path = base_path.join(file.path());
        if let Some(parent) = file_path.parent() {
            fs::create_dir_all(parent).unwrap_or_else(|e| {
                eprintln!("error: failed to create directory: {}", e);
                std::process::exit(1);
            });
        }

        let contents = file.contents();
        if let Ok(text) = std::str::from_utf8(contents) {
            let modified = text.replace("threeton-starter", module);
            fs::write(&file_path, modified).unwrap_or_else(|e| {
                eprintln!(
                    "error: failed to write file '{}': {}",
                    file.path().display(),
                    e
                );
                std::process::exit(1);
            });
        } else {
            fs::write(&file_path, contents).unwrap_or_else(|e| {
                eprintln!(
                    "error: failed to write file '{}': {}",
                    file.path().display(),
                    e
                );
                std::process::exit(1);
            });
        }
    }

    for subdir in dir.dirs() {
        write_dir(subdir, base_path, module);
    }
}
