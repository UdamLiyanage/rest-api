package main

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"time"
)

func getRules(c echo.Context) error {
	var (
		crud Operations = Configuration{
			Collection: ruleCollection,
		}
		rules []Rule
	)
	response, err := crud.Read(nil)
	if err != nil {
		panic(err)
	}
	err = response.All(context.Background(), &rules)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, rules)
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
		if result.Err() == mongo.ErrNoDocuments {
			return c.JSON(404, nil)
		}
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
	//rule.Urn = c.Get("resourceUrn").(string)
	rule.Urn = uuid.New().String()
	rule.CreatedAt = time.Now()
	rule.UpdatedAt = time.Now()
	_, err = crud.Create(rule)
	if err != nil {
		panic(err)
	}
	return c.JSON(201, rule)
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
	rule.UpdatedAt = time.Now()
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
