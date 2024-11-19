package todo

import (
	"context"
	"errors"
	"example/go-blog-api/internal/db"
	"fmt"
	"strings"
)

type Item struct {
	Task   string
	Status string
}

type Service struct {
	db *db.DB
}

func NewService(db *db.DB) *Service {
	return &Service{
		db: db,
	}
}

func (svc *Service) Add(todo string) error {

	items, err := svc.GetAll()
	if err != nil {
		return fmt.Errorf("failed to read from db: %w", err)
	}

	for _, s := range items {
		if s.Task == todo {
			return errors.New("todo already exists")
		}
	}

	if err := svc.db.InsertItem(context.Background(), db.Item{
		Task:   todo,
		Status: "To_be_started",
	}); err != nil {
		return fmt.Errorf("failed to insert item: %w", err)
	}
	return nil 
}

func (svc *Service) GetAll() ([]Item, error) {
	var results []Item
	items, err := svc.db.GetAllItems(context.Background())
	if err != nil {
		return nil, fmt.Errorf("faild to read from db: %w", err)
	}

	for _, item := range items {
		results = append(results, Item{
			Task:   item.Task,
			Status: item.Status,
		})
	}
	return results, nil
}

func (svc *Service) Search(query string) []string {
	var results []string
	for _, todo := range svc.todos {
		if strings.Contains(todo.Task, query) {
			results = append(results, todo.Task)
		}
	}
	return results
}
