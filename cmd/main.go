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
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}


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
	if err != nil {
		log.Fatal(err)
	}

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

		queries := strings.Split(string(content), ";")
		for i := 0; i < len(queries)-1; i++ {
			_, err = db.Exec(queries[i] + ";")
			if err != nil {
				log.Fatal(err)
			}
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
