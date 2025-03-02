package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"test-project/internal/config"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable",
		cfg.PostgresUser,
		os.Getenv("POSTGRES_PASSWORD"),
		cfg.PostgresDB,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %s", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Ошибка получения базы данных: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	if err := applyMigrations(sqlDB); err != nil {
		log.Fatalf("Ошибка применения миграций: %v", err)
	}

	return db, nil
}

func applyMigrations(db *sql.DB) error {
	log.Println("Начинаем миграции...")
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	migrationsDir := "./db/migrations"

	if err := goose.Up(db, migrationsDir); err != nil {
		return fmt.Errorf("Ошибка выполнения миграций: %w", err)
	}

	log.Println("Миграции успешно применены")
	return nil
}
