package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func ReadConfig(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func runMigrations(config *Config) error {
	sourceURL := "file://migrations"
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	m, err := migrate.New(sourceURL, dbURL)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	log.Println("Migrations applied successfully")
	return nil
}

func InitDB() (*sql.DB, *Config, error) {
	cfg, err := ReadConfig("config/local.yaml")
	if err != nil {
		return nil, nil, fmt.Errorf("could not read config: %w", err)
	}

	dbConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open database: %w", err)
	}

	err = runMigrations(cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("could not run migrations: %w", err)
	}

	log.Println("Database initialized successfully")
	return db, cfg, nil
}

func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("could not close database: %v", err)
	}
	log.Println("Database closed successfully")
}
