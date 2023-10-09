package handle

import (
	"fmt"
	"net/http"
	"strconv"
	"todolist/storage"

	"github.com/gin-gonic/gin"
)

// GetTasks retrieve doing task from DB and send to client
func GetTasks(ctx *gin.Context) {
	todos, err := storage.GetTasks(storage.DoingTasks)
	if err != nil {
		ctx.JSON(500, "error: something went wrong with MarshalIndent")
		return
	}
	ctx.IndentedJSON(http.StatusOK, todos)
}

// CreateTask creates a new task
func CreateTask(ctx *gin.Context) {
	var newTask storage.Task

	// Call BindJSON to bind the received JSON to
	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.JSON(500, "error: something went wrong with BindJSON")
		return
	}

	// Add the new task to DB
	if err := storage.NewTask(newTask.Title, newTask.Description); err != nil {
		ctx.JSON(500, "error: something went wrong with storage.NewTask")
	}

	// Respond to client
	GetTasks(ctx)
}

// UpdateTask update a task by its id
func UpdateTask(ctx *gin.Context) {

	// Call BindJSON to bind the received JSON to
	var newTask storage.Task
	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.JSON(500, "error: something went wrong with BindJSON")
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	// fmt.Printf("id = %s\n, newTask = %v\n", id, newTask)

	// Update todo to the DB.
	if err := storage.UpdateTask(id, newTask.Title, newTask.Description); err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	// Respond to client
	GetTasks(ctx)
}

// DeleteTask deletes by id
func DeleteTask(ctx *gin.Context) {

	idtext := ctx.Param("id")
	id, _ := strconv.Atoi(idtext[1:]) // Param return: idtext = ":id"

	if err := storage.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, "error: something went wrong with DeleteTask")
		return
	}

	// Respond to client
	GetTasks(ctx)
}
