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
	response, err := crud.Read(nil)
	if err != nil {
		panic(err)
	}
	return c.JSON(200, response)
}

func getEnterprise(c echo.Context) error {
	var (
		crud Operations = Configuration{
			Collection: enterpriseCollection,
		}
		ent Enterprise
	)
	result := crud.Index(c.Param("id"))
	if result.Err() != nil {
		panic(result.Err())
	}
	err := result.Decode(&ent)
	if err != nil {
		panic(err)
	}
	return c.JSON(200, ent)
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
	response, err := crud.Update(c.Param("id"))
	if err != nil {
		panic(err)
	}
	return c.JSON(200, response)
}

func deleteEnterprise(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: enterpriseCollection,
	}
	status, err := crud.Delete(c.Param("urn"))
	if err != nil || !status {
		return c.JSON(401, nil)
	}
	return c.JSON(500, nil)
}
