package leadercluster

import (
	"fmt"
	"log"
	"time"

	"github.com/Isaac-Franklyn/task-scheduler/internal/domain"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
)

type RaftCluster struct {
	N       int // number of nodes to create
	Cluster []*domain.Node
}

func NewRaftCluster(n int) *RaftCluster {
	return &RaftCluster{N: n}
}

func (raftcluster *RaftCluster) StartCluster() {

	n := raftcluster.N

	for i := 0; i < n; i++ {

		id := fmt.Sprintf("node-%d", i+1)
		addr := fmt.Sprintf("127.0.0.1.%d", 9000+i)
		node := createRaftNode(id, addr)
		raftcluster.Cluster = append(raftcluster.Cluster, node)
	}
}

func createRaftNode(id, addr string) *domain.Node {
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(id)

	store, err := raftboltdb.NewBoltStore(fmt.Sprintf("%s-log.bolt", id))
	if err != nil {
		log.Fatalf("Error creating store: %v", err)
	}

	transport, err := raft.NewTCPTransport(addr, nil, 3, time.Second, nil)
	if err != nil {
		log.Fatalf("Error creating transport: %v", err)
	}

	snapshots, err := raft.NewFileSnapshotStore(".", 1, nil)
	if err != nil {
		log.Fatalf("Error creating snapshot store: %v", err)
	}

	node := &domain.Node{ID: id}

	raftNode, err := raft.NewRaft(config, nil, store, store, snapshots, transport)
	if err != nil {
		log.Fatalf("Error starting Raft: %v", err)
	}

	node.Raft = raftNode
	return node
}

func (raftcluster *RaftCluster) GetLeader() (*domain.Node, error) {

	for _, node := range raftcluster.Cluster {
		if node.Raft.State() == raft.Leader {
			return node, nil
		}
	}

	return &domain.Node{}, fmt.Errorf("no leader available")
}
