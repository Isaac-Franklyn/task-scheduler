package adapters

import (
	"fmt"
	"net/http"

	"github.com/Isaac-Franklyn/task-scheduler/internal/application/ports"
	"github.com/Isaac-Franklyn/task-scheduler/internal/domain"
	"github.com/gin-gonic/gin"
)

type HTTPInputAdapter struct {
	Ctx        *gin.Context
	ApiService ports.ApiService
}

func NewHTTPInputAdapter(c *gin.Context, api ports.ApiService) *HTTPInputAdapter {
	return &HTTPInputAdapter{Ctx: c, ApiService: api}
}

func (Input *HTTPInputAdapter) SendInputToApi() error {

	var task *domain.Task

	if err := Input.Ctx.ShouldBindJSON(&task); err != nil {
		Input.Ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return nil
	}

	err := Input.ApiService.NewInput(task)
	if err != nil {
		return fmt.Errorf("error sending task to api: %v", err)
	}

	return nil
}
