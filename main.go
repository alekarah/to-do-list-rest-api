package main

import (
	"log"
	"os"
	"todo/pkg/db"
	"todo/pkg/server"
)

func main() {
	// Получаем путь к базе данных из переменной окружения
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		// Если переменная окружения не задана, используем значение по умолчанию
		dbFile = "scheduler.db"
	}

	// Инициализируем базу данных
	err := db.Init(dbFile)
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}
	defer db.Close()

	log.Printf("База данных успешно инициализирована: %s", dbFile)

	// Создаем конфигурацию сервера
	config := server.DefaultConfig()

	// Создаем новый сервер
	srv := server.NewServer(config)

	// Настраиваем маршруты
	srv.SetupRoutes()

	// Запускаем сервер
	log.Fatal(srv.Start())
}
