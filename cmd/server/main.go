package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	product "test-project"
	"test-project/internal/config"
	repository "test-project/internal/database/postgres"
	handler "test-project/internal/handlers"
	"time"
)

func main() {
	cfg := config.Load()

	db, err := repository.Connect(cfg)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных", err)
		os.Exit(1)
	}

	repo := repository.NewRepository(db)

	handlerInstance := handler.NewHandler(repo)
	router := handlerInstance.InitRoutes()
	srv := new(product.Server)
	go func() {
		if err := srv.Run("8080", router); err != nil {
			log.Fatalf("Ошибка запуска сервера: %s", err)
		}
	}()

	log.Println("Сервер запущен на порту 8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	log.Println("Завершение работы сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Сервер был вынужден выключиться: %s", err)
	}
}
