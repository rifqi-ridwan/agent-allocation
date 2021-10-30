package agent

import (
	"agent-allocation/domain"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/bgentry/que-go"
)

type handler struct {
	service IAgentService
}

func NewHandler(service IAgentService) *handler {
	return &handler{service}
}

func (h *handler) HandleAssignAgentTask(job *que.Job) error {
	var payload domain.QueuePayload
	err := json.Unmarshal(job.Args, &payload)
	if err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v", err)
	}

	ctx := context.TODO()
	err = h.service.AssignAgent(ctx, payload)
	if err != nil {
		return fmt.Errorf("assign agent failed: %v", err)
	}

	log.Printf("assign agent success for room_id: %s", payload.RoomID)
	return nil
}
