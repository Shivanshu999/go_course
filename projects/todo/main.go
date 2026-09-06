package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"todo/db"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

var todos []Todo

func addTodo() {

	scanner := bufio.NewScanner(os.Stdin)

	var title string

	fmt.Print("Enter the title:")
	if scanner.Scan() {
		title = scanner.Text()
	}

	todo := Todo{
		ID:        len(todos) + 1,
		Title:     title,
		Completed: false,
	}
	todos = append(todos, todo)

	err := scanner.Err()

	if err != nil {
		fmt.Println("Error reading input:", err)
	}

	fmt.Println("Todo added successfully!")
}

func listTodos() {
	fmt.Println("\n--- Todo List ---")
	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return
	}
	for _, todo := range todos {
		status := "Pending"

		if todo.Completed {
			status = "Completed"
		}
		fmt.Printf("ID: %d, Title: %s, Status: %s\n", todo.ID, todo.Title, status)
	}
}

func showMenu() {

	fmt.Println("\n======================")
	fmt.Println("       TODO APP")
	fmt.Println("======================")
	fmt.Println("1. Add Todo")
	fmt.Println("2. List Todos")
	fmt.Println("3. Complete Todo")
	fmt.Println("4. Delete Todo")
	fmt.Println("5. Exit")
	fmt.Println("======================")

}

func completedTodo() {
	var id int

	fmt.Print("Enter todo id:")
	fmt.Scan(&id)

	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = true
			fmt.Println("Todo marked as completed!")
			return
		}
	}
	fmt.Println("Todo not found.")
}

func deleteTodo() {

	var id int
	fmt.Print("Enter todo ID to delete: ")
	fmt.Scan(&id)
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println("Todo deleted successfully!")
			return
		}
	}
	fmt.Println("Todo not found.")

}

func main() {

	conn, err := db.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close(context.Background())
	fmt.Println("connected to DB")

	for {
		showMenu()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			addTodo()
		case 2:
			listTodos()
		case 3:
			completedTodo()
		case 4:
			deleteTodo()
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}
