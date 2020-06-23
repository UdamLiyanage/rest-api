package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"time"
)

func getActions(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: actionCollection,
	}
	response, err := crud.Read(nil)
	if err != nil {
		panic(err)
	}
	return c.JSON(200, response)
}

func getAction(c echo.Context) error {
	var (
		crud Operations = Configuration{
			Collection: actionCollection,
		}
		action Action
	)
	result := crud.Index(c.Param("id"))
	if result.Err() != nil {
		panic(result.Err())
	}
	err := result.Decode(&action)
	if err != nil {
		panic(err)
	}
	return c.JSON(200, action)
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
	action.CreatedAt = time.Now()
	action.UpdatedAt = time.Now()
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
	action.UpdatedAt = time.Now()
	response, err := crud.Update(c.Param("id"), bson.M{
		"$set": action,
	})
	if err != nil {
		panic(err)
	}
	if response.MatchedCount == 0 {
		return c.JSON(404, "document with urn not found")
	} else {
		if response.ModifiedCount == 0 {
			return c.JSON(304, "document not modified")
		}
	}
	return c.JSON(200, action)
}

func deleteAction(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: actionCollection,
	}
	status, err := crud.Delete(c.Param("urn"))
	if err != nil || !status {
		return c.JSON(500, nil)
	}
	return c.JSON(204, nil)
}
