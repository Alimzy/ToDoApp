package routes
 
import (
    "github.com/gin-gonic/gin"
    "gotask/handlers"
    "gotask/middleware"
)
 
func SetupRoutes(r *gin.Engine) {
    api := r.Group("/api/v1")
 
    // Public routes — no auth needed
    auth := api.Group("/auth")
    {
        auth.POST("/register", handlers.RegisterUser)
        auth.POST("/login",    handlers.LoginUser)
    }
 
    // Protected routes — RequireAuth middleware runs first
    tasks := api.Group("/tasks")
    tasks.Use(middleware.RequireAuth)
    {
        tasks.GET("",      handlers.GetTasks)
        tasks.POST("",     handlers.CreateTask)
        tasks.GET("/:id",  handlers.GetTask)
        tasks.PUT("/:id",  handlers.UpdateTask)
        tasks.DELETE("/:id", handlers.DeleteTask)
    }
}
