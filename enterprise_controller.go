package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
)

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
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		panic(err)
	}
	enterprise, err := UnmarshalEnterprise(body)
	if err != nil {
		panic(err)
	}
	response, err := crud.Create(enterprise)
	if err != nil {
		panic(err)
	}
	return c.JSON(201, response)
}

func updateEnterprise(c echo.Context) error {
	var (
		crud Operations = Configuration{
			Collection: enterpriseCollection,
		}
		enterprise Enterprise
	)
	err := json.NewDecoder(c.Request().Body).Decode(&enterprise)
	if err != nil {
		panic(err)
	}
	response, err := crud.Update()
	if err != nil {
		panic(err)
	}
	return c.JSON(200, response)
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
