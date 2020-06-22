package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
)

func getDevices(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: deviceCollection,
	}
	response, err := crud.Read(nil)
	if err != nil {
		panic(err)
	}
	return c.JSON(200, response)
}

func getDevice(c echo.Context) error {
	var (
		crud Operations = Configuration{
			Collection: deviceCollection,
		}
		device Device
	)
	result := crud.Index(c.Param("id"))
	if result.Err() != nil {
		panic(result.Err())
	}
	err := result.Decode(&device)
	if err != nil {
		panic(err)
	}
	return c.JSON(200, device)
}

func createDevice(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: deviceCollection,
	}
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		panic(err)
	}
	device, err := UnmarshalDevice(body)
	if err != nil {
		panic(err)
	}
	response, err := crud.Create(device)
	if err != nil {
		panic(err)
	}
	return c.JSON(201, response)
}

func updateDevice(c echo.Context) error {
	var (
		crud Operations = Configuration{
			Collection: deviceCollection,
		}
		device Device
	)
	err := json.NewDecoder(c.Request().Body).Decode(&device)
	if err != nil {
		panic(err)
	}
	response, err := crud.Update(c.Param("id"), bson.M{
		"$set": device,
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
	return c.JSON(200, device)
}

func deleteDevice(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: deviceCollection,
	}
	status, err := crud.Delete(c.Param("urn"))
	if err != nil || !status {
		return c.JSON(401, nil)
	}
	return c.JSON(500, nil)
}
