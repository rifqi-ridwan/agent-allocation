package customer

import (
	"agent-allocation/domain"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/bgentry/que-go"
)

type repository struct {
	Client *que.Client
}

type ICustomerRepository interface {
	InsertQueue(ctx context.Context, payload []byte) error
}

func NewRepository(c *que.Client) ICustomerRepository {
	return &repository{c}
}

func (r *repository) InsertQueue(ctx context.Context, payload []byte) error {
	job := que.Job{
		Type: domain.TypeAllocateAgent,
		Args: payload,
	}
	err := r.Client.Enqueue(&job)
	if err != nil {
		err := fmt.Sprintf("could not enqueue task: %v", err)
		return errors.New(err)
	}
	log.Printf("successfully enqueue job, id: %d", job.ID)
	return nil
}
