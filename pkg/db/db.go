package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

// db хранит глобальный указатель на базу данных
var db *sql.DB

// schema содержит SQL команды для создания таблицы и индекса
const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(256) NOT NULL DEFAULT "",
    comment TEXT NOT NULL DEFAULT "",
    repeat VARCHAR(128) NOT NULL DEFAULT ""
);

CREATE INDEX idx_scheduler_date ON scheduler(date);
`

// Init инициализирует базу данных
func Init(dbFile string) error {
	// Проверяем существование файла базы данных
	_, err := os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	// Открываем базу данных
	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	// Проверяем соединение
	if err = db.Ping(); err != nil {
		return err
	}

	// Если файл не существовал, создаем таблицу и индекс
	if install {
		_, err = db.Exec(schema)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetDB возвращает указатель на базу данных
func GetDB() *sql.DB {
	return db
}

// Close закрывает соединение с базой данных
func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
