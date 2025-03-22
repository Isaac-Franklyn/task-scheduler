package adapters

import (
	"net/http"

	"github.com/Isaac-Franklyn/task-scheduler/internal/domain"
	"github.com/gin-gonic/gin"
)

type HTTPInputAdapter struct {
	Ctx *gin.Context
}

func NewHTTPInputAdapter(c *gin.Context) *HTTPInputAdapter {
	return &HTTPInputAdapter{Ctx: c}
}

func (Input *HTTPInputAdapter) GetTaskFromClient() (domain.Task, error) {

	var task domain.Task

	if err := Input.Ctx.ShouldBindJSON(&task); err != nil {
		Input.Ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return domain.Task{}, nil
	}

	return task, nil
}
