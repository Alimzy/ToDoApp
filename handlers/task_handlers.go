package handlers

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"gotask/models"

	"github.com/gin-gonic/gin"
)

// In-memory store (replaced by a real DB in Weekend 2)
var (
	tasks  []models.Task
	nextID uint       = 1
	mu     sync.Mutex // protects tasks from concurrent access
)

// GetTasks returns all tasks as JSON
func GetTasks(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// CreateTask reads JSON from the request body and adds a new task
func CreateTask(c *gin.Context) {
	var input models.CreateTaskInput

	// ShouldBindJSON parses the body AND validates required fields
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	newTask := models.Task{
		ID:          nextID,
		Title:       input.Title,
		Description: input.Description,
		Status:      models.StatusPending,
		CreatedAt:   time.Now(),
	}
	nextID++
	tasks = append(tasks, newTask)

	c.JSON(http.StatusCreated, gin.H{"data": newTask})
}

// GetTask returns a single task by its ID from the URL
func GetTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, t := range tasks {
		if t.ID == uint(id) {
			c.JSON(http.StatusOK, gin.H{"data": t})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}
