package todo

import "errors"

type Service struct {
	todos []string
}

func NewService() *Service {
	return &Service{todos: make([]string, 0),}
}


func (svc *Service) Add(todo string) error {

	for _, s := range svc.todos{
		if s == todo{
			return errors.New("todo already exists")
		}
	}
	svc.todos = append(svc.todos, todo)
	return nil
}


func (svc *Service) GetAll() ([]string) {
return svc.todos
}