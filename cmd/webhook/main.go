package main

import (
	"agent-allocation/internal/webhook/customer"
	"fmt"
	"log"
	"os"

	_ "agent-allocation/util"

	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	redisAddress := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddress})
	defer client.Close()

	custRepo := customer.NewRepository(client)
	custService := customer.NewService(custRepo)
	custHandler := customer.NewAPI(custService)

	e.POST("/customagentallocation", custHandler.CreateQueue)

	serverAddr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal(e.Start(serverAddr))
}
