package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/kevinblischke0301/migrater/internal/env"
)

func GetDB(env *env.Env) (*sql.DB, error) {
	switch env.DBType {

	case "mysql":
		cfg := mysql.NewConfig()
		cfg.Net = env.DBNetwork
		cfg.Addr = fmt.Sprintf("%s:%s", env.DBHost, env.DBPort)
		cfg.User = env.DBUser
		cfg.Passwd = env.DBPassword

		db, err := sql.Open("mysql", cfg.FormatDSN())

		return db, err

	case "sqlite":
		db, err := sql.Open("sqlite3", env.DBDatabase)

		return db, err

	default:
		return nil, errors.New(fmt.Sprintf("%s isn't a supported database type", env.DBType))
	}
}
