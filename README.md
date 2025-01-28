# Migrater

Migrater is a leightweight command-line tool to perform migrations from sql-files.

## Usage

The following parameters must be set as environment variables or inside the file `.env` in the root folder:

**MySQL**:

- `DB_TYPE`: `"mysql"` for MySQL
- `DB_NETWORK`: the used network (e. g. `"tcp"`)
- `DB_HOST`: the host on which MySQL runs
- `DB_PORT`: the port on which MySQL listens
- `DB_DATABASE`: the name of the database to access
- `DB_USER`: the username of the user
- `DB_PASSWORD`: the password of the user
- `MIGRATION_DIR`: the path to the directory of the migration files

**SQLite**:

- `DB_TYPE`: `"sqlite"` for SQLite
- `DB_DATABASE`: the path to the database file to access
- `MIGRATION_DIR`: the path to the directory of the migration files

Every sql-file inside `MIGRATION_DIR` is executed sequentially in alphanumerical order of the file names. Therefore it is recommended to give every file name an ordering prefix.
