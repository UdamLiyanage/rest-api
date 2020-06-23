package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"time"
)

func getUsers(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: userCollection,
	}
	response, err := crud.Read(nil)
	if err != nil {
		panic(err)
	}
	return c.JSON(200, response)
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
		panic(result.Err())
	}
	err := result.Decode(&user)
	if err != nil {
		panic(err)
	}
	return c.JSON(200, user)
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
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	response, err := crud.Create(user)
	if err != nil {
		panic(err)
	}
	return c.JSON(201, response)
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
