package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
)

func getRules(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: ruleCollection,
	}
	rules := crud.Read()
	return c.JSON(200, rules)
}

func getRule(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: ruleCollection,
	}
	rule := crud.Index()
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
	response, err := crud.Update()
	if err != nil {
		panic(err)
	}
	return c.JSON(200, response)
}

func deleteRule(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: ruleCollection,
	}
	status := crud.Delete()
	if status {
		return c.JSON(401, nil)
	}
	return c.JSON(500, nil)
}
