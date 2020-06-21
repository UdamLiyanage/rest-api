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
	response, err := crud.Read(nil)
	if err != nil {
		panic(err)
	}
	return c.JSON(200, response)
}

func getDeviceSchema(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: deviceSchemaCollection,
	}
	response, err := crud.Index(c.Param("id"))
	if err != nil {
		panic(err)
	}
	return c.JSON(200, response)
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
	response, err := crud.Update(c.Param("id"))
	if err != nil {
		panic(err)
	}
	return c.JSON(200, response)
}

func deleteDeviceSchema(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: deviceSchemaCollection,
	}
	status, err := crud.Delete(c.Param("urn"))
	if err != nil || !status {
		return c.JSON(401, nil)
	}
	return c.JSON(500, nil)
}
