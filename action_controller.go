package main

import "github.com/labstack/echo/v4"

func getActions(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: actionCollection,
	}
	actions := crud.Read()
	return c.JSON(200, actions)
}

func getAction(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: actionCollection,
	}
	action := crud.Index()
	return c.JSON(200, action)
}

func createAction(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: actionCollection,
	}
	action := crud.Create()
	return c.JSON(201, action)
}

func updateAction(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: actionCollection,
	}
	action := crud.Update()
	return c.JSON(200, action)
}

func deleteAction(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: actionCollection,
	}
	status := crud.Delete()
	if status {
		return c.JSON(401, nil)
	}
	return c.JSON(500, nil)
}
