package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"agent-allocation/domain"
	"agent-allocation/internal/worker/agent"
	_ "agent-allocation/util"
	"agent-allocation/util/db"

	"github.com/bgentry/que-go"
)

func main() {
	pgxpool, err := db.Setup(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to create database connection: %v", err)
	}
	defer pgxpool.Close()

	agentRepo := agent.NewRepository()
	agentService := agent.NewService(agentRepo)
	agentHandler := agent.NewHandler(agentService)

	wm := que.WorkMap{
		domain.TypeAllocateAgent: agentHandler.HandleAssignAgentTask,
	}

	client := que.NewClient(pgxpool)
	workers := que.NewWorkerPool(client, wm, 1)
	wakeUpInterval := 2 * time.Second
	workers.Interval = wakeUpInterval

	// Catch signal so we can shutdown gracefully
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	go workers.Start()

	// Wait for a signal
	sig := <-sigCh
	log.Printf("singal: %s. signal received. shutting down.", sig)

	workers.Shutdown()
}
