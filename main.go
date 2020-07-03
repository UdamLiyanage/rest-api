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

/*func a(c echo.Context, userUrn string) bool {
	if (c.Param("id") == "" && c.Request().Method != "POST") || len(strings.Split(c.Path(), "/")) > 5 {
		//Index method or User specific show
		return true
	}
	c.Set("userUrn", userUrn)
	requestUrl := func() string {
		basePath := os.Getenv("AUTHORIZATION_SERVER_API") + c.Path()
		log.Info(basePath + "?userUrn=" + userUrn + "&resourceUrn=" + c.Param("id"))
		return basePath + "?userUrn=" + userUrn + "&schemaUrn=" + c.Param("id")
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
		return readOp()
	case "POST":
		return createOp()
	case "PUT":
		return readOp()
	case "DELETE":
		return deleteOp()
	}
	return false
}*/

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

	authenticatedGroup := e.Group("/api/v1")

	authenticatedGroup.Use(authenticateRequest)

	authenticatedGroup.GET("/enterprises", getEnterprises)
	authenticatedGroup.GET("/enterprises/:id", getEnterprise)
	authenticatedGroup.GET("/users", getUsers)
	authenticatedGroup.GET("/users/:id", getUser)
	authenticatedGroup.GET("/devices", getDevices)
	authenticatedGroup.GET("/devices/:id", getDevice)
	authenticatedGroup.GET("/device-schemas", getDeviceSchemas)
	authenticatedGroup.GET("/device-schemas/:id", getDeviceSchema)
	authenticatedGroup.GET("/actions", getActions)
	authenticatedGroup.GET("/actions/:id", getAction)
	authenticatedGroup.GET("/rules", getRules)
	authenticatedGroup.GET("/rules/:id", getRule)

	authenticatedGroup.GET("/users/devices", getUserDevices)
	authenticatedGroup.GET("/users/device-schemas", getUserDeviceSchemas)
	authenticatedGroup.GET("/devices/:id/rules", getDeviceRules)
	authenticatedGroup.GET("/device-schemas/:id/actions", getDeviceSchemaActions)
	authenticatedGroup.GET("/actions/:id/rules", getActionRules)

	authenticatedGroup.POST("/enterprises", createEnterprise)
	authenticatedGroup.POST("/users", createUser)
	authenticatedGroup.POST("/devices", createDevice)
	authenticatedGroup.POST("/device-schemas", createDeviceSchema)
	authenticatedGroup.POST("/actions", createAction)
	authenticatedGroup.POST("/rules", createRule)

	authenticatedGroup.PUT("/enterprises/:id", updateEnterprise)
	authenticatedGroup.PUT("/users/:id", updateUser)
	authenticatedGroup.PUT("/devices/:id", updateDevice)
	authenticatedGroup.PUT("/device-schemas/:id", updateDeviceSchema)
	authenticatedGroup.PUT("/actions/:id", updateAction)
	authenticatedGroup.PUT("/rules/:id", updateRule)

	authenticatedGroup.DELETE("/enterprises/:id", deleteEnterprise)
	authenticatedGroup.DELETE("/users/:id", deleteUser)
	authenticatedGroup.DELETE("/devices/:id", deleteDevice)
	authenticatedGroup.DELETE("/device-schemas/:id", deleteDeviceSchema)
	authenticatedGroup.DELETE("/actions/:id", deleteAction)
	authenticatedGroup.DELETE("/rules/:id", deleteRule)

	return e
}

func main() {
	r := setupRouter()
	r.Logger.Fatal(r.Start(":8000"))
}
