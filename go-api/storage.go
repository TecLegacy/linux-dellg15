package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	DB *sql.DB
}

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

func NewPostgesStore() (*PostgresStore, error) {
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
		DB: db,
	}, nil

}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
	query := `create table if not exists Accounts(
		id serial primary key,
		firstname varchar(50),
		lastname varchar(50),
		accountnumber serial,
		balance integer,
		created_at timestamp
	)`
	result, err := s.DB.Exec(query)

	fmt.Printf("result of query %v ", result)
	return err
}

// Implementing Storage Interface
func (s *PostgresStore) CreateAccount(*Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(int) error {
	return nil
}
func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgresStore) GetAccountByID(int) (*Account, error) {
	return nil, nil
}
