package domain

import "github.com/hashicorp/raft"

type Task struct {
	ID       string `json:"id"`
	Payload  any    `json:"payload"`
	Type     string `json:"type"`
	Priority int    `json:"priority"`
	Status   string `json:"status"`
	Retries  int    `json:"retries"`
}

type Node struct {
	ID   string
	Raft *raft.Raft
}
