package adapters

import "github.com/Isaac-Franklyn/task-scheduler/internal/domain"

type RaftInput struct{}

func NewRaftInput() *RaftInput {
	return &RaftInput{}
}

func (r *RaftInput) SendTaskToRaftLeader(task *domain.Task) {

	
}
