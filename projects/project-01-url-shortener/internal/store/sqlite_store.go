package store

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"letscode/project-01-url-shortener/internal/model"

	_ "github.com/mattn/go-sqlite3"
)

// SQLiteStore implements Store using a SQLite database with a unique code and URL map.
type SQLiteStore struct {
	db *sql.DB
}

// NewSQLiteStore opens a SQLite-backed URL repository.
func NewSQLiteStore(dsn string) (*SQLiteStore, error) {
	if strings.TrimSpace(dsn) == "" {
		dsn = ":memory:"
	}

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping sqlite: %w", err)
	}

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS short_urls (
			code TEXT PRIMARY KEY,
			url TEXT NOT NULL UNIQUE,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`); err != nil {
		db.Close()
		return nil, fmt.Errorf("create short_urls table: %w", err)
	}

	return &SQLiteStore{db: db}, nil
}

func (s *SQLiteStore) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}

func (s *SQLiteStore) Save(code, url string) error {
	if strings.TrimSpace(code) == "" {
		return fmt.Errorf("code is required")
	}
	if strings.TrimSpace(url) == "" {
		return fmt.Errorf("url is required")
	}

	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	var existingURL string
	err = tx.QueryRow(`SELECT url FROM short_urls WHERE code = ?`, code).Scan(&existingURL)
	switch {
	case err == nil && existingURL == url:
		return tx.Commit()
	case err == nil && existingURL != url:
		return fmt.Errorf("code already exists")
	case err != nil && !errors.Is(err, sql.ErrNoRows):
		return fmt.Errorf("lookup code: %w", err)
	}

	var existingCode string
	err = tx.QueryRow(`SELECT code FROM short_urls WHERE url = ?`, url).Scan(&existingCode)
	switch {
	case err == nil && existingCode == code:
		return tx.Commit()
	case err == nil && existingCode != code:
		return fmt.Errorf("url already exists with different code")
	case err != nil && !errors.Is(err, sql.ErrNoRows):
		return fmt.Errorf("lookup url: %w", err)
	}

	_, err = tx.Exec(`INSERT INTO short_urls (code, url) VALUES (?, ?)`, code, url)
	if err != nil {
		return fmt.Errorf("insert short_url: %w", err)
	}

	return tx.Commit()
}

func (s *SQLiteStore) Lookup(code string) (string, bool) {
	var url string
	err := s.db.QueryRow(`SELECT url FROM short_urls WHERE code = ?`, code).Scan(&url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", false
		}
		return "", false
	}
	return url, true
}

func (s *SQLiteStore) FindByURL(url string) (string, bool) {
	var code string
	err := s.db.QueryRow(`SELECT code FROM short_urls WHERE url = ?`, url).Scan(&code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", false
		}
		return "", false
	}
	return code, true
}

func (s *SQLiteStore) All() ([]model.URLMapping, error) {
	rows, err := s.db.Query(`SELECT code, url FROM short_urls ORDER BY created_at ASC, code ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]model.URLMapping, 0)
	for rows.Next() {
		var mapping model.URLMapping
		if err := rows.Scan(&mapping.Code, &mapping.URL); err != nil {
			return nil, err
		}
		out = append(out, mapping)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
