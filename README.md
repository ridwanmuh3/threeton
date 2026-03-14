```
         _____ _                   _
        |_   _| |                 | |
          | | | |__  _ __ ___  ___| |_ ___  _ __
          | | | '_ \| '__/ _ \/ _ \ __/ _ \| '_ \
          | | | | | | | |  __/  __/ || (_) | | | |
          \_/ |_| |_|_|  \___|\___|\__\___/|_| |_|
                    Develop by ridwanmuh3
```

A Rust-based CLI tool for scaffolding Go projects with a clean architecture template.

## Installation

```bash
cargo install --path .
```

## Usage

```bash
threeton init --name my-app --module github.com/username/my-app
```

### Options

| Flag       | Short | Description                                      |
| ---------- | ----- | ------------------------------------------------ |
| `--name`   | `-n`  | Project directory name                           |
| `--module` | `-m`  | Go module path (e.g., `github.com/user/project`) |

## Generated Folder Structure

```
my-app/
├── cmd/
│   ├── app/main.go          # Application entrypoint
│   └── seed/seeder.go       # Database seeder
├── internal/
│   ├── config/              # App bootstrap, Fiber, GORM, Viper, Zap, Validator
│   ├── delivery/http/
│   │   ├── handler/         # HTTP handlers
│   │   └── route/           # Route definitions
│   ├── exception/           # Error handling
│   ├── model/               # Request/response models
│   ├── repository/          # Data access layer
│   └── service/             # Business logic
├── test/                    # Tests
├── api/                     # API specs (placeholder)
├── migrations/              # DB migrations (placeholder)
├── pkg/util/                # Shared utilities (placeholder)
├── .env                     # Environment variables
├── Dockerfile               # Multi-stage Docker build
├── docker-compose.yml       # PostgreSQL
├── Makefile                 # Dev/build/compose commands
├── go.mod
└── go.sum
```

## Template Stack

- **Framework**: [Fiber v3](https://github.com/gofiber/fiber)
- **ORM**: [GORM](https://gorm.io/) with PostgreSQL
- **Config**: [Viper](https://github.com/spf13/viper)
- **Logging**: [Zap](https://github.com/uber-go/zap)
- **Validation**: [go-playground/validator](https://github.com/go-playground/validator)
- **JSON**: [Sonic](https://github.com/bytedance/sonic)

## Future Works

- Add capabilities for choosing different template framework (e.g: echo, gin, etc)
- Add capabilities for choosing different databases (e.g: MySQL, SQLite, MongoDB)

## License

[MIT](https://choosealicense.com/licenses/mit/)
