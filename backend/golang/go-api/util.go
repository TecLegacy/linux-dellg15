package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ScanIntoAccountFromRow(row *sql.Row) (*Account, error) {
	account := new(Account)
	err := row.Scan(&account.ID, &account.FirstName, &account.LastName, &account.AccNumber, &account.Balance, &account.CreatedAt)
	return account, err
}

func ScanIntoAccountFromRows(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.AccNumber, &account.Balance, &account.CreatedAt)
	return account, err
}

func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id provided. %s", idStr)
	}

	return id, nil

}
