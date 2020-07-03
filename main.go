package main

import (
	"bytes"
	"encoding/json"
	"errors"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"strings"
)

var (
	enterpriseCollection   *mongo.Collection
	userCollection         *mongo.Collection
	deviceCollection       *mongo.Collection
	actionCollection       *mongo.Collection
	deviceSchemaCollection *mongo.Collection
	ruleCollection         *mongo.Collection
)

type (
	Jwks struct {
		Keys []JSONWebKeys `json:"keys"`
	}

	JSONWebKeys struct {
		Kty string   `json:"kty"`
		Kid string   `json:"kid"`
		Use string   `json:"use"`
		N   string   `json:"n"`
		E   string   `json:"e"`
		X5c []string `json:"x5c"`
	}
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

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(os.Getenv("AUTH0_CERT"))

	if err != nil {
		return cert, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}

func authenticateRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userUrn string
		jM := jwtmiddleware.New(jwtmiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				aud := os.Getenv("AUTH0_AUDIENCE")
				checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
				if !checkAud {
					return token, errors.New("invalid audience")
				}
				iss := os.Getenv("AUTH0_ISSUER")
				checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
				if !checkIss {
					return token, errors.New("invalid issuer")
				}

				cert, err := getPemCert(token)
				if err != nil {
					return nil, err
				}

				t, err := jwt.ParseWithClaims(token.Raw, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
					res, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
					return res, nil
				})
				if err != nil {
					return nil, err
				}
				x := t.Claims.(jwt.MapClaims)
				sub := strings.Split(x["sub"].(string), "|")
				userUrn = sub[1]

				result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
				return result, nil
			},
			SigningMethod: jwt.SigningMethodRS256,
		})
		err := jM.CheckJWT(c.Response().Writer, c.Request())
		if err != nil {
			return echo.NewHTTPError(500, "Internal server error!")
		}
		return next(c)
	}
}

func authorizeRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authStatus := false
		requestUrl := func() string {
			basePath := os.Getenv("AUTHORIZATION_SERVER_API") + c.Path()
			log.Info(basePath + "?userUrn=" + c.Get("userUrn").(string) + "&resourceUrn=" + c.Param("id"))
			return basePath + "?userUrn=" + c.Get("userUrn").(string) + "&schemaUrn=" + c.Param("id")
		}
		createOp := func() bool {
			var responseBody map[string]string
			println("Create Op")
			body, err := json.Marshal(c.Request().Body)
			if err != nil {
				panic(err)
			}
			resp, err := http.Post(requestUrl(), "application/json", bytes.NewBuffer(body))
			if err != nil {
				log.Error(err)
				return false
			}
			defer func() {
				_ = resp.Body.Close()
			}()
			log.Info(resp.StatusCode)
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			if err != nil {
				log.Error(err)
				return false
			}
			c.Set("resourceUrn", responseBody["urn"])
			if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
				return true
			}
			return false
		}
		readOp := func() bool {
			println("Read Op")
			resp, err := http.Get(requestUrl())
			if err != nil {
				log.Error(err)
			}
			if resp.StatusCode == http.StatusOK {
				return true
			}
			return false
		}
		deleteOp := func() bool {
			println("Delete Op")
			client := &http.Client{}

			req, err := http.NewRequest("DELETE", requestUrl(), nil)
			if err != nil {
				log.Error(err)
				return false
			}

			resp, err := client.Do(req)
			if err != nil {
				log.Error(err)
				return false
			}
			if resp.StatusCode == http.StatusOK {
				return true
			}
			return false
		}
		switch c.Request().Method {
		case "GET":
			authStatus = readOp()
		case "POST":
			authStatus = createOp()
		case "PUT":
			authStatus = readOp()
		case "DELETE":
			authStatus = deleteOp()
		}

		if !authStatus {
			return echo.NewHTTPError(401, "Unauthorized")
		}
		return next(c)
	}
}

func setupRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339} method=${method}, uri=${uri}, status=${status} path=${path} latency=${latency_human}\n",
	}))

	e.POST("/auth0/users", createUser)

	/*
		API Groups. Grouped according to the handler method.
	*/
	showGroup := e.Group("/api/v1")
	showGroup.Use(authenticateRequest)
	showGroup.Use(authorizeRequest)

	indexGroup := e.Group("/api/v1")
	showGroup.Use(authenticateRequest)

	saveGroup := e.Group("/api/v1")
	saveGroup.Use(authenticateRequest)
	saveGroup.Use(authorizeRequest)

	updateGroup := e.Group("/api/v1")
	updateGroup.Use(authenticateRequest)
	updateGroup.Use(authorizeRequest)

	deleteGroup := e.Group("/api/v1")
	deleteGroup.Use(authenticateRequest)
	deleteGroup.Use(authorizeRequest)

	/*
		Routes - Prefix /api/v1
	*/
	indexGroup.GET("/enterprises", getEnterprises)
	showGroup.GET("/enterprises/:id", getEnterprise)
	indexGroup.GET("/users", getUsers)
	showGroup.GET("/users/:id", getUser)
	indexGroup.GET("/devices", getDevices)
	showGroup.GET("/devices/:id", getDevice)
	indexGroup.GET("/device-schemas", getDeviceSchemas)
	showGroup.GET("/device-schemas/:id", getDeviceSchema)
	indexGroup.GET("/actions", getActions)
	showGroup.GET("/actions/:id", getAction)
	indexGroup.GET("/rules", getRules)
	showGroup.GET("/rules/:id", getRule)

	showGroup.GET("/users/devices", getUserDevices)
	showGroup.GET("/users/device-schemas", getUserDeviceSchemas)
	showGroup.GET("/devices/:id/rules", getDeviceRules)
	showGroup.GET("/device-schemas/:id/actions", getDeviceSchemaActions)
	showGroup.GET("/actions/:id/rules", getActionRules)

	saveGroup.POST("/enterprises", createEnterprise)
	saveGroup.POST("/users", createUser)
	saveGroup.POST("/devices", createDevice)
	saveGroup.POST("/device-schemas", createDeviceSchema)
	saveGroup.POST("/actions", createAction)
	saveGroup.POST("/rules", createRule)

	updateGroup.PUT("/enterprises/:id", updateEnterprise)
	updateGroup.PUT("/users/:id", updateUser)
	updateGroup.PUT("/devices/:id", updateDevice)
	updateGroup.PUT("/device-schemas/:id", updateDeviceSchema)
	updateGroup.PUT("/actions/:id", updateAction)
	updateGroup.PUT("/rules/:id", updateRule)

	deleteGroup.DELETE("/enterprises/:id", deleteEnterprise)
	deleteGroup.DELETE("/users/:id", deleteUser)
	deleteGroup.DELETE("/devices/:id", deleteDevice)
	deleteGroup.DELETE("/device-schemas/:id", deleteDeviceSchema)
	deleteGroup.DELETE("/actions/:id", deleteAction)
	deleteGroup.DELETE("/rules/:id", deleteRule)

	return e
}

func main() {
	r := setupRouter()
	r.Logger.Fatal(r.Start(":8000"))
}
