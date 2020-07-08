package main

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"time"
)

func getUsers(c echo.Context) error {
	var (
		crud Operations = Configuration{
			Collection: userCollection,
		}
		users []User
	)
	response, err := crud.Read(nil)
	if err != nil {
		panic(err)
	}
	err = response.All(context.Background(), &users)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, users)
}

func getUser(c echo.Context) error {
	var (
		crud Operations = Configuration{
			Collection: userCollection,
		}
		user User
	)
	result := crud.Index(c.Param("id"))
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return c.JSON(404, nil)
		}
		panic(result.Err())
	}
	err := result.Decode(&user)
	if err != nil {
		panic(err)
	}
	return c.JSON(200, user)
}

func getUserDevices(c echo.Context) error {
	var (
		crud Operations = Configuration{
			Collection: deviceCollection,
		}
		devices []Device
	)
	response, err := crud.Read(bson.M{
		"user": c.Get("userUrn").(string),
	})
	if err != nil {
		return c.JSON(500, err)
	}
	err = response.All(context.Background(), &devices)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, devices)
}

func getUserDeviceSchemas(c echo.Context) error {
	var (
		crud Operations = Configuration{
			Collection: deviceSchemaCollection,
		}
		schemas []DeviceSchema
	)
	response, err := crud.Read(bson.M{
		"user": c.Get("userUrn").(string),
	})
	if err != nil {
		return c.JSON(500, err)
	}
	err = response.All(context.Background(), &schemas)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, schemas)
}

func createUser(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: userCollection,
	}
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		panic(err)
	}
	user, err := UnmarshalUser(body)
	if err != nil {
		panic(err)
	}
	//user.Urn = c.Get("resourceUrn").(string)
	user.Urn = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err = crud.Create(user)
	if err != nil {
		panic(err)
	}
	return c.JSON(201, user)
}

func updateUser(c echo.Context) error {
	var (
		crud Operations = Configuration{
			Collection: userCollection,
		}
		user User
	)
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	user.UpdatedAt = time.Now()
	response, err := crud.Update(c.Param("id"), bson.M{
		"$set": user,
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
	return c.JSON(200, user)
}

func deleteUser(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: userCollection,
	}
	status, err := crud.Delete(c.Param("urn"))
	if err != nil || !status {
		return c.JSON(500, nil)
	}
	return c.JSON(204, nil)
}
