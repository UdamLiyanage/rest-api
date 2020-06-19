package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setupRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339} method=${method}, uri=${uri}, status=${status} path=${path} latency=${latency_human}\n",
	}))

	e.Use(middleware.Recover())

	e.GET("/enterprises", getEnterprises)
	e.GET("/enterprises/:id", getEnterprise)
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)

	e.POST("/enterprises", createEnterprise)
	e.POST("/users", createUser)

	e.PUT("/enterprises/:id", updateEnterprise)
	e.PUT("/users/:id", updateUser)

	e.DELETE("/enterprises/:id", deleteEnterprise)
	e.DELETE("/users/:id", deleteUser)

	return e
}

func main() {
	r := setupRouter()
	r.Logger.Fatal(r.Start(":8000"))
}
