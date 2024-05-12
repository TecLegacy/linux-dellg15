package main

import "database/sql"

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
