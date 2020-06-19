package main

import "github.com/labstack/echo/v4"

func getUsers(c echo.Context) error {
	var config = Configuration{
		Collection: userCollection,
	}
	users := config.read()
	return c.JSON(200, users)
}

func getUser(c echo.Context) error {
	var config = Configuration{
		Collection: userCollection,
	}
	user := config.read()
	return c.JSON(200, user)
}

func createUser(c echo.Context) error {
	var config = Configuration{
		Collection: userCollection,
	}
	user := config.create()
	return c.JSON(201, user)
}

func updateUser(c echo.Context) error {
	var config = Configuration{
		Collection: userCollection,
	}
	user := config.read()
	return c.JSON(200, user)
}

func deleteUser(c echo.Context) error {
	var config = Configuration{
		Collection: userCollection,
	}
	status := config.delete()
	if status {
		return c.JSON(401, nil)
	}
	return c.JSON(500, nil)
}
