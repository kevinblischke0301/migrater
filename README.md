# Migrater

Migrater is a leightweight command-line tool to perform migrations and according rollbacks in the development process with sql files.

## Usage

### Options

Migrater takes a command-line flag to signal the execution of either migrations or rollbacks. Thereby migrations are performed per default when no command-line flags are provided.

These can be prefixed with one or two dashes. The following command-line flags exists:

- `migrate`: perform migrations
- `rollback`: perform rollbacks

### Configurations

Migrater uses environment variables for configuration. These can be set inside the file `.env` in the root folder to simulate and overwrite environment variables. The following environment variables must be set:

**MySQL**:

- `DB_TYPE`: `"mysql"` for MySQL
- `DB_NETWORK`: the used network (e. g. `"tcp"`)
- `DB_HOST`: the host on which MySQL runs
- `DB_PORT`: the port on which MySQL listens
- `DB_DATABASE`: the name of the database to access
- `DB_USER`: the username of the user
- `DB_PASSWORD`: the password of the user
- `MIGRATION_DIR`: the path to the directory of the sql files for migrations (only for migrations)
- `ROLLBACK_DIR`: the path to the directory of the sql files for rollbacks (only for rollbacks)

**SQLite**:

- `DB_TYPE`: `"sqlite"` for SQLite
- `DB_DATABASE`: the path to the database file to access
- `MIGRATION_DIR`: the path to the directory of the sql files for migrations (only for migrations)
- `ROLLBACK_DIR`: the path to the directory of the sql files for rollbacks (only for rollbacks)

### Execution

Migrater executes the contents of every file of a specified directory on a specified database. This happens in ascending alphanumerical order of the file names if a migration is performed and in descending alphanumerical order of the file names if a rollback is performed. Therefore it is recommended to only place sql files in these directories and to give every file name an ordering prefix.

Examples for the usage of sql files for migrations and rollbacks can be found in the directories `example_migrations` and `example_rollbacks`. Thereby the directories are containing according examples for MySQL and SQLite.

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

Thereby an according command-line flag must be appended to that command.

### Building with Makefile

Migrater can be build with Makefile. This requires Go version 1.23.5 or higher, Make version 4.3 or higher and a command-line that supports bash commands.

To build Migrater with Make, run the following command in the root folder of the project:

```bash
make build
```

The executable Migrater file can then be found in the directory `bin` with the name `migrater.exe`.

To build and directly run Migrate, run one of the following according commands in the root folder of the project to perform a migration or a rollback:

```bash
make migrate
```

```bash
make rollback
```

Thereby the executable Migrater file is only rebuild when the source code was changed since the last build process.

To remoe the executable Migrater file, run the following command in the root folder of the project:

```bash
make clean
```
