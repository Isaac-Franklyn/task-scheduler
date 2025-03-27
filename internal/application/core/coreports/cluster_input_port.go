package coreports

import (
	"github.com/Isaac-Franklyn/task-scheduler/internal/domain"
)

type ClusterInputPort interface {
	SendTaskToCluster(task *domain.Task) error
}
