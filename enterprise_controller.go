package main

import "github.com/labstack/echo/v4"

func getEnterprises(c echo.Context) error {
	var config = Configuration{
		Collection: enterpriseCollection,
	}
	enterprises := config.read()
	return c.JSON(200, enterprises)
}

func getEnterprise(c echo.Context) error {
	var config = Configuration{
		Collection: enterpriseCollection,
	}
	enterprise := config.read()
	return c.JSON(200, enterprise)
}

func createEnterprise(c echo.Context) error {
	var config = Configuration{
		Collection: enterpriseCollection,
	}
	enterprise := config.create()
	return c.JSON(201, enterprise)
}

func updateEnterprise(c echo.Context) error {
	var config = Configuration{
		Collection: enterpriseCollection,
	}
	enterprise := config.read()
	return c.JSON(200, enterprise)
}

func deleteEnterprise(c echo.Context) error {
	var config = Configuration{
		Collection: enterpriseCollection,
	}
	status := config.delete()
	if status {
		return c.JSON(401, nil)
	}
	return c.JSON(500, nil)
}
