package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
)

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
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		panic(err)
	}
	schema, err := UnmarshalDeviceSchema(body)
	if err != nil {
		panic(err)
	}
	response, err := crud.Create(schema)
	if err != nil {
		panic(err)
	}
	return c.JSON(201, response)
}

func updateDeviceSchema(c echo.Context) error {
	var (
		crud Operations = Configuration{
			Collection: deviceSchemaCollection,
		}
		schema DeviceSchema
	)
	err := json.NewDecoder(c.Request().Body).Decode(&schema)
	if err != nil {
		panic(err)
	}
	response := crud.Update()
	return c.JSON(200, response)
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
