package api

import (
	"fmt"

	ports "github.com/Isaac-Franklyn/task-scheduler/internal/application/ports"
	"github.com/google/uuid"
)

type APIGateway struct {
	Input ports.InputService
}

func NewAPIGateway(input ports.InputService) *APIGateway {
	return &APIGateway{Input: input}
}

func (apigateway *APIGateway) ValidateTask() error {

	task, err := apigateway.Input.GetTaskFromClient()
	if err != nil {
		return err
	}

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

	return nil
}
