package user

import (
	"database/sql"
	"fmt"

	"github.com/teclegacy/golang-ecom/types"
)

type StoreRepo struct {
	db *sql.DB
}

func NewStoreRepo(db *sql.DB) *StoreRepo {
	return &StoreRepo{
		db: db,
	}
}

func ScanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	u := new(types.User)
	if err := rows.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.CreatedAt); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *StoreRepo) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM user WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)

	for rows.Next() {
		u, err = ScanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *StoreRepo) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (s *StoreRepo) CreateUser(user types.User) error {
	stmt, err := s.db.Prepare("INSERT INTO users (fistName, lastName, email, password, created_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
