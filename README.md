# Migrater

This is a leightweight command-line tool to execute migrations from sql-files.

## Usage

The following parameters must be set as environment variables:

**MySQL**:

- DB_TYPE = "mysql"
- DB_NETWORK = <your_database_network>
- DB_HOST = <your_database_host>
- DB_PORT = <your_database_port>
- DB_USER = <your_database_username>
- DB_PASSWORD = <your_database_user_password>
- MIGRATION_DIR = <your_directory_with_migration_files>

These can be defined inside a file of name `.env`.

Every sql-file inside `MIGRATION_DIR` is executed sequentially in alphabetical order of the files.
