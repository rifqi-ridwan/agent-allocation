package customer

import (
	"agent-allocation/domain"
	"context"
	"encoding/json"
)

type service struct {
	repo ICustomerRepository
}

type ICustomerService interface {
	CreateQueue(ctx context.Context, customer domain.Customer) error
}

func NewService(repo ICustomerRepository) ICustomerService {
	return &service{repo}
}

func (s *service) CreateQueue(ctx context.Context, customer domain.Customer) error {
	payload, err := json.Marshal(customer)
	if err != nil {
		return err
	}
	return s.repo.InsertQueue(ctx, payload)
}
