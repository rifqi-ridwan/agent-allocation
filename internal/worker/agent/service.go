package agent

import (
	"agent-allocation/domain"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

type service struct {
	repository IAgentRepository
}

type IAgentService interface {
	AssignAgent(ctx context.Context, customer domain.QueuePayload) error
}

func NewService(repo IAgentRepository) IAgentService {
	return &service{repo}
}

func (s *service) AssignAgent(ctx context.Context, customer domain.QueuePayload) error {
	retry := 0
	isAssign := false
	for !isAssign && retry < 5 {
		agents, err := s.repository.GetAllAgents(ctx)
		if err != nil {
			log.Printf("failed to get agents: %v", err)
		} else {
			for _, agent := range agents {
				if agent.CurrentCustomerCount < 2 && agent.IsAvailable {
					agentID := strconv.Itoa(agent.ID)
					err := s.repository.AssignAgent(ctx, customer.RoomID, agentID)
					if err != nil {
						log.Printf("failed to assign agent: %s to room: %s error: %v", agentID, customer.RoomID, err)
						continue
					}
					isAssign = true
					break
				}
			}
		}
		if !isAssign {
			log.Printf("failed to assign agent, retry after 30s")
			time.Sleep(30 * time.Second)
			retry += 1
		}
	}

	if !isAssign {
		err := fmt.Sprintf("failed to assign agent to room : %s, and already retry 5 times", customer.RoomID)
		return errors.New(err)
	}
	return nil
}
