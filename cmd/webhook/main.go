package main

import (
	"agent-allocation/internal/webhook/customer"
	"agent-allocation/util/db"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bgentry/que-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	pgxpool, err := db.Setup(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to create database connection: %v", err)
	}
	defer pgxpool.Close()

	client := que.NewClient(pgxpool)

	custRepo := customer.NewRepository(client)
	custService := customer.NewService(custRepo)
	custHandler := customer.NewAPI(custService)

	e.GET("/healthz", healthzHandler)
	e.POST("/customagentallocation", custHandler.CreateQueue)

	serverAddr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal(e.Start(serverAddr))
}

func healthzHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message: ": "ok"})
}
