package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
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
	var (
		crud Operations = Configuration{
			Collection: deviceSchemaCollection,
		}
		schema DeviceSchema
	)
	result := crud.Index(c.Param("id"))
	if result.Err() != nil {
		panic(result.Err())
	}
	err := result.Decode(&schema)
	if err != nil {
		panic(err)
	}
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
	response, err := crud.Update(c.Param("id"), bson.M{
		"$set": schema,
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
	return c.JSON(200, schema)
}

func deleteDeviceSchema(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: deviceSchemaCollection,
	}
	status, err := crud.Delete(c.Param("urn"))
	if err != nil || !status {
		return c.JSON(500, nil)
	}
	return c.JSON(204, nil)
}
