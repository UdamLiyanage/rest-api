package main

import (
	"encoding/json"
	"errors"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
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

func setupRouter() *echo.Echo {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var email string
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
						panic(err.Error())
					}

					t, err := jwt.ParseWithClaims(token.Raw, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
						res, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
						return res, nil
					})
					if err != nil {
						panic(err)
					}
					x := t.Claims.(jwt.MapClaims)
					email = x["https://abydub.com/email"].(string)
					println(x["https://abydub.com/email"].(string))

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
	})
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339} method=${method}, uri=${uri}, status=${status} path=${path} latency=${latency_human}\n",
	}))

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

	e.GET("/users/:id/devices", getUserDevices)
	e.GET("/users/:id/device-schemas", getUserDeviceSchemas)
	e.GET("/devices/:id/rules", getDeviceRules)
	e.GET("/device-schemas/:id/actions", getDeviceSchemaActions)
	e.GET("/actions/:id/rules", getActionRules)

	e.POST("/enterprises", createEnterprise)
	e.POST("/users", createUser)
	e.POST("/devices", createDevice)
	e.POST("/device-schemas", createDeviceSchema)
	e.POST("/actions", createAction)
	e.POST("/rules", createRule)

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
