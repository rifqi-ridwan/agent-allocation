package agent

import (
	"agent-allocation/domain"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

type handler struct {
	service IAgentService
}

func NewHandler(service IAgentService) *handler {
	return &handler{service}
}

func (h *handler) HandleAssignAgentTask(ctx context.Context, task *asynq.Task) error {
	var payload domain.Customer
	err := json.Unmarshal(task.Payload(), &payload)
	if err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	err = h.service.AssignAgent(ctx, payload)
	if err != nil {
		return fmt.Errorf("assign agent failed: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("assign agent success for room_id: %s", payload.RoomID)
	return nil
}
