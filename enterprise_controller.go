package main

import "github.com/labstack/echo/v4"

func getEnterprises(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: enterpriseCollection,
	}
	enterprises := crud.Read()
	return c.JSON(200, enterprises)
}

func getEnterprise(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: enterpriseCollection,
	}
	enterprise := crud.Index()
	return c.JSON(200, enterprise)
}

func createEnterprise(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: enterpriseCollection,
	}
	enterprise := crud.Create()
	return c.JSON(201, enterprise)
}

func updateEnterprise(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: enterpriseCollection,
	}
	enterprise := crud.Update()
	return c.JSON(200, enterprise)
}

func deleteEnterprise(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: enterpriseCollection,
	}
	status := crud.Delete()
	if status {
		return c.JSON(401, nil)
	}
	return c.JSON(500, nil)
}
