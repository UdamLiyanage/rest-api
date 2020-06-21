package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
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
	response, err := crud.Update(c.Param("id"))
	if err != nil {
		panic(err)
	}
	return c.JSON(200, response)
}

func deleteUser(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: userCollection,
	}
	status, err := crud.Delete(c.Param("urn"))
	if err != nil || !status {
		return c.JSON(401, nil)
	}
	return c.JSON(500, nil)
}
