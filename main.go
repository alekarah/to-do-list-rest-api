package main

import (
	"log"
	"to-do-list-rest-api/pkg/server"
)

func main() {
	// Создаем конфигурацию сервера
	config := server.DefaultConfig()

	// Создаем новый сервер
	srv := server.NewServer(config)

	// Настраиваем маршруты
	srv.SetupRoutes()

	// Запускаем сервер
	log.Fatal(srv.Start())
}
