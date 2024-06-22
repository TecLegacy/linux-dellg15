package user

import (
	"database/sql"
	"fmt"
	"log"

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
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
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

	log.Println(u)
	return u, nil
}

func (s *StoreRepo) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	u := new(types.User)
	if rows.Next() {
		u, err = ScanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil

}

func (s *StoreRepo) CreateUser(user types.User) error {
	stmt, err := s.db.Prepare("INSERT INTO users (firstName, lastName, email, password, createdAt) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
