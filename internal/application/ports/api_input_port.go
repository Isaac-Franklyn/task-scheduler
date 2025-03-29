package ports

import "github.com/Isaac-Franklyn/task-scheduler/internal/domain"

type ApiService interface {
	NewInput(task *domain.Task) error
	ValidateTask(task *domain.Task) error
}
