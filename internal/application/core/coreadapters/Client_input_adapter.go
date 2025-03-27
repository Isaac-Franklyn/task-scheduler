package coreadapters

import (
	"encoding/json"
	"fmt"
	"time"

	leadercluster "github.com/Isaac-Franklyn/task-scheduler/internal/application/core/coreapplication/LeaderCluster"
	"github.com/Isaac-Franklyn/task-scheduler/internal/domain"
)

type ClusterInputAdapter struct {
	Raftcluster *leadercluster.RaftCluster
}

func NewClusterInputAdapter(raftCluster *leadercluster.RaftCluster) *ClusterInputAdapter {
	return &ClusterInputAdapter{Raftcluster: raftCluster}
}

func (clusterinput *ClusterInputAdapter) SendTaskToCluster(task *domain.Task) error {

	node, err := clusterinput.Raftcluster.GetLeader()
	if err != nil {
		return fmt.Errorf("failed to get a leader in the cluster: %v", err)
	}

	taskBytes, err := json.Marshal(task)
	if err != nil {
		return fmt.Errorf("failed to marshal task: %v", err)
	}

	future := node.Raft.Apply(taskBytes, time.Second*2)

	if err := future.Error(); err != nil {
		return fmt.Errorf("failed to apply task to raft log: %v", err)
	}

	fmt.Println("task successfully sent to leader: ", task.ID)
	return nil
}
