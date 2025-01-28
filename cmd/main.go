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

type Env struct {
	DBType       string
	DBNetwork    string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	MigrationDir string
}

func main() {
	err := godotenv.Load()
	AbortIf(err)

	env := Env{
		DBType:       os.Getenv("DB_TYPE"),
		DBNetwork:    os.Getenv("DB_NETWORK"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		MigrationDir: os.Getenv("MIGRATION_DIR"),
	}

	db, err := GetDB(&env)
	AbortIf(err)

	err = db.Ping()
	AbortIf(err)

	entries, err := os.ReadDir(env.MigrationDir)
	AbortIf(err)

	for _, entry := range entries {
		info, err := entry.Info()
		AbortIf(err)

		file, err := os.Open(fmt.Sprintf("%s/%s", env.MigrationDir, info.Name()))
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

func GetDB(env *Env) (*sql.DB, error) {
	switch env.DBType {

	case "mysql":
		cfg := mysql.NewConfig()
		cfg.Net = env.DBNetwork
		cfg.Addr = fmt.Sprintf("%s:%s", env.DBHost, env.DBPort)
		cfg.User = env.DBUser
		cfg.Passwd = env.DBPassword

		db, err := sql.Open("mysql", cfg.FormatDSN())

		return db, err

	default:
		return nil, errors.New(fmt.Sprintf("%s isn't a supported database type", env.DBType))
	}
}

func AbortIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
