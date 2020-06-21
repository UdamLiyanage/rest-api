package main

import (
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
	device := crud.Index()
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
	var crud Operations = Configuration{
		Collection: deviceCollection,
	}
	device := crud.Update()
	return c.JSON(200, device)
}

func deleteDevice(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: deviceCollection,
	}
	status := crud.Delete()
	if status {
		return c.JSON(401, nil)
	}
	return c.JSON(500, nil)
}
