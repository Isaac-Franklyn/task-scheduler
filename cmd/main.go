package main

import (
	"log"
	"net/http"

	"github.com/Isaac-Franklyn/task-scheduler/internal/application/adapters"
	api "github.com/Isaac-Franklyn/task-scheduler/internal/core/Api"
	raft "github.com/Isaac-Franklyn/task-scheduler/internal/core/Raft"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//start a raft cluster
	raftCluster := raft.NewRaftCluster(5)
	raftCluster.StartCluster()

	//start a api gateway
	APIGateway := api.NewAPIGateway(raftCluster)

	r.POST("/task", func(c *gin.Context) {

		newHTTPAdapter := adapters.NewHTTPInputAdapter(c, APIGateway)

		err := newHTTPAdapter.SendInputToApi()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(200, gin.H{"message": "Task validated and submitted"})
	})

	log.Println("task-scheduler running on port 8080")
	r.Run(":8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
