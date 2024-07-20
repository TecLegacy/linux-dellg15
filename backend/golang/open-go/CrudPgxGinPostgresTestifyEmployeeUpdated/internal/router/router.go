package router

import (
	"employeeTestify/internal/config"
	"employeeTestify/internal/controller"
	"employeeTestify/internal/repository"
	"employeeTestify/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	employeeRepo := repository.NewEmployeeRepository(config.DB)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeController := controller.NewEmployeeController(employeeService)

	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		log.Print("Manual log test")
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	r.POST("/employees", employeeController.CreateEmployee)
	r.GET("/employees", employeeController.GetEmployees)
	r.GET("/employees/:id", employeeController.GetEmployeeByID)
	r.PUT("/employees/:id", employeeController.UpdateEmployee)
	r.DELETE("/employees/:id", employeeController.DeleteEmployee)
	r.GET("/employeesPaging", employeeController.GetEmployeesPaging)

	return r
}
