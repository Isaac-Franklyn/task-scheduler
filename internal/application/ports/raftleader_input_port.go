package ports

import "github.com/Isaac-Franklyn/task-scheduler/internal/domain"

type TaskService interface {
	SendTaskToRaftLeader(task *domain.Task) error
}
