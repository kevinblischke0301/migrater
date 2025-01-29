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

		queries := strings.SplitAfter(string(fileContent), ";")
		for i := 0; i < len(queries)-1; i++ {
			_, err = db.Exec(queries[i])

			if err != nil {
				log.Fatal(fmt.Sprintf(
					"Error while performing migration file \"%s\":\n%s", fileName, err,
				))
			}
		}
	}
}
