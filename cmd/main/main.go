package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/kevinblischke0301/migrater/internal/db"
	"github.com/kevinblischke0301/migrater/internal/env"
	"github.com/kevinblischke0301/migrater/internal/service"
	"github.com/kevinblischke0301/migrater/internal/arg"
)

func main() {
	command, err := arg.ParseArg(os.Args)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while parsing command-line arguments:\n%s", err))
	}

	err = godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(fmt.Sprintf("Error while reading \".env\" file:\n%s", err))
	}

	env := env.Env{
		DBType:       os.Getenv("DB_TYPE"),
		DBNetwork:    os.Getenv("DB_NETWORK"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBDatabase:   os.Getenv("DB_DATABASE"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		MigrationDir: os.Getenv("MIGRATION_DIR"),
		RollbackDir: os.Getenv("ROLLBACK_DIR"),
	}

	db, err := db.GetDB(&env)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while opening database:\n%s", err))
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while connecting to database:\n%s", err))
	}

	switch command {
	case arg.MIGRATE:
		service.Migrate(env.MigrationDir, db)
		fmt.Println("Migration completed.")
	case arg.ROLLBACK:
		service.Rollback(env.RollbackDir, db)
		fmt.Println("Rollback completed.")
	}
}
