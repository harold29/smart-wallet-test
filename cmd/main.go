package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {

	router := gin.Default()

	viper.SetConfigFile("./pkg/common/envs/dev.yaml")

	// conf, err := config.LoadConfig()

	// if err != nil {
	// 	fmt.Printf("Error loading configuration, %s", err)
	// 	return
	// }

	// dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.Db.PostgresHost, conf.Db.PostgresUser, conf.Db.PostgresPass, conf.Db.PostgresDB, conf.Db.PostgresPort)

	// potato := db.Init(dbInfo)
	var envs map[string]string
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	issuerURL, _ := url.Parse(envs["AUTH0_ISSUER_URL"])
	audience := envs["AUTH0_AUDIENCE"]

	fmt.Printf("POTATO")
	fmt.Printf(envs["AUTH0_ISSUER_URL"])
	fmt.Printf(audience)

	provider := jwks.NewCachingProvider(issuerURL, time.Duration(5*time.Minute))

	jwtValidator, _ := validator.New(provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{audience},
	)

	jwtMiddleware := jwtmiddleware.New(jwtValidator.ValidateToken)
	// router.Use(adapter.Wrap(jwtMiddleware.CheckJWT))
	checkJwt := adapter.Wrap(jwtMiddleware.CheckJWT)

	router.GET("/ping", checkJwt, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.Run(":8080")
}
