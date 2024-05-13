package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

type Storage interface {
	CreateAccount(*Account) (*Account, error)
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
	GetAllAccounts() ([]*Account, error)
}

// NewPostgresStore creates a new instance of PostgresStore and establishes a connection to the PostgreSQL database.
func NewPostgresStore() (*PostgresStore, error) {
	// Connection to Postgres Docker img
	connStr := "user=postgres dbname=postgres password=mysecretpassword sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect with postgres in file storage.go: %v", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping postgres in file storage.go: %v", err)
	}

	return &PostgresStore{
		db: db,
	}, nil

}

// Init initializes the PostgresStore by creating the account table if it doesn't exist.
func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

// CreateAccountTable creates the account table in the PostgreSQL database if it doesn't exist.
func (s *PostgresStore) CreateAccountTable() error {
	query := `create table if not exists account(
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		account_number serial,
		balance serial,
		created_at timestamp
	)`
	result, err := s.db.Exec(query)

	if err != nil {
		return fmt.Errorf("failed to create account table in file storage.go: %v", err)
	}

	fmt.Printf("result of query %+v ", result)
	return nil
}

/*
--------------------------------------------------------
*/
// Implementing Storage Interface

// CreateAccount creates a new account in the PostgreSQL database.
func (s *PostgresStore) CreateAccount(a *Account) (*Account, error) {
	query := `INSERT INTO account(
			first_name,
			last_name,
			account_number,
			balance,
			created_at
			)
			VALUES($1,$2,$3,$4,$5)
			RETURNING *
			`
	row := s.db.QueryRow(
		query,
		a.FirstName,
		a.LastName,
		a.AccNumber,
		a.Balance,
		a.CreatedAt)

	account, err := ScanIntoAccountFromRow(row)
	if err != nil {
		return nil, fmt.Errorf("failed to create account in file storage.go: %v", err)
	}

	return account, nil
}

// DeleteAccount deletes an account from the PostgreSQL database based on the provided ID.
func (s *PostgresStore) DeleteAccount(id int) error {
	query := `SELECT id FROM account WHERE id = $1`

	row := s.db.QueryRow(query, id)

	var retrievedID int
	err := row.Scan(&retrievedID)

	if err == sql.ErrNoRows {
		return fmt.Errorf("no account found with ID %d", id)
	} else if err != nil {
		return err
	}

	// If the account exists
	query = `DELETE FROM account WHERE id = $1`

	_, err = s.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

// UpdateAccount updates an existing account in the PostgreSQL database.
func (s *PostgresStore) UpdateAccount(a *Account) error {
	// TODO: Implement the logic to update an existing account in the database.
	return nil
}

// GetAccountByID retrieves an account from the PostgreSQL database based on the provided ID.
func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	query := `SELECT * FROM account
		WHERE id = $1
	`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query account with id %v in file storage.go: %v", id, err)
	}

	for rows.Next() {
		account, err := ScanIntoAccountFromRows(rows)
		return account, err
	}

	return nil, fmt.Errorf("no account found for id %v in file storage.go", id)
}

// GetAllAccounts retrieves all accounts from the PostgreSQL database.
func (s *PostgresStore) GetAllAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, fmt.Errorf("failed to query all accounts in file storage.go: %v", err)
	}

	accounts := []*Account{}
	for rows.Next() {
		account, err := ScanIntoAccountFromRows(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan account from rows in file storage.go: %v", err)
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}
