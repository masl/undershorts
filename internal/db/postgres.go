package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/masl/undershorts/internal/utils"
)

type PostgresClient struct {
	db *sql.DB
}

type Short struct {
	ID        int
	ShortURL  string
	LongURL   string
	CreatedAt time.Time
}

func NewPostgres() (*PostgresClient, error) {
	db, err := sql.Open("postgres", utils.GetEnv("POSTGRES_CONNECTION_STRING", "user=postgres password=postgres dbname=postgres sslmode=disable"))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresClient{
		db: db,
	}, nil
}

func (client *PostgresClient) Close() error {
	return client.db.Close()
}

func (client *PostgresClient) AddURL(shortURL string, longURL string) error {

	_, err := client.db.Query("INSERT INTO urls (short_url, long_url) VALUES ($1, $2);", shortURL, longURL)
	if err != nil {
		return err
	}

	return nil
}

func (client *PostgresClient) GetShortByShortURL(shortURL string) (*Short, error) {
	rows, err := client.db.Query("SELECT id, short_url, long_url, created_at FROM urls WHERE short_url = $1;", shortURL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var short Short
	for rows.Next() {
		if err := rows.Scan(&short.ID, &short.ShortURL, &short.LongURL, &short.CreatedAt); err != nil {
			return nil, err
		}
	}

	return &short, nil
}
