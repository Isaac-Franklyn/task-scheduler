package main

import (
	"log"
	"net/http"

	"github.com/Isaac-Franklyn/task-scheduler/internal/application/adapters"
	"github.com/Isaac-Franklyn/task-scheduler/internal/application/core/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/task", func(c *gin.Context) {

		newHTTPAdapter := adapters.NewHTTPInputAdapter(c)
		newAPIgateway := api.NewAPIGateway(newHTTPAdapter)

		if err := newAPIgateway.ValidateTask(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(), // or err.Error() for the exact string
			})
			return
		}

		c.JSON(200, gin.H{"message": "Task validated and submitted"})
	})

	log.Println("task-scheduler running on port 8080")
	r.Run(":8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
