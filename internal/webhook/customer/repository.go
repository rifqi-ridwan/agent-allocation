package customer

import (
	"agent-allocation/domain"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

type repository struct {
	Client *asynq.Client
}

type ICustomerRepository interface {
	InsertQueue(ctx context.Context, payload []byte) error
}

func NewRepository(c *asynq.Client) ICustomerRepository {
	return &repository{c}
}

func (r *repository) InsertQueue(ctx context.Context, payload []byte) error {
	task := asynq.NewTask(domain.TypeAllocateAgent, payload)
	info, err := r.Client.Enqueue(task)
	if err != nil {
		err := fmt.Sprintf("could not enqueue task: %v", err)
		return errors.New(err)
	}
	log.Printf("enqueued task: id=%s queue%s", info.ID, info.Queue)
	return nil
}
