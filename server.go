package main

import (
	ctrl "code/ctrl"
	"code/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	db := db.Connect()

	println(db)

	e.POST("/users", ctrl.CreateUser)
	e.GET("/users", ctrl.ListUsers)
	e.GET("/users/:id", ctrl.GetUser)
	e.PUT("/users/:id", ctrl.UpdateUser)
	e.DELETE("/users/:id", ctrl.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
