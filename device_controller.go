package main

import "github.com/labstack/echo/v4"

func getDevices(c echo.Context) error {
	var config = Configuration{
		Collection: deviceCollection,
	}
	devices := config.read()
	return c.JSON(200, devices)
}

func getDevice(c echo.Context) error {
	var config = Configuration{
		Collection: deviceCollection,
	}
	device := config.read()
	return c.JSON(200, device)
}

func createDevice(c echo.Context) error {
	var config = Configuration{
		Collection: deviceCollection,
	}
	device := config.create()
	return c.JSON(201, device)
}

func updateDevice(c echo.Context) error {
	var config = Configuration{
		Collection: deviceCollection,
	}
	device := config.update()
	return c.JSON(200, device)
}

func deleteDevice(c echo.Context) error {
	var config = Configuration{
		Collection: deviceCollection,
	}
	status := config.delete()
	if status {
		return c.JSON(401, nil)
	}
	return c.JSON(500, nil)
}
