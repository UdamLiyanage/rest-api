package main

import "github.com/labstack/echo/v4"

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
	rule := crud.Create()
	return c.JSON(201, rule)
}

func updateRule(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: ruleCollection,
	}
	rule := crud.Update()
	return c.JSON(200, rule)
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
