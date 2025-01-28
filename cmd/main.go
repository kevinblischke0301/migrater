package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	AbortIf(err)

	db, err := GetDB(GetEnv("DB_TYPE"))
	AbortIf(err)

	err = db.Ping()
	AbortIf(err)

	entries, err := os.ReadDir(GetEnv("MIGRATION_DIR"))
	AbortIf(err)

	for _, entry := range entries {
		info, err := entry.Info()
		AbortIf(err)

		file, err := os.Open(fmt.Sprintf("%s%s", GetEnv("MIGRATION_DIR"), info.Name()))
		AbortIf(err)

		content, err := io.ReadAll(file)
		AbortIf(err)

		queries := strings.Split(string(content), ";")
		for i := 0; i < len(queries)-1; i++ {
			_, err = db.Exec(queries[i] + ";")
			AbortIf(err)
		}
	}

	fmt.Println("Migration completed.")
}

func GetDB(dbType string) (*sql.DB, error) {
	switch dbType {

	case "mysql":
		cfg := mysql.NewConfig()
		cfg.Net = GetEnv("DB_NETWORK")
		cfg.Addr = fmt.Sprintf("%s:%s", GetEnv("DB_HOST"), GetEnv("DB_PORT"))
		cfg.User = GetEnv("DB_USER")
		cfg.Passwd = GetEnv("DB_PASSWORD")

		db, err := sql.Open("mysql", cfg.FormatDSN())

		return db, err

	default:
		return nil, errors.New(fmt.Sprintf("%s isn't a supported database type", dbType))
	}
}

func GetEnv(key string) string {
	env, ok := os.LookupEnv(key)

	if !ok {
		log.Fatal(fmt.Sprintf("No environment variable \"%s\" set.", key))
	}

	return env
}

func AbortIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
