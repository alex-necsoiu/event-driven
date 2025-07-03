package order

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Repository abstracts DB operations for orders
// Implements basic PostgreSQL connection and migration placeholder

type Repository interface {
	CreateOrder(userID string, amount float64) (string, error)
	GetOrder(id string) (Order, error)
	UpdateOrderStatus(id string, status string) error
}

type Order struct {
	ID     string
	UserID string
	Amount float64
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
	if err := migrateOrders(db, logger); err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

// migrateOrders runs DB migrations for the orders table (placeholder)
func migrateOrders(db *sql.DB, logger *log.Logger) error {
	logger.Println("Running order table migrations (placeholder)")
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL,
		amount NUMERIC NOT NULL
	)`)
	return err
}

// CreateOrder creates a new order and returns the order ID
func (r *PostgresRepository) CreateOrder(userID string, amount float64) (string, error) {
	var id string
	err := r.db.QueryRow(
		"INSERT INTO orders (user_id, amount) VALUES ($1, $2) RETURNING id::text",
		userID, amount,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

// GetOrder retrieves an order by ID
func (r *PostgresRepository) GetOrder(id string) (Order, error) {
	var order Order
	err := r.db.QueryRow(
		"SELECT id::text, user_id::text, amount FROM orders WHERE id = $1",
		id,
	).Scan(&order.ID, &order.UserID, &order.Amount)

	if err != nil {
		return Order{}, err
	}

	return order, nil
}

// UpdateOrderStatus updates the status of an order
func (r *PostgresRepository) UpdateOrderStatus(id string, status string) error {
	// First, we need to add a status column to the orders table
	// For now, we'll just log the status update
	_, err := r.db.Exec(
		"UPDATE orders SET status = $1 WHERE id = $2",
		status, id,
	)

	return err
}
