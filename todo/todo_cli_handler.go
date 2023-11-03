package todo

import (
	"bufio"
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"test-mongo/models"
)

var counter int = 0

func randomUUID() string {
	uuidBytes := make([]byte, 16)
	_, err := rand.Read(uuidBytes)
	if err != nil {
		return ""
	}

	// Set the version and variant bits
	uuidBytes[6] = (uuidBytes[6] & 0x0F) | 0x40 // Version 4
	uuidBytes[8] = (uuidBytes[8] & 0x3F) | 0x80 // Variant 1

	// Format the UUID
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", uuidBytes[0:4], uuidBytes[4:6], uuidBytes[6:8], uuidBytes[8:10], uuidBytes[10:])
	return uuid
}

// GET - /todos
// return all todos
type TodoCLIHandler struct {
	scanner *bufio.Scanner
	uc      *TodoUsecase
}

func NewTodoHandler(uc *TodoUsecase) *TodoCLIHandler {
	return &TodoCLIHandler{
		scanner: bufio.NewScanner(os.Stdin),
		uc:      uc,
	}
}

func (h *TodoCLIHandler) GetTodos(ctx context.Context) []models.Todo {
	todos, err := h.uc.GetTodos(ctx)
	if err != nil {
		panic(err)
	}

	return todos
}

func (h *TodoCLIHandler) GetTodo(ctx context.Context, id int) models.Todo {
	todos, err := h.uc.GetTodos(ctx)
	if err != nil {
		panic(err)
	}

	return todos[id]
}

func (h TodoCLIHandler) InsertTodo(ctx context.Context, name string) {

	go func(todoTitle string) {
		fmt.Println("Inserting ToDo...")
		counter++
		if err := h.uc.InsertTodo(ctx, models.Todo{
			UUID:      randomUUID(),
			Title:     todoTitle,
			Completed: false,
			Index:     counter,
		}); err != nil {
			fmt.Println("Abandoned all hope to insert")
			return
		}
		fmt.Println("Insert Successful")
	}(name)
}
