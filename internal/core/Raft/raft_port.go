package raft

import "github.com/Isaac-Franklyn/task-scheduler/internal/domain"

type RaftClusterService interface {
	SendTaskToCluster(task *domain.Task) error
}
