package router

import (
	"github.com/gin-gonic/gin"
	"ottoDigital/handler"
)

func NewPublicRoute(r *gin.RouterGroup) {
	h := handler.NewHandler()

	userRoute := r.Group("/users")
	{
		userRoute.POST("/", h.Register)
		//userRoute.POST("/", h.Login)
		userRoute.GET("/", h.GetUsers)
		userRoute.GET("/:id", h.GetUserByID)
		userRoute.PUT("/:id", h.UpdateUser)
		userRoute.DELETE("/:id", h.DeleteUser)
	}

	tasksRoute := r.Group("/tasks")
	{
		tasksRoute.POST("/", h.CreateTask)
		tasksRoute.GET("/", h.GetTasks)
		tasksRoute.GET("/:id", h.GetTaskByID)
		tasksRoute.PUT("/:id", h.UpdateTask)
		tasksRoute.DELETE("/:id", h.DeleteTask)
	}
}
