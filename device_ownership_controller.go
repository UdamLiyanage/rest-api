package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func attachDeviceOwner(c echo.Context) error {
	var (
		crud Operations = Configuration{
			Collection: deviceCollection,
		}
		deviceOwnership DeviceOwnership
	)
	err := json.NewDecoder(c.Request().Body).Decode(&deviceOwnership)
	if err != nil {
		panic(err)
	}
	response, err := crud.Update(c.Param("urn"), bson.M{
		"$set": bson.M{
			"owner_urn": deviceOwnership.OwnerUrn,
		},
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
	return c.JSON(200, deviceOwnership)
}

func attachDeviceEnterprise(c echo.Context) error {
	var (
		crud Operations = Configuration{
			Collection: deviceCollection,
		}
		deviceOwnership DeviceOwnership
	)
	err := json.NewDecoder(c.Request().Body).Decode(&deviceOwnership)
	if err != nil {
		panic(err)
	}
	response, err := crud.Update(c.Param("urn"), bson.M{
		"$set": bson.M{
			"enterprise_urn": deviceOwnership.OwnerUrn,
		},
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
	return c.JSON(200, deviceOwnership)
}
