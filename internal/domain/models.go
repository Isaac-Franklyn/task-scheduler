package domain

type Task struct {
	ID       string `json:"id"`
	Payload  any    `json:"payload"`
	Type     string `json:"type"`
	Priority int    `json:"priority"`
	Status   string `json:"status"`
	Retries  int    `json:"retries"`
}
