package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"shortlink/internals/domain/service"
	"syscall"

	"shortlink/internals/presentation"
	"shortlink/internals/repository"
)

func main() {
	dbUrl := os.Getenv("DB_URL")
	dbPool, err := repository.NewPostgresDbPool(dbUrl)
	if err != nil {
		log.Fatal("Cant init connection to postgres db: %s", err)
	}

	repository := repository.NewPostgresRepository(dbPool)
	service := service.NewLinkService(repository)
	handler := presentation.NewHandler(service)

	presentation.AddRoutes(handler)
	server := new(presentation.Server)

	err = server.LaunchServer(handler)
	if err != nil {
		log.Fatal("Cant launch server: %s", err.Error())
	}
	log.Printf("Server started successfuly.")

	waitForExit()

	dbPool.Close()

	err = server.GracefulShutdown(context.Background())
	if err != nil {
		log.Fatal("Get error while server shutting down: %s", err.Error())
	}
}

func waitForExit() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit
}
