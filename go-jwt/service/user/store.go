package user

import (
	"database/sql"
	"fmt"

	"github.com/teclegacy/golang-ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
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

	return u, nil
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

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}
