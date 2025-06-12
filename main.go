package main

import (
	"log"
	"todo/pkg/db"
	"todo/pkg/server"
)

func main() {
	// Инициализируем базу данных
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}
	defer db.Close()

	log.Println("База данных успешно инициализирована")

	// Создаем конфигурацию сервера
	config := server.DefaultConfig()

	// Создаем новый сервер
	srv := server.NewServer(config)

	// Настраиваем маршруты
	srv.SetupRoutes()

	// Запускаем сервер
	log.Fatal(srv.Start())
}
