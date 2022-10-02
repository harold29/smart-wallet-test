package main

import (
	"fmt"
	"net/http"
	"smart_wallet/pkg/common/config"
	"smart_wallet/pkg/common/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	router := gin.Default()

	viper.SetConfigFile("./pkg/common/envs/dev.yaml")

	conf, err := config.LoadConfig()

	if err != nil {
		fmt.Printf("Error loading configuration, %s", err)
		return
	}

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.Db.PostgresHost, conf.Db.PostgresUser, conf.Db.PostgresPass, conf.Db.PostgresDB, conf.Db.PostgresPort)

	potato := db.Init(dbInfo)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.Run(":8080")
}
