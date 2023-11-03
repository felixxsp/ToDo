package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"test-mongo/todo"

	"github.com/gin-gonic/gin"
)

// 3 - layers of a web application or clean architecture
// ----------------------------------------------------
// 1. Handler (deals with http requests) - GIN
// 2. Business Logic / Service / Usecase (deals with business logic) - GO
// 3. Data layer - MONGO

func Homepage() Menu {
	fmt.Print("1. View List; 2. Input new ToDo; 3. Exit: ")
	var num int
	fmt.Scan(&num)
	return Menu(num)
}

type Menu int

const (
	ViewList        Menu = 1
	CreateTodo      Menu = 2
	CancelAnyInsert Menu = 3
	Exit            Menu = 4
)

func main() {
	mainCtx := context.Background()
	//var cancel context.CancelFunc
	router := gin.Default()
	todoRepo := todo.NewTodoMemoryDataLayer()
	todoUC := todo.NewTodoUsecase(todoRepo)
	todoH := todo.NewTodoHandler(todoUC)

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	router.GET("/view", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, todoH.GetTodos(mainCtx))
	})

	router.GET("/view/:id", func(c *gin.Context) {
		idstr := c.Param("id")
		id, _ := strconv.Atoi(idstr)
		c.IndentedJSON(http.StatusOK, todoH.GetTodo(mainCtx, id))
	})

	router.POST("/add/:content", func(ctx *gin.Context) {
		name := ctx.Param("content")
		todoH.InsertTodo(ctx, name)
		ctx.IndentedJSON(http.StatusCreated, fmt.Sprintf("Succesfully Append: %s", name))
	})

	router.PUT("/complete/:index", func(ctx *gin.Context) {
		id := ctx.Param("index")
		index, _ := strconv.Atoi(id)
		todoRepo.UpdateTodo(mainCtx, index)
		ctx.IndentedJSON(http.StatusAccepted, fmt.Sprintf("Succesfully Updated: ", todoH.GetTodo(mainCtx, index)))
	})

	router.DELETE("/delete/:index", func(ctx *gin.Context) {
		id := ctx.Param("index")
		index, _ := strconv.Atoi(id)
		todoRepo.DeleteTodo(mainCtx, index)
		ctx.IndentedJSON(http.StatusAccepted, "Succesfully Delete Element")
	})

	router.Run(":8080")

	/*
		for {
			UserChoice := Homepage()
			if UserChoice == Exit {
				break
			}

			if UserChoice == ViewList {
				todoH.GetTodos(mainCtx)
			} else if UserChoice == CreateTodo {
				ctx, c := context.WithCancel(mainCtx)
				cancel = c
				todoH.InsertTodo(ctx)
			} else if UserChoice == CancelAnyInsert {
				cancel()
			}
		}*/

}
