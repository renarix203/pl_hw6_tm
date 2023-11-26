package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	id          int
	title       string
	description string
	date        string
	priority    string
	status      string
}

func main() {

	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt :=
		`create table if not exists tasks (id integer, title text, description text, date text, priority text, status text);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []Item
	num := 0

	recoverItems(&tasks, db)
	num = tasks[len(tasks)-1].id

	for {
		showMenu()
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			createTask(&tasks, &num, db)
		case 2:
			showTasks(tasks)
		case 3:
			editTask(tasks, db)
		case 4:
			eraseTask(&tasks, db)
		case 5:
			showStatistics(tasks)
		default:
			fmt.Println("Invalid option")
		}
	}
}

func showMenu() {
	fmt.Println("Task Manager")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("1. Create a task")
	fmt.Println("2. Show existing tasks")
	fmt.Println("3. Edit a task")
	fmt.Println("4. Erase a task")
	fmt.Println("5. Statistics")
	fmt.Println()
}

func createTask(tasks *[]Item, num *int, db *sql.DB) {
	fmt.Println("Title:")
	var title string
	fmt.Scanln(&title)
	fmt.Println("Description:")
	var description string
	fmt.Scanln(&description)
	fmt.Println("Due Date:")
	var date string
	fmt.Scanln(&date)
	fmt.Println("Priority:")
	var priority string
	fmt.Scanln(&priority)
	status := "TO DO"

	*num++

	var task Item
	task.id = *num
	task.title = title
	task.description = description
	task.date = date
	task.priority = priority
	task.status = status

	*tasks = append(*tasks, task)

	sqlStmt := `insert into tasks values (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(sqlStmt, task.id, task.title, task.description, task.date, task.priority, task.status)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("Task Saved Scuccessfully!")
	fmt.Println()
}

func showTasks(tasks []Item) {
	fmt.Println("1. All Tasks")
	fmt.Println("2. TO DO Tasks")
	fmt.Println()
	var option int
	fmt.Scanln(&option)
	fmt.Println()
	switch option {
	case 1:
		fmt.Println("All Tasks:")
		for i := 0; i < len(tasks); i++ {
			fmt.Println("Title: ", tasks[i].title)
			fmt.Println("Description: ", tasks[i].description)
			fmt.Println("Due Date: ", tasks[i].date)
			fmt.Println("Priority: ", tasks[i].priority)
			fmt.Println("Status: ", tasks[i].status)
			fmt.Println()
		}
	case 2:
		fmt.Println("TO DO Tasks:")
		for i := 0; i < len(tasks); i++ {
			if tasks[i].status == "TO DO" {
				fmt.Println("Title: ", tasks[i].title)
				fmt.Println("Description: ", tasks[i].description)
				fmt.Println("Due Date: ", tasks[i].date)
				fmt.Println("Priority: ", tasks[i].priority)
				fmt.Println("Status: ", tasks[i].status)
				fmt.Println()
			}
		}
	}
}

func editTask(tasks []Item, db *sql.DB) {
	fmt.Println("What task do you want to edit?")
	fmt.Println()
	for i := 0; i < len(tasks); i++ {
		fmt.Println(i+1, tasks[i].title)
	}
	fmt.Println()
	var option int
	fmt.Scanln(&option)
	fmt.Println()
	fmt.Println("What do you want to edit?")
	fmt.Println()
	fmt.Println("1. Title")
	fmt.Println("2. Description")
	fmt.Println("3. Due Date")
	fmt.Println("4. Priority")
	fmt.Println("5. Status")
	fmt.Println()
	var option2 int
	fmt.Scanln(&option2)
	fmt.Println()
	switch option2 {
	case 1:
		fmt.Println("New Title:")
		var title string
		fmt.Scanln(&title)
		tasks[option-1].title = title

		sqlStmt := `update tasks set title = ? where id = ?`
		_, err := db.Exec(sqlStmt, title, (tasks)[option-1].id)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Title Updated")
	case 2:
		fmt.Println("New Description:")
		var description string
		fmt.Scanln(&description)
		tasks[option-1].description = description

		sqlStmt := `update tasks set description = ? where id = ?`
		_, err := db.Exec(sqlStmt, description, (tasks)[option-1].id)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Description Updated")
	case 3:
		fmt.Println("New Due Date:")
		var date string
		fmt.Scanln(&date)
		tasks[option-1].date = date

		sqlStmt := `update tasks set date = ? where id = ?`
		_, err := db.Exec(sqlStmt, date, (tasks)[option-1].id)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Due Date Updated")
	case 4:
		fmt.Println("New Priority:")
		var priority string
		fmt.Scanln(&priority)
		tasks[option-1].priority = priority

		sqlStmt := `update tasks set priority = ? where id = ?`
		_, err := db.Exec(sqlStmt, priority, (tasks)[option-1].id)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Priority Updated")
	case 5:
		var status string
		status = "DONE"
		tasks[option-1].status = status

		sqlStmt := `update tasks set status = ? where id = ?`
		_, err := db.Exec(sqlStmt, status, (tasks)[option-1].id)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Status Changed to DONE")
	}
}

func eraseTask(tasks *[]Item, db *sql.DB) {
	fmt.Println("What task do you want to erase?")
	fmt.Println()
	for i := 0; i < len(*tasks); i++ {
		fmt.Println(i+1, (*tasks)[i].title)
	}
	fmt.Println()
	var option int
	fmt.Scanln(&option)
	fmt.Println()

	sqlStmt := `delete from tasks where id = ?`
	_, err := db.Exec(sqlStmt, (*tasks)[option-1].id)
	if err != nil {
		log.Fatal(err)
	}

	*tasks = append((*tasks)[:option-1], (*tasks)[option:]...)
	fmt.Println("Task Erased")
	fmt.Println()
}

func showStatistics(tasks []Item) {
	fmt.Println("Statistics:")
	fmt.Println()
	n := 0
	done := 0
	for i := 0; i < len(tasks); i++ {
		if tasks[i].status == "DONE" {
			done++
		}
		n++
	}
	fmt.Println("Total Tasks: ", n)
	fmt.Println("Total Completed Tasks: ", done)
	fmt.Println("Total Pending Tasks: ", n-done)
	fmt.Println("Percentage of Pending tasks: ", (n-done)*100/n, "%")
	fmt.Println()
}

func recoverItems(tasks *[]Item, db *sql.DB) {
	rows, err := db.Query("select * from tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var task Item
		err = rows.Scan(&task.id, &task.title, &task.description, &task.date, &task.priority, &task.status)
		if err != nil {
			log.Fatal(err)
		}
		*tasks = append(*tasks, task)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
