package repository

type TestRepository struct{}

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}

func (r *TestRepository) SayHello(name string) string {
	return "Hello " + name
}
