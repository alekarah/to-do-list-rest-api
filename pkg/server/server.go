package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Config содержит конфигурацию сервера
type Config struct {
	Port   string
	WebDir string
}

// Server представляет наш веб-сервер
type Server struct {
	config *Config
	mux    *http.ServeMux
}

// NewServer создает новый экземпляр сервера
func NewServer(config *Config) *Server {
	return &Server{
		config: config,
		mux:    http.NewServeMux(),
	}
}

// SetupRoutes настраивает маршруты сервера
func (s *Server) SetupRoutes() {
	// Проверяем существование директории с веб-файлами
	if _, err := os.Stat(s.config.WebDir); os.IsNotExist(err) {
		log.Fatalf("Директория %s не найдена", s.config.WebDir)
	}

	// Настраиваем файл-сервер для обслуживания статических файлов
	fileServer := http.FileServer(http.Dir(s.config.WebDir))
	s.mux.Handle("/", fileServer)
}

// Start запускает сервер
func (s *Server) Start() error {
	addr := ":" + s.config.Port

	// Проверяем, откуда взят порт
	if os.Getenv("TODO_PORT") != "" {
		fmt.Printf("Сервер запущен на http://localhost:%s (порт из переменной окружения TODO_PORT)\n", s.config.Port)
	} else {
		fmt.Printf("Сервер запущен на http://localhost:%s (порт по умолчанию)\n", s.config.Port)
	}
	fmt.Printf("Обслуживаем файлы из директории: %s\n", s.config.WebDir)

	return http.ListenAndServe(addr, s.mux)
}

// DefaultConfig возвращает конфигурацию по умолчанию
func DefaultConfig() *Config {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	return &Config{
		Port:   port,
		WebDir: "./web",
	}
}
