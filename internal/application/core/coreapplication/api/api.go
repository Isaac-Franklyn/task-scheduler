package api

import (
	"fmt"

	"github.com/Isaac-Franklyn/task-scheduler/internal/application/core/coreadapters"
	leadercluster "github.com/Isaac-Franklyn/task-scheduler/internal/application/core/coreapplication/LeaderCluster"
	"github.com/Isaac-Franklyn/task-scheduler/internal/application/core/coreports"
	"github.com/Isaac-Franklyn/task-scheduler/internal/application/ports"
	"github.com/Isaac-Franklyn/task-scheduler/internal/domain"
	"github.com/google/uuid"
)

type APIGateway struct {
	Raftcluster *leadercluster.RaftCluster
}

func NewAPIGateway(cluster *leadercluster.RaftCluster) *APIGateway {
	return &APIGateway{Raftcluster: cluster}
}

func (apigateway *APIGateway) ValidateTask(task domain.Task) error {

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
	var clusterPort coreports.ClusterInputPort = coreadapters.NewClusterInputAdapter(apigateway.Raftcluster)
	err := clusterPort.SendTaskToCluster(&task)
	if err != nil {
		return fmt.Errorf("error sending task to the raft cluster: %v", err)
	}

	return nil
}

func (apigateway *APIGateway) NewInput(Input ports.InputService) error {

	task, err := Input.GetTaskFromClient()
	if err != nil {
		return err
	}

	err = apigateway.ValidateTask(*task)
	if err != nil {
		return err
	}

	return nil
}
