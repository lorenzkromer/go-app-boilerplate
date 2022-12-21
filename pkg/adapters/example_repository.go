package adapters

type ExampleRepository struct{}

func (r *ExampleRepository) Hello() string {
	return "Hello"
}

func NewExampleRepository() *ExampleRepository {
	return &ExampleRepository{}
}
