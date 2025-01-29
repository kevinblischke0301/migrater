package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"

	"github.com/kevinblischke0301/migrater/internal/db"
	"github.com/kevinblischke0301/migrater/internal/env"
)

func main() {
	err := godotenv.Load()
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

	entries, err := os.ReadDir(env.MigrationDir)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while reading migrations directory:\n%s", err))
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			log.Fatal(fmt.Sprintf("Error while reading file informations:\n%s", err))
		}

		fileName := info.Name()
		pathName := filepath.Join(env.MigrationDir, fileName)
		file, err := os.Open(pathName)
		if err != nil {
			log.Fatal(fmt.Sprintf("Error while opening file \"%s\":\n%s", fileName, err))
		}

		fileContent, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(fmt.Sprintf("Error while reading file \"%s\":\n%s", fileName, err))
		}

		queries := strings.SplitAfter(string(fileContent), ";")
		for i := 0; i < len(queries)-1; i++ {
			_, err = db.Exec(queries[i])

			if err != nil {
				log.Fatal(fmt.Sprintf("Error while performing migration file \"%s\":\n%s", fileName, err))
			}
		}
	}

	fmt.Println("Migration completed.")
}
