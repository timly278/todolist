package main

import (
	"fmt"

	"os"
	"todolist/storage"
	"todolist/handle"
	"github.com/gin-gonic/gin"
)



func main() {

	// Open DB
	if err := storage.Open(); err != nil {
		fmt.Println("can't open!", err)
		os.Exit(1)
	}
	defer storage.Close()

	server := gin.Default()

	server.GET("/todos", handle.GetTasks)
	server.POST("/todo", handle.CreateTask)
	server.PUT("/todo/:id", handle.UpdateTask)
	server.DELETE("/todo/:id", handle.DeleteTask)

	fmt.Println("Application is running on port 8080")
	server.Run(":8080")

}
