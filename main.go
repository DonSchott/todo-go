package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// ANSI color codes
const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorRed    = "\033[31m"
	ColorDim    = "\033[2m"
)

// Todo represents a single todo item
type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func todoFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".todo.json")
}

func loadTodos() []Todo {
	data, err := os.ReadFile(todoFilePath())
	if err != nil {
		return []Todo{}
	}
	
	var todos []Todo
	json.Unmarshal(data, &todos)
	return todos
}

func saveTodos(todos []Todo) {
	data, _ := json.MarshalIndent(todos, "", "  ")
	os.WriteFile(todoFilePath(), data, 0644)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [add|list|done|remove] [args]")
		return
	}

	command := os.Args[1]
	
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo add <text>")
			return
		}
		todos := loadTodos()
		maxID := 0
		for _, t := range todos {
			if t.ID > maxID {
				maxID = t.ID
			}
		}
		newTodo := Todo{ID: maxID + 1, Text: os.Args[2], Done: false}
		todos = append(todos, newTodo)
		saveTodos(todos)
		fmt.Printf("%sAdded: %s%s\n", ColorGreen, newTodo.Text, ColorReset)		
		
	case "list":
		todos := loadTodos()
		if len(todos) == 0 {
			fmt.Println("No todos yet!")
			return
		}
		for _, todo := range todos {
			if todo.Done {
				fmt.Printf("%s[x]%s %s%d  %s%s\n", 
					ColorGreen, ColorReset,
					ColorGreen+ColorDim, todo.ID, todo.Text, ColorReset)
			} else {
				fmt.Printf("%s[ ]%s %d  %s\n", 
					ColorYellow, ColorReset, todo.ID, todo.Text)
			}
		}
		
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo done <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		todos := loadTodos()
		found := false
		for i := range todos {
			if todos[i].ID == id {
				todos[i].Done = true
				found = true
				saveTodos(todos)
				fmt.Printf("%sCompleted: %s%s\n", ColorGreen, todos[i].Text, ColorReset)
				break
			}
		}
		if !found {
			fmt.Printf("%sTodo not found%s\n", ColorRed, ColorReset)
		}
		
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo remove <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		todos := loadTodos()
		found := false
		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				found = true
				saveTodos(todos)
				fmt.Printf("%sRemoved: %s%s\n", ColorYellow, todo.Text, ColorReset)
				break
			}
		}
		if !found {
			fmt.Printf("%sTodo not found%s\n", ColorRed, ColorReset)
		}
		
	default:
		fmt.Println("Unknown command:", command)
	}
}
