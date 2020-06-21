package main

import "github.com/labstack/echo/v4"

func getUsers(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: userCollection,
	}
	users := crud.Read()
	return c.JSON(200, users)
}

func getUser(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: userCollection,
	}
	user := crud.Index()
	return c.JSON(200, user)
}

func createUser(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: userCollection,
	}
	user := crud.Create()
	return c.JSON(201, user)
}

func updateUser(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: userCollection,
	}
	user := crud.Update()
	return c.JSON(200, user)
}

func deleteUser(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: userCollection,
	}
	status := crud.Delete()
	if status {
		return c.JSON(401, nil)
	}
	return c.JSON(500, nil)
}
