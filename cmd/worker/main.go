package main

import (
	"fmt"
	"log"
	"os"

	"agent-allocation/domain"
	"agent-allocation/internal/worker/agent"
	_ "agent-allocation/util"

	"github.com/hibiken/asynq"
)

func main() {
	redisAddress := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddress},
		asynq.Config{
			Concurrency: 1,
		},
	)

	agentRepo := agent.NewRepository()
	agentService := agent.NewService(agentRepo)
	agentHandler := agent.NewHandler(agentService)

	mux := asynq.NewServeMux()
	mux.HandleFunc(domain.TypeAllocateAgent, agentHandler.HandleAssignAgentTask)
	err := srv.Run(mux)
	if err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
