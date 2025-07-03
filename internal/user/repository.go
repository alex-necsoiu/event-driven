package user

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// To support MongoDB or other DBs, implement the Repository interface and add a factory method.

// Repository abstracts DB operations for users
// Implements basic PostgreSQL connection and migration placeholder

type Repository interface {
	CreateUser(name, email string) (string, error)
	GetUser(id string) (User, error)
}

type User struct {
	ID    string
	Name  string
	Email string
}

// PostgresRepository implements Repository using PostgreSQL
type PostgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository creates a new PostgresRepository
func NewPostgresRepository(connStr string, logger *log.Logger) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	// Run migrations (placeholder)
	if err := migrateUsers(db, logger); err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

// migrateUsers runs DB migrations for the users table (placeholder)
func migrateUsers(db *sql.DB, logger *log.Logger) error {
	logger.Println("Running user table migrations (placeholder)")
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	)`)
	return err
}

// CreateUser creates a new user and returns the user ID
func (r *PostgresRepository) CreateUser(name, email string) (string, error) {
	var id string
	err := r.db.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id::text",
		name, email,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

// GetUser retrieves a user by ID
func (r *PostgresRepository) GetUser(id string) (User, error) {
	var user User
	err := r.db.QueryRow(
		"SELECT id::text, name, email FROM users WHERE id = $1",
		id,
	).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
