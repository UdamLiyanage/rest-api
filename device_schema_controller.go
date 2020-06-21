package main

import "github.com/labstack/echo/v4"

func getDeviceSchemas(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: deviceSchemaCollection,
	}
	schemas := crud.Read()
	return c.JSON(200, schemas)
}

func getDeviceSchema(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: deviceSchemaCollection,
	}
	schema := crud.Index()
	return c.JSON(200, schema)
}

func createDeviceSchema(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: deviceSchemaCollection,
	}
	schema := crud.Create()
	return c.JSON(201, schema)
}

func updateDeviceSchema(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: deviceSchemaCollection,
	}
	schema := crud.Update()
	return c.JSON(200, schema)
}

func deleteDeviceSchema(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: deviceSchemaCollection,
	}
	status := crud.Delete()
	if status {
		return c.JSON(401, nil)
	}
	return c.JSON(500, nil)
}
