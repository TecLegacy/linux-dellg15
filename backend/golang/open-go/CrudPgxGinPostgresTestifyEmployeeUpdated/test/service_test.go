// package test

// import (
// 	"context"
// 	"employeeTestify/internal/entity"
// 	"employeeTestify/internal/service"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MockEmployeeRepository struct {
// 	mock.Mock
// }

// func (m *MockEmployeeRepository) CreateEmployee(ctx context.Context, employee entity.Employee) error {
// 	args := m.Called(ctx, employee)
// 	return args.Error(0)
// }

// func (m *MockEmployeeRepository) GetEmployees(ctx context.Context) ([]entity.Employee, error) {
// 	args := m.Called(ctx)
// 	return args.Get(0).([]entity.Employee), args.Error(1)
// }

// func (m *MockEmployeeRepository) GetEmployeeByID(ctx context.Context, id int) (entity.Employee, error) {
// 	args := m.Called(ctx, id)
// 	return args.Get(0).(entity.Employee), args.Error(1)
// }

// func (m *MockEmployeeRepository) UpdateEmployee(ctx context.Context, employee entity.Employee) error {
// 	args := m.Called(ctx, employee)
// 	return args.Error(0)
// }

// func (m *MockEmployeeRepository) DeleteEmployee(ctx context.Context, id int) error {
// 	args := m.Called(ctx, id)
// 	return args.Error(0)
// }

// func TestEmployeeService_CreateEmployee(t *testing.T) {
// 	mockRepo := new(MockEmployeeRepository)
// 	service := service.NewEmployeeService(mockRepo)

// 	employee := entity.Employee{Name: "John Doe"}
// 	mockRepo.On("CreateEmployee", mock.Anything, employee).Return(nil)

// 	err := service.CreateEmployee(context.Background(), employee)
// 	assert.NoError(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// func TestEmployeeService_GetEmployees(t *testing.T) {
// 	mockRepo := new(MockEmployeeRepository)
// 	service := service.NewEmployeeService(mockRepo)

// 	expectedEmployees := []entity.Employee{{ID: 1, Name: "John Doe"}}
// 	mockRepo.On("GetEmployees", mock.Anything).Return(expectedEmployees, nil)

// 	employees, err := service.GetEmployees(context.Background())
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedEmployees, employees)
// 	mockRepo.AssertExpectations(t)
// }

// func TestEmployeeService_GetEmployeesPaging(t *testing.T) {
// 	mockRepo := new(MockEmployeeRepository)
// 	service := service.NewEmployeeService(mockRepo)

// 	expectedEmployees := []entity.Employee{{ID: 1, Name: "John Doe"}}
// 	mockRepo.On("GetEmployeesPaging", mock.Anything).Return(expectedEmployees, nil)

// 	employees, err := service.GetEmployeesPaging(context.Background(),1,3,"","")
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedEmployees, employees)
// 	mockRepo.AssertExpectations(t)
// }
// func TestEmployeeService_GetEmployeeByID(t *testing.T) {
// 	mockRepo := new(MockEmployeeRepository)
// 	service := service.NewEmployeeService(mockRepo)

// 	expectedEmployee := entity.Employee{ID: 1, Name: "John Doe"}
// 	mockRepo.On("GetEmployeeByID", mock.Anything, 1).Return(expectedEmployee, nil)

// 	employee, err := service.GetEmployeeByID(context.Background(), 1)
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedEmployee, employee)
// 	mockRepo.AssertExpectations(t)
// }

// func TestEmployeeService_UpdateEmployee(t *testing.T) {
// 	mockRepo := new(MockEmployeeRepository)
// 	service := service.NewEmployeeService(mockRepo)

// 	employee := entity.Employee{ID: 1, Name: "John Doe"}
// 	mockRepo.On("UpdateEmployee", mock.Anything, employee).Return(nil)

// 	err := service.UpdateEmployee(context.Background(), employee)
// 	assert.NoError(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// func TestEmployeeService_DeleteEmployee(t *testing.T) {
// 	mockRepo := new(MockEmployeeRepository)
// 	service := service.NewEmployeeService(mockRepo)

// 	mockRepo.On("DeleteEmployee", mock.Anything, 1).Return(nil)

// 	err := service.DeleteEmployee(context.Background(), 1)
// 	assert.NoError(t, err)
// 	mockRepo.AssertExpectations(t)
// }
