package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DBConfig struct {
	PostgresHost            string `mapstructure:"PG_HOST"`
	PostgresHostReplica     string `mapstructure:"PG_HOST_REPLICA"`
	PostgresUser            string `mapstructure:"PG_USER"`
	PostgresDB              string `mapstructure:"PG_DB"`
	PostgresPass            string `mapstructure:"PG_PASS"`
	PostgresPort            string `mapstructure:"PG_PORT"`
	PostgresPortReplica     string `mapstructure:"PG_PORT_REPLICA"`
	PostgresConnMaxIdleTime string `mapstructure:"PG_CONN_MAX_IDLE_TIME"`
	PostgresConnMaxLifeTime string `mapstructure:"PG_CONN_MAX_LIFE_TIME"`
	PostgresMaxIdleConn     string `mapstructure:"PG_MAX_IDLE_CONN"`
	PostgresMaxOpenConn     string `mapstructure:"PG_MAX_OPEN_CONN"`
}

type Config struct {
	Db DBConfig `mapstructure:"DB"`
}

func LoadConfig() (c Config, e error) {
	viper.AddConfigPath("pkg/common/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return
	}

	viper.SetDefault("POSTGRES_HOST", "127.0.0.1")

	err = viper.Unmarshal(&c)

	return
}
