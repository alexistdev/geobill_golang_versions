package repository

import (
	"database/sql"
	"errors"
	"geobill_golang_versions/models"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrDuplicateUsername = errors.New("username already exists")
)

type Repository interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
}

type MySQLRepository struct {
	DB *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{DB: db}
}

func (r *MySQLRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (username, password, role, created_at) VALUES (?, ?, ?, NOW())"
	result, err := r.DB.Exec(query, user.Username, user.Password, user.Role)
	if err != nil {
		// Assuming MySQL error 1062 is duplicate entry.
		// For simplicity, just returning error. Real prod code checks driver specific errors.
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}

func (r *MySQLRepository) GetUserByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, password, role, created_at FROM users WHERE username = ?"
	row := r.DB.QueryRow(query, username)

	var user models.User
	// Scan directly. If parseTime=true is not used in DSN, we might need to handle time parsing.
	// But assuming it will be set.
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
