package service

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Migrate(dir string, db *sql.DB) {
	queryFiles := readQueries(dir)

	for i, queries := range queryFiles {
		for _, query := range queries {
			_, err := db.Exec(query)

			if err != nil {
				log.Fatal(fmt.Sprintf(
					"Error while performing migration file %d:\n%s", i, err,
				))
			}
		}
	}
}

func Rollback(dir string, db *sql.DB) {
	queryFiles := readQueries(dir)
	queryFilesLen := len(queryFiles)

	for i := queryFilesLen - 1; i >= 0; i-- {
		for _, query := range queryFiles[i] {
			_, err := db.Exec(query)

			if err != nil {
				log.Fatal(fmt.Sprintf(
					"Error while performing rollback file %d:\n%s", queryFilesLen - i, err,
				))
			}
		}
	}
}

func readQueries(dir string) [][]string {
	queries := make([][]string, 0)

	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while reading migrations directory:\n%s", err))
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			log.Fatal(fmt.Sprintf("Error while reading file informations:\n%s", err))
		}

		fileName := info.Name()
		pathName := filepath.Join(dir, fileName)
		file, err := os.Open(pathName)
		if err != nil {
			log.Fatal(fmt.Sprintf("Error while opening file \"%s\":\n%s", fileName, err))
		}

		fileContent, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(fmt.Sprintf("Error while reading file \"%s\":\n%s", fileName, err))
		}

		unfilteredQueries := strings.SplitAfter(string(fileContent), ";")
		filteredQueries := make([]string, 0)
		for i := 0; i < len(unfilteredQueries)-1; i++ {
			filteredQueries = append(filteredQueries, unfilteredQueries[i])
		}

		queries = append(queries, filteredQueries)
	}

	return queries
}
