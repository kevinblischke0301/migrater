package env

type Env struct {
	DBType       string
	DBNetwork    string
	DBHost       string
	DBPort       string
	DBDatabase   string
	DBUser       string
	DBPassword   string
	MigrationDir string
	RollbackDir  string
}
