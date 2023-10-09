package storage

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	Index       int    `json:"index"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

const (
	DoingTasks string = "doingtasks"
	DoneTasks  string = "donetasks"
)

var pool *sql.DB // connect to database

// Open
func Open() error {
	var err error
	pool, err = sql.Open("sqlite3", "./storage/todolist.db")
	return err
}

// Close
func Close() {
	pool.Close()
}

// NewTask
func NewTask(title, des string) error {
	_, err := pool.Exec("INSERT INTO doingtasks (title, description) VALUES (?,?)", title, des)
	return err
}

// GetTasks returns all doing tasks or done tasks in order
func GetTasks(state string) ([]Task, error) {
	rows, err := pool.Query("SELECT * FROM " + state)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []Task
	i := 0
	for rows.Next() {
		var id int
		var title string
		var descri string
		var result Task
		if err := rows.Scan(&id, &title, &descri); err != nil {
			fmt.Println("error read rows,,,", err)
			return nil, err
		}
		result.Index = id
		result.Title = title
		result.Description = descri
		results = append(results, result)
		i++
	}

	return results, nil
}

// UpdateTask help update title or description of a task in DB
// pass "" if want to keep the collumn still
func UpdateTask(id int, title, des string) error {

	// check if id exists?
	rows, err := pool.Query("SELECT * FROM doingtasks where id = ?", id)
	if err != nil {
		return errors.New("error: query id error")
	}

	if !rows.Next() {
		return errors.New("error: no id exist")
	}

	rows.Close()
	_, err = pool.Exec("UPDATE doingtasks SET title = ?, description = ? WHERE id = ?", title, des, id)
	return err
}

// DeleteTask
func DeleteTask(id int) error {

	_, err := pool.Exec("DELETE FROM doingtasks WHERE id = ?", id)
	return err
}
