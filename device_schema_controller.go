package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"time"
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
		if result.Err() == mongo.ErrNoDocuments {
			return c.JSON(404, nil)
		}
		panic(result.Err())
	}
	err := result.Decode(&schema)
	if err != nil {
		panic(err)
	}
	return c.JSON(200, schema)
}

func getDeviceSchemaActions(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: actionCollection,
	}
	response, err := crud.Read(bson.M{
		"schema": c.Param("id"),
	})
	if err != nil {
		return c.JSON(500, err)
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
	schema.Urn = c.Get("resourceUrn").(string)
	schema.CreatedAt = time.Now()
	schema.UpdatedAt = time.Now()
	_, err = crud.Create(schema)
	if err != nil {
		panic(err)
	}
	return c.JSON(201, schema)
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
	schema.UpdatedAt = time.Now()
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
