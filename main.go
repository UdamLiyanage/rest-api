package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

var (
	enterpriseCollection   *mongo.Collection
	userCollection         *mongo.Collection
	deviceCollection       *mongo.Collection
	actionCollection       *mongo.Collection
	deviceSchemaCollection *mongo.Collection
	ruleCollection         *mongo.Collection
)

func init() {
	databaseClient := connect()
	enterpriseCollection = databaseCollection(databaseClient, os.Getenv("ENTERPRISES_DB"), os.Getenv("ENTERPRISES_COLLECTION"))
	userCollection = databaseCollection(databaseClient, os.Getenv("USERS_DB"), os.Getenv("USERS_COLLECTION"))
	deviceSchemaCollection = databaseCollection(databaseClient, os.Getenv("DEVICE_SCHEMAS_DB"), os.Getenv("DEVICE_SCHEMAS_COLLECTION"))
	deviceCollection = databaseCollection(databaseClient, os.Getenv("DEVICES_DB"), os.Getenv("DEVICES_COLLECTION"))
	actionCollection = databaseCollection(databaseClient, os.Getenv("ACTIONS_DB"), os.Getenv("ACTIONS_COLLECTION"))
	ruleCollection = databaseCollection(databaseClient, os.Getenv("RULES_DB"), os.Getenv("RULES_COLLECTION"))
}

func setupRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339} method=${method}, uri=${uri}, status=${status} path=${path} latency=${latency_human}\n",
	}))

	e.Use(middleware.Recover())

	e.GET("/enterprises", getEnterprises)
	e.GET("/enterprises/:id", getEnterprise)
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)
	e.GET("/devices", getDevices)
	e.GET("/devices/:id", getDevice)
	e.GET("/device-schemas", getDeviceSchemas)
	e.GET("/device-schemas/:id", getDeviceSchema)
	e.GET("/actions", getActions)
	e.GET("/actions/:id", getAction)
	e.GET("/rules", getRules)
	e.GET("/rules/:id", getRule)

	e.POST("/enterprises", createEnterprise)
	e.POST("/users", createUser)
	e.POST("/devices", createDevice)
	e.POST("/device-schemas", createDeviceSchema)
	e.POST("/actions", createAction)
	e.POST("/rules", createRule)
	e.POST("/devices/:id/attach/user", attachDeviceOwner)

	e.PUT("/enterprises/:id", updateEnterprise)
	e.PUT("/users/:id", updateUser)
	e.PUT("/devices/:id", updateDevice)
	e.PUT("/device-schemas/:id", updateDeviceSchema)
	e.PUT("/actions/:id", updateAction)
	e.PUT("/rules/:id", updateRule)

	e.DELETE("/enterprises/:id", deleteEnterprise)
	e.DELETE("/users/:id", deleteUser)
	e.DELETE("/devices/:id", deleteDevice)
	e.DELETE("/device-schemas/:id", deleteDeviceSchema)
	e.DELETE("/actions/:id", deleteAction)
	e.DELETE("/rules/:id", deleteRule)

	return e
}

func main() {
	r := setupRouter()
	r.Logger.Fatal(r.Start(":8000"))
}
