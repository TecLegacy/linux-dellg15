package service

import (
	"context"
	"employeeTestify/internal/entity"
	"employeeTestify/internal/repository"
)

type EmployeeService interface {
	CreateEmployee(ctx context.Context, employee entity.Employee) error
	GetEmployees(ctx context.Context) ([]entity.Employee, error)
	GetEmployeeByID(ctx context.Context, id int) (entity.Employee, error)
	UpdateEmployee(ctx context.Context, employee entity.Employee) error
	DeleteEmployee(ctx context.Context, id int) error
	GetEmployeesPaging(ctx context.Context, page int, pageSize int, sortBy string, order string) ([]entity.Employee, error)
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) CreateEmployee(ctx context.Context, employee entity.Employee) error {
	return s.repo.CreateEmployee(ctx, employee)
}

func (s *employeeService) GetEmployees(ctx context.Context) ([]entity.Employee, error) {
	return s.repo.GetEmployees(ctx)
}

func (s *employeeService) GetEmployeeByID(ctx context.Context, id int) (entity.Employee, error) {
	return s.repo.GetEmployeeByID(ctx, id)
}

func (s *employeeService) UpdateEmployee(ctx context.Context, employee entity.Employee) error {
	return s.repo.UpdateEmployee(ctx, employee)
}

func (s *employeeService) DeleteEmployee(ctx context.Context, id int) error {
	return s.repo.DeleteEmployee(ctx, id)
}

func (s *employeeService) GetEmployeesPaging(ctx context.Context, page int, pageSize int, sortBy string, order string) ([]entity.Employee, error) {

	// PAGE is LIMIT
	// PAGESIZE IF offset

	return s.repo.GetEmployeesPaging(ctx, page, pageSize, sortBy, order)
}
