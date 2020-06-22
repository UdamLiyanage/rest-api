package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
)

func getRules(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: ruleCollection,
	}
	response, err := crud.Read(nil)
	if err != nil {
		panic(err)
	}
	return c.JSON(200, response)
}

func getRule(c echo.Context) error {
	var (
		crud Operations = Configuration{
			Collection: ruleCollection,
		}
		rule Rule
	)
	result := crud.Index(c.Param("id"))
	if result.Err() != nil {
		panic(result.Err())
	}
	err := result.Decode(&rule)
	if err != nil {
		panic(err)
	}
	return c.JSON(200, rule)
}

func createRule(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: ruleCollection,
	}
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		panic(err)
	}
	rule, err := UnmarshalRule(body)
	if err != nil {
		panic(err)
	}
	response, err := crud.Create(rule)
	if err != nil {
		panic(err)
	}
	return c.JSON(201, response)
}

func updateRule(c echo.Context) error {
	var (
		crud Operations = Configuration{
			Collection: ruleCollection,
		}
		rule Rule
	)
	err := json.NewDecoder(c.Request().Body).Decode(&rule)
	if err != nil {
		panic(err)
	}
	response, err := crud.Update(c.Param("id"), bson.M{
		"$set": rule,
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
	return c.JSON(200, rule)
}

func deleteRule(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: ruleCollection,
	}
	status, err := crud.Delete(c.Param("urn"))
	if err != nil || !status {
		return c.JSON(500, nil)
	}
	return c.JSON(204, nil)
}
