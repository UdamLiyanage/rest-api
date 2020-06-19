package main

import "github.com/labstack/echo/v4"

func getDeviceSchemas(c echo.Context) error {
	var config = Configuration{
		Collection: deviceSchemaCollection,
	}
	schemas := config.read()
	return c.JSON(200, schemas)
}

func getDeviceSchema(c echo.Context) error {
	var config = Configuration{
		Collection: deviceSchemaCollection,
	}
	schema := config.read()
	return c.JSON(200, schema)
}

func createDeviceSchema(c echo.Context) error {
	var config = Configuration{
		Collection: deviceSchemaCollection,
	}
	schema := config.create()
	return c.JSON(201, schema)
}

func updateDeviceSchema(c echo.Context) error {
	var config = Configuration{
		Collection: deviceSchemaCollection,
	}
	schema := config.update()
	return c.JSON(200, schema)
}

func deleteDeviceSchema(c echo.Context) error {
	var config = Configuration{
		Collection: deviceSchemaCollection,
	}
	status := config.delete()
	if status {
		return c.JSON(401, nil)
	}
	return c.JSON(500, nil)
}
