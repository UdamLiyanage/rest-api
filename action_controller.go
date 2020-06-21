package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
)

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
	response, err := crud.Index(c.Param("id"))
	if err != nil {
		panic(err)
	}
	return c.JSON(200, response)
}

func createAction(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: actionCollection,
	}
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		panic(err)
	}
	action, err := UnmarshalAction(body)
	if err != nil {
		panic(err)
	}
	response, err := crud.Create(action)
	if err != nil {
		panic(err)
	}
	return c.JSON(201, response)
}

func updateAction(c echo.Context) error {
	var (
		crud Operations = Configuration{
			Collection: actionCollection,
		}
		action Action
	)
	err := json.NewDecoder(c.Request().Body).Decode(&action)
	if err != nil {
		panic(err)
	}
	response, err := crud.Update(c.Param("id"))
	if err != nil {
		panic(err)
	}
	return c.JSON(200, response)
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
