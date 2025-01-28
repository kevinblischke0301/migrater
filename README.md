# Migrater

This is a leightweight command-line tool to execute migrations from sql-files.

## Usage

The following parameters must be set as environment variables:

**MySQL**:

- `DB_TYPE`: `"mysql"` for MySQL
- `DB_NETWORK`: the used network (e. g. `"tcp"`)
- `DB_HOST`: the host on which MySQL runs
- `DB_PORT`: the port on which MySQL listens
- `DB_USER`: the username of the user
- `DB_PASSWORD`: the password of the user
- `MIGRATION_DIR`: the directory path of the migration files

These can be defined inside a file of name `.env` in the root folder.

Every sql-file inside `MIGRATION_DIR` is executed sequentially in alphabetical order of the file names.
