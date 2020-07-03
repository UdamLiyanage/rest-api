package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"time"
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
		if result.Err() == mongo.ErrNoDocuments {
			return c.JSON(404, nil)
		}
		panic(result.Err())
	}
	err := result.Decode(&device)
	if err != nil {
		panic(err)
	}
	return c.JSON(200, device)
}

func getDeviceRules(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: ruleCollection,
	}
	response, err := crud.Read(bson.M{
		"device": c.Param("id"),
	})
	if err != nil {
		return c.JSON(500, err)
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
	//device.Urn = c.Get("resourceUrn").(string)
	//device.User = c.Get("userUrn").(string)
	device.Urn = uuid.New().String()
	device.CreatedAt = time.Now()
	device.UpdatedAt = time.Now()
	_, err = crud.Create(device)
	if err != nil {
		panic(err)
	}
	return c.JSON(201, device)
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
	device.UpdatedAt = time.Now()
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
		return c.JSON(500, nil)
	}
	return c.JSON(204, nil)
}
