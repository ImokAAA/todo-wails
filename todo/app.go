package main

import (
	"context"
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type App struct {
	ctx context.Context
	db  *sql.DB
}

type Task struct {
	ID       int64  `json:"id"`
	Text     string `json:"text"`
	DateTime string `json:"dateTime"`
	IsDone   bool   `json:"isDone"`
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	var err error
	a.db, err = sql.Open("sqlite", "tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	query := `
        CREATE TABLE IF NOT EXISTS tasks (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            text TEXT NOT NULL,
            dateTime TEXT NOT NULL,
            isDone INTEGER NOT NULL
        );
    `
	_, err = a.db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) shutdown(ctx context.Context) {
	if a.db != nil {
		a.db.Close()
	}
}

func (a *App) AddTask(text, dateTime string) (int64, error) {
	log.Println("Adding Task:", text, dateTime)
	query := "INSERT INTO tasks (text, dateTime, isDone) VALUES (?, ?, 0)"
	result, err := a.db.Exec(query, text, dateTime)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (a *App) GetTasks() ([]Task, error) {
	rows, err := a.db.Query("SELECT id, text, dateTime, isDone FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		var isDone int
		err := rows.Scan(&task.ID, &task.Text, &task.DateTime, &isDone)
		if err != nil {
			return nil, err
		}
		task.IsDone = isDone == 1
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (a *App) UpdateTask(id int64, isDone bool) error {
	query := "UPDATE tasks SET isDone = ? WHERE id = ?"
	_, err := a.db.Exec(query, isDone, id)
	return err
}

func (a *App) DeleteTask(id int64) error {
	query := "DELETE FROM tasks WHERE id = ?"
	_, err := a.db.Exec(query, id)
	return err
}
