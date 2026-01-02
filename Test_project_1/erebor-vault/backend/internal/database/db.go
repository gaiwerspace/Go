package database
import ("database/sql"; "embed"; "fmt"; "log"; "github.com/golang-migrate/migrate/v4"; "github.com/golang-migrate/migrate/v4/database/postgres"; "github.com/golang-migrate/migrate/v4/source/iofs"; _ "github.com/lib/pq")
//go:embed migrations/*.sql
var migrationFS embed.FS
func NewDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil { return nil, err }
	if err := db.Ping(); err != nil { return nil, err }
	log.Println("Database connected")
	return db, nil
}
func RunMigrations(db *sql.DB) error {
	_, _ = db.Exec("CREATE SCHEMA IF NOT EXISTS public")
	driver, err := postgres.WithInstance(db, &postgres.Config{MigrationsTable: "schema_migrations", SchemaName: "public"})
	if err != nil { return err }
	d, err := iofs.New(migrationFS, "migrations")
	if err != nil { return err }
	m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)
	if err != nil { return err }
	if err := m.Up(); err != nil && err != migrate.ErrNoChange { return err }
	log.Println("Migrations complete")
	return nil
}
