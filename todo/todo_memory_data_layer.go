package todo

import (
	"context"
	"test-mongo/models"
)

type TodoMemoryDataLayer struct {
	todos []models.Todo
}

func NewTodoMemoryDataLayer() *TodoMemoryDataLayer {
	return &TodoMemoryDataLayer{
		todos: []models.Todo{
			{
				UUID:      "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
				Title:     "Out of Bound",
				Completed: false,
				Index:     0,
			},
		},
	}
}

func (data *TodoMemoryDataLayer) GetTodos(ctx context.Context) ([]models.Todo, error) {
	return data.todos, nil
}

func (data *TodoMemoryDataLayer) InsertTodo(ctx context.Context, todo models.Todo) error {
	//select {
	//case <-time.After(10 * time.Second):
	data.todos = append(data.todos, todo)
	//case <-ctx.Done():
	//	return ctx.Err()
	//}
	return nil
}

func (data *TodoMemoryDataLayer) UpdateTodo(ctx context.Context, index int) {
	data.todos[index].Completed = true
}

func (data *TodoMemoryDataLayer) DeleteTodo(ctx context.Context, index int) {
	var newData []models.Todo
	for idx := range data.todos {
		if idx < index {
			newData = append(newData, data.todos[idx])
		} else if idx > index {
			data.todos[idx].Index--
			newData = append(newData, data.todos[idx])
		}
	}
	data.todos = newData
}
