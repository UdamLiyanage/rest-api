package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
)

func getDevices(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: deviceCollection,
	}
	devices := crud.Read()
	return c.JSON(200, devices)
}

func getDevice(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: deviceCollection,
	}
	response, err := crud.Index(c.Param("id"))
	if err != nil {
		panic(err)
	}
	return c.JSON(200, response)
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
	response, err := crud.Update(c.Param("id"))
	if err != nil {
		panic(err)
	}
	return c.JSON(200, response)
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
