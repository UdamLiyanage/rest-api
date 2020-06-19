package main

import "github.com/labstack/echo/v4"

func getActions(c echo.Context) error {
	var config = Configuration{
		Collection: actionCollection,
	}
	actions := config.read()
	return c.JSON(200, actions)
}

func getAction(c echo.Context) error {
	var config = Configuration{
		Collection: actionCollection,
	}
	action := config.read()
	return c.JSON(200, action)
}

func createAction(c echo.Context) error {
	var config = Configuration{
		Collection: actionCollection,
	}
	action := config.create()
	return c.JSON(201, action)
}

func updateAction(c echo.Context) error {
	var config = Configuration{
		Collection: actionCollection,
	}
	action := config.update()
	return c.JSON(200, action)
}

func deleteAction(c echo.Context) error {
	var config = Configuration{
		Collection: actionCollection,
	}
	status := config.delete()
	if status {
		return c.JSON(401, nil)
	}
	return c.JSON(500, nil)
}
