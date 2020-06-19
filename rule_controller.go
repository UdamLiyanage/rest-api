package main

import "github.com/labstack/echo/v4"

func getRules(c echo.Context) error {
	var config = Configuration{
		Collection: ruleCollection,
	}
	rules := config.read()
	return c.JSON(200, rules)
}

func getRule(c echo.Context) error {
	var config = Configuration{
		Collection: ruleCollection,
	}
	rule := config.read()
	return c.JSON(200, rule)
}

func createRule(c echo.Context) error {
	var config = Configuration{
		Collection: ruleCollection,
	}
	rule := config.create()
	return c.JSON(201, rule)
}

func updateRule(c echo.Context) error {
	var config = Configuration{
		Collection: ruleCollection,
	}
	rule := config.read()
	return c.JSON(200, rule)
}

func deleteRule(c echo.Context) error {
	var config = Configuration{
		Collection: ruleCollection,
	}
	status := config.delete()
	if status {
		return c.JSON(401, nil)
	}
	return c.JSON(500, nil)
}
