package agent

import (
	"agent-allocation/domain"
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

type repository struct {
}

type IAgentRepository interface {
	GetAllAgents(ctx context.Context) ([]domain.Agent, error)
	AssignAgent(ctx context.Context, roomID string, agentID string) error
}

const (
	postMethod = "POST"
	getMethod  = "GET"
)

var (
	baseURL         = os.Getenv("QISCUS_BASE_URL")
	allAgentsPath   = os.Getenv("QISCUS_ALL_AGENT_PATH")
	assignAgentPath = os.Getenv("QISCUS_ASSIGN_AGENT_PATH")
	appID           = os.Getenv("QISCUS_APP_ID")
	secretKey       = os.Getenv("QISCUS_SECRET_KEY")
)

func NewRepository() IAgentRepository {
	return &repository{}
}

func (r *repository) GetAllAgents(ctx context.Context) ([]domain.Agent, error) {
	var err error
	var result domain.AllAgent

	request, err := http.NewRequest(getMethod, baseURL+allAgentsPath, nil)
	if err != nil {
		return result.Data.Agent.Data, err
	}
	request.Header.Set("Qiscus-App-Id", appID)
	request.Header.Set("Qiscus-Secret-Key", secretKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return result.Data.Agent.Data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return result.Data.Agent.Data, err
	}
	log.Printf("get all agents")
	return result.Data.Agent.Data, nil
}

func (r *repository) AssignAgent(ctx context.Context, roomID string, agentID string) error {
	var err error

	param := url.Values{}
	param.Set("room_id", roomID)
	param.Set("agent_id", agentID)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest(postMethod, baseURL+assignAgentPath, payload)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Qiscus-App-Id", appID)
	request.Header.Set("Qiscus-Secret-Key", secretKey)

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}
	log.Printf("assign agent: %s to room: %s", agentID, roomID)
	return nil
}
