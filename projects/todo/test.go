package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"todo/db"

	"github.com/jackc/pgx/v5"
)

type Todos struct {
	ID        int
	Title     string
	Completed bool
}

func addsTodo(conn *pgx.Conn, scanner *bufio.Scanner) {

	fmt.Print("Enter todo title: ")

	if !scanner.Scan() {
		fmt.Println("Error reading input")
		return
	}

	title := strings.TrimSpace(scanner.Text())

	if title == "" {
		fmt.Println("Todo title cannot be empty.")
		return
	}

	_, err := conn.Exec(
		context.Background(),
		"INSERT INTO todos (title) VALUES ($1)",
		title,
	)

	if err != nil {
		fmt.Println("Error adding todo:", err)
		return
	}

	fmt.Println("Todo added successfully!")
}

func listsTodos(conn *pgx.Conn) {

	rows, err := conn.Query(
		context.Background(),
		"SELECT id, title, completed FROM todos ORDER BY id",
	)

	if err != nil {
		fmt.Println("Error fetching todos:", err)
		return
	}

	defer rows.Close()

	fmt.Println("\n--- Todo List ---")

	found := false

	for rows.Next() {

		found = true

		var todo Todo

		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Completed,
		)

		if err != nil {
			fmt.Println("Error reading todo:", err)
			return
		}

		status := "Pending"

		if todo.Completed {
			status = "Completed"
		}

		fmt.Printf(
			"ID: %d | Title: %s | Status: %s\n",
			todo.ID,
			todo.Title,
			status,
		)
	}

	if !found {
		fmt.Println("No todos found.")
	}
}

func completesTodo(conn *pgx.Conn, scanner *bufio.Scanner) {

	fmt.Print("Enter todo ID: ")

	if !scanner.Scan() {
		fmt.Println("Error reading input")
		return
	}

	id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	if err != nil {
		fmt.Println("Please enter a valid ID.")
		return
	}

	result, err := conn.Exec(
		context.Background(),
		"UPDATE todos SET completed = true WHERE id = $1",
		id,
	)

	if err != nil {
		fmt.Println("Error completing todo:", err)
		return
	}

	if result.RowsAffected() == 0 {
		fmt.Println("Todo not found.")
		return
	}

	fmt.Println("Todo marked as completed!")
}

func deletesTodo(conn *pgx.Conn, scanner *bufio.Scanner) {

	fmt.Print("Enter todo ID: ")

	if !scanner.Scan() {
		fmt.Println("Error reading input")
		return
	}

	id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	if err != nil {
		fmt.Println("Please enter a valid ID.")
		return
	}

	result, err := conn.Exec(
		context.Background(),
		"DELETE FROM todos WHERE id = $1",
		id,
	)

	if err != nil {
		fmt.Println("Error deleting todo:", err)
		return
	}

	if result.RowsAffected() == 0 {
		fmt.Println("Todo not found.")
		return
	}

	fmt.Println("Todo deleted successfully!")
}

func showMenus() {

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

func mainer() {

	conn, err := db.ConnectDB()

	if err != nil {
		fmt.Println("Database error:", err)
		return
	}

	defer conn.Close(context.Background())

	fmt.Println("Connected to Neon PostgreSQL!")

	scanner := bufio.NewScanner(os.Stdin)

	for {

		showMenus()

		fmt.Print("Enter your choice: ")

		if !scanner.Scan() {
			fmt.Println("Error reading input")
			return
		}

		choice, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		switch choice {

		case 1:
			addsTodo(conn, scanner)

		case 2:
			listsTodos(conn)

		case 3:
			completesTodo(conn, scanner)

		case 4:
			deletesTodo(conn, scanner)

		case 5:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
