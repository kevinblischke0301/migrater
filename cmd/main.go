package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	Abort(err)

	cfg := mysql.NewConfig()
	cfg.Net = GetEnv("DB_NETWORK")
	cfg.Addr = fmt.Sprintf("%s:%s", GetEnv("DB_HOST"), GetEnv("DB_PORT"))
	cfg.User = GetEnv("DB_USER")
	cfg.Passwd = GetEnv("DB_PASSWORD")

	db, err := sql.Open("mysql", cfg.FormatDSN())
	Abort(err)

	err = db.Ping()
	Abort(err)

	entries, err := os.ReadDir(GetEnv("MIGRATION_DIR"))
	Abort(err)

	for _, entry := range entries {
		info, err := entry.Info()
		Abort(err)

		file, err := os.Open(fmt.Sprintf("%s%s", GetEnv("MIGRATION_DIR"), info.Name()))
		Abort(err)

		content, err := io.ReadAll(file)
		Abort(err)

		_, err = db.Exec(string(content))
		Abort(err)
	}

	fmt.Println("Migration completed.")
}

func GetEnv(key string) string {
	env, ok := os.LookupEnv(key)

	if !ok {
		log.Fatal(fmt.Sprintf("No environment variable \"%s\" set.", key))
	}

	return env
}

func Abort(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
