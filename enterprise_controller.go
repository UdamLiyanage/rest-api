package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"time"
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
		if result.Err() == mongo.ErrNoDocuments {
			return c.JSON(404, nil)
		}
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
	//enterprise.Urn = c.Get("resourceUrn").(string)
	enterprise.Urn = uuid.New().String()
	enterprise.CreatedAt = time.Now()
	enterprise.UpdatedAt = time.Now()
	_, err = crud.Create(enterprise)
	if err != nil {
		panic(err)
	}
	return c.JSON(201, enterprise)
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
	enterprise.UpdatedAt = time.Now()
	response, err := crud.Update(c.Param("id"), bson.M{
		"$set": enterprise,
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
	return c.JSON(200, enterprise)
}

func deleteEnterprise(c echo.Context) error {
	var crud Operations = Configuration{
		Collection: enterpriseCollection,
	}
	status, err := crud.Delete(c.Param("urn"))
	if err != nil || !status {
		return c.JSON(500, nil)
	}
	return c.JSON(204, nil)
}
