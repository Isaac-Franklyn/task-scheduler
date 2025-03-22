package ports

import (
	"github.com/Isaac-Franklyn/task-scheduler/internal/domain"
)

type InputService interface {
	GetTaskFromClient() (domain.Task, error)
}
