package services

import "github.com/lorenzkromer/go-app-boilerplate/pkg/ports"

type ExampleService struct {
	exampleRepository ports.ExampleInterface
}

func (s *ExampleService) Hello() string {
	return s.exampleRepository.Hello()
}

func NewExampleService(exampleRepo ports.ExampleInterface) *ExampleService {
	return &ExampleService{
		exampleRepository: exampleRepo,
	}
}
