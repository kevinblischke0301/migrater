package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/kevinblischke0301/migrater/internal/db"
	"github.com/kevinblischke0301/migrater/internal/env"
)

func main() {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
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
	}

	db, err := db.GetDB(&env)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	entries, err := os.ReadDir(env.MigrationDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Open(fmt.Sprintf("%s/%s", env.MigrationDir, info.Name()))
		if err != nil {
			log.Fatal(err)
		}

		content, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		queries := strings.SplitAfter(string(content), ";")
		for i := 0; i < len(queries)-1; i++ {
			_, err = db.Exec(queries[i])
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Println("Migration completed.")
}
