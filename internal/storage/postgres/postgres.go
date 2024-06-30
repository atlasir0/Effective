
package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"

	_ "github.com/lib/pq"
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

func createDatabase(config *Config) error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)", config.Database.Name).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		log.Println("Database already exists")
		return nil
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", config.Database.Name))
	if err != nil {
		return err
	}

	log.Println("Database created successfully")
	return nil
}

func runMigrations(config *Config) error {
	migrateCmd := exec.Command("migrate", "-path", "migrations", "-database", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name), "up")
	migrateCmd.Stdout = log.Writer()
	migrateCmd.Stderr = log.Writer()
	return migrateCmd.Run()
}

func InitDB() (*sql.DB, *Config, error) {
	cfg, err := ReadConfig("config/local.yaml")
	if err != nil {
		return nil, nil, fmt.Errorf("could not read config: %w", err)
	}

	err = createDatabase(cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("could not create database: %w", err)
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

func dropTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS tasks;")
	return err
}

func dropDatabase(config *Config) error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", config.Database.Name))
	return err
}

func CloseDB(db *sql.DB, config *Config) {
	err := dropTable(db)
	if err != nil {
		log.Fatalf("could not drop table: %v", err)
	}

	db.Close()

	err = dropDatabase(config)
	if err != nil {
		log.Fatalf("could not drop database: %v", err)
	}

	log.Println("Database closed successfully")
}


