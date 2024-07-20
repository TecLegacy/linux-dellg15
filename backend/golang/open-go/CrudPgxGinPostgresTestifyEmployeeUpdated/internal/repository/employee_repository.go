package repository

import (
	"context"
	"employeeTestify/internal/entity"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, employee entity.Employee) error
	GetEmployees(ctx context.Context) ([]entity.Employee, error)
	GetEmployeeByID(ctx context.Context, id int) (entity.Employee, error)
	UpdateEmployee(ctx context.Context, employee entity.Employee) error
	DeleteEmployee(ctx context.Context, id int) error
	GetEmployeesPaging(ctx context.Context, offset int, limit int, sortBy string, order string) ([]entity.Employee, error)
}

type employeeRepository struct {
	db *pgxpool.Pool
}

func NewEmployeeRepository(db *pgxpool.Pool) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) CreateEmployee(ctx context.Context, employee entity.Employee) error {
	_, err := r.db.Exec(ctx, "INSERT INTO employees (name) VALUES ($1)", employee.Name)
	return err
}

func (r *employeeRepository) GetEmployees(ctx context.Context) ([]entity.Employee, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []entity.Employee
	for rows.Next() {
		var employee entity.Employee
		if err := rows.Scan(&employee.ID, &employee.Name); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (r *employeeRepository) GetEmployeeByID(ctx context.Context, id int) (entity.Employee, error) {
	var employee entity.Employee
	err := r.db.QueryRow(ctx, "SELECT id, name FROM employees WHERE id = $1", id).Scan(&employee.ID, &employee.Name)
	if err != nil {
		return employee, err
	}
	return employee, nil
}

func (r *employeeRepository) UpdateEmployee(ctx context.Context, employee entity.Employee) error {
	_, err := r.db.Exec(ctx, "UPDATE employees SET name = $1 WHERE id = $2", employee.Name, employee.ID)
	return err
}

func (r *employeeRepository) DeleteEmployee(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "DELETE FROM employees WHERE id = $1", id)
	return err
}

func (r *employeeRepository) GetEmployeesPaging(ctx context.Context, limit int, offset int, sortBy string, order string) ([]entity.Employee, error) {
	var employees []entity.Employee
	// PAGE is LIMIT
	// PAGESIZE IF offset

	query := fmt.Sprintf("SELECT id, name FROM employees ORDER BY %s %s LIMIT $1 OFFSET $2", sortBy, order) //

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var employee entity.Employee

		err := rows.Scan(&employee.ID, &employee.Name)
		if err != nil {

			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}
