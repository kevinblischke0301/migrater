# Migrater

Migrater is a leightweight command-line tool to perform migrations in the development process from sql files.

## Usage

Migrater executes the contents of every file of a specified directory in alphanumerical order of the file names on a specified database. Therefore it is recommended to only place sql files in that directory and to give every file name an ordering prefix. There should also exist statements to rollback migrations when the same migration is applied multiple times, because Migrater tries to execute all sql files again, regardless of wether they were executed and have taken effect already.

Examples for the usage of sql files for migrations can be found in the directories `example_migrations/migrations_mysql` and `example_migrations/migrations_sqlite`. Thereby the directories are reflecting according example migrations for MySQL and SQLite.

The following parameters must be set as environment variables or inside the file `.env` in the root folder:

**MySQL**:

- `DB_TYPE`: `"mysql"` for MySQL
- `DB_NETWORK`: the used network (e. g. `"tcp"`)
- `DB_HOST`: the host on which MySQL runs
- `DB_PORT`: the port on which MySQL listens
- `DB_DATABASE`: the name of the database to access
- `DB_USER`: the username of the user
- `DB_PASSWORD`: the password of the user
- `MIGRATION_DIR`: the path to the directory of the sql files

**SQLite**:

- `DB_TYPE`: `"sqlite"` for SQLite
- `DB_DATABASE`: the path to the database file to access
- `MIGRATION_DIR`: the path to the directory of the sql files

## Build

Migrater must be build locally.

### Building with Go CLI

Migrater can be build with the Go CLI. This requires Go version 1.23.5 or higher.

To build Migrater with the Go CLI, run the following command in the root folder of the project:

```bash
go build -o ./bin/migrater.exe ./cmd/main/main.go
```

The executable Migrater file can then be found in the directory `bin` with the name `migrater.exe`.

To only run Migrater without building an executable file with the Go CLI, run the following command in the root folder of the project:

```bash
go run ./cmd/main/main.go
```

### Building with Makefile

Migrater can be build with Makefile. This requires Go version 1.23.5 or higher, Make version 4.3 or higher and a command-line that supports bash commands.

To build Migrater with Make, run the following command in the root folder of the project:

```bash
make build
```

The executable Migrater file can then be found in the directory `bin` with the name `migrater.exe`.

To build and directly run Migrate, run the following command in the root folder of the project:

```bash
make run
```

Thereby the executable Migrater file is only rebuild when the source code was changed since the last build process.

To remoe the executable Migrater file, run the following command in the root folder of the project:

```bash
make clean
```
