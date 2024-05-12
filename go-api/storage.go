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

func NewPostgresStore() (*PostgresStore, error) {
	// Connection to Postgres Docker img
	connStr := "user=postgres dbname=postgres password=mysecretpassword sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect with postgres %v", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil

}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

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

	fmt.Printf("result of query %+v ", result)
	return err
}

/*
--------------------------------------------------------
*/
// Implementing Storage Interface
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
	row := s.db.QueryRow(query, a.FirstName, a.LastName, a.AccNumber, a.Balance, a.CreatedAt)
	account, err := ScanIntoAccountFromRow(row)
	if err != nil {
		return nil, fmt.Errorf("sq %v", err)
	}

	return account, nil
}

func (s *PostgresStore) DeleteAccount(int) error {
	return nil
}

func (s *PostgresStore) UpdateAccount(a *Account) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	query := `SELECT * FROM account
		WHERE id = $1
	`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		account, err := ScanIntoAccountFromRows(rows)
		return account, err
	}

	return nil, fmt.Errorf("no account found for id %v", id)
}

func (s *PostgresStore) GetAllAccounts() ([]*Account, error) {

	rows, err := s.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account, err := ScanIntoAccountFromRows(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}
