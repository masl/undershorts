package storage

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/masl/undershorts/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type Storage interface {
	Ping() bool
	CreateUser(*models.UserRequest) (*models.User, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PSQL_HOST"), os.Getenv("PSQL_PORT"), os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASSWORD"), os.Getenv("PSQL_DATABASE"),
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil

}

// initialize the database
func (s *PostgresStore) Init() error {
	// install uuid extension
	_, err := s.db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		CREATE EXTENSION IF NOT EXISTS "citext";`)
	if err != nil {
		return err
	}

	// create user table
	_, err = s.db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id uuid DEFAULT uuid_generate_v4(),
		email citext NOT NULL UNIQUE,
		password_hash varchar(255) NOT NULL,
		created_at timestamp NOT NULL DEFAULT NOW(),
		PRIMARY KEY (id)
	);`)
	if err != nil {
		return err
	}

	return nil
}

// create user with email and password
func (s *PostgresStore) CreateUser(userReq *models.UserRequest) (*models.User, error) {
	query := `INSERT INTO users
	(email, password_hash)
	values ($1, $2)
	RETURNING id, email, password_hash, created_at`

	// make password hash
	hash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// run query
	rows, err := s.db.Query(
		query,
		userReq.Email,
		hash,
	)
	if err != nil {
		return nil, err
	}

	// return created user
	rows.Next()
	user := &models.User{}
	err = rows.Scan(
		&user.Id,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *PostgresStore) Ping() bool {
	err := s.db.Ping()
	if err != nil {
		return false
	}

	return true
}
