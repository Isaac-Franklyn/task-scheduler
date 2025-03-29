package api

import (
	"fmt"

	raft "github.com/Isaac-Franklyn/task-scheduler/internal/core/Raft"
	"github.com/Isaac-Franklyn/task-scheduler/internal/domain"
	"github.com/google/uuid"
)

type APIGateway struct {
	RaftService raft.RaftClusterService
}

func NewAPIGateway(cluster raft.RaftClusterService) *APIGateway {
	return &APIGateway{RaftService: cluster}
}

func (apigateway *APIGateway) ValidateTask(task *domain.Task) error {

	if task.Payload == nil {
		return fmt.Errorf("payload is empty")
	}

	if task.Type != "api_call" {
		return fmt.Errorf("wrong task type, expected: 'api_call'")
	}

	if task.Priority <= 0 || task.Priority > 10 {
		return fmt.Errorf("priority is out of bounds, allowed: 1-10")
	}

	if task.Retries < 0 || task.Retries > 5 {
		return fmt.Errorf("retry limit exceeded, allowed: 0-5")
	}

	task.Retries++ // default operation included, user input not considered default input retry number
	task.ID = uuid.New().String()
	task.Status = "Pending"

	//sending task to raft cluster.
	err := apigateway.RaftService.SendTaskToCluster(task)
	if err != nil {
		return err
	}

	return nil
}

func (apigateway *APIGateway) NewInput(task *domain.Task) error {

	err := apigateway.ValidateTask(task)
	if err != nil {
		return err
	}

	return nil
}
