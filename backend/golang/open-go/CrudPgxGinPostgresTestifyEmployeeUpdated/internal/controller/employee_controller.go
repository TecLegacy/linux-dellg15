package controller

import (
	"context"
	"employeeTestify/internal/entity"
	"employeeTestify/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(service service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func (c *EmployeeController) CreateEmployee(ctx *gin.Context) {
	var employee entity.Employee
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.CreateEmployee(context.Background(), employee)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

func (c *EmployeeController) GetEmployees(ctx *gin.Context) {
	employees, err := c.service.GetEmployees(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, employees)
}

func (c *EmployeeController) GetEmployeeByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	employee, err := c.service.GetEmployeeByID(context.Background(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, employee)
}

func (c *EmployeeController) UpdateEmployee(ctx *gin.Context) {
	var employee entity.Employee
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.UpdateEmployee(context.Background(), employee)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, employee)
}

func (c *EmployeeController) DeleteEmployee(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = c.service.DeleteEmployee(context.Background(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}

func (c *EmployeeController) GetEmployeesPaging(ctx *gin.Context) {
	// PAGE is LIMIT
	// PAGESIZE IF offset
	var limit, offset int

	pageStr := ctx.DefaultQuery("page", "2")          // limit
	pageSizeStr := ctx.DefaultQuery("page_size", "1") // offset
	sortBy := ctx.DefaultQuery("sort_by", "id")
	order := ctx.DefaultQuery("order", "asc")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	//page 1 -> 5 records  , 0 -5
	// page 2 -> 5 records , 5 - 10
	//page -> limit =5 , 10 - 15  // page * 5

	if pageSize == 1 {
		limit = page
		offset = 0
	} else {
		limit = page
		offset = pageSize*limit - limit
	}

	employees, err := c.service.GetEmployeesPaging(context.Background(), limit, offset, sortBy, order)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, employees)
}
