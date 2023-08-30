package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

var cfg *Config

func Init(path string) {
	if path == "" {
		path = "config"
	}
	initViper(path)
	LoadConfig()
}

func initViper(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("connot read config file: %s", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func LoadConfig() *Config {
	if cfg == nil {
		cfg = &Config{
			Mysql: Mysql{
				Host:         viper.GetString("mysql.host"),
				Port:         viper.GetInt("mysql.port"),
				Database:     viper.GetString("mysql.database"),
				Timeout:      viper.GetDuration("mysql.timeout"),
				MaxIdleConns: viper.GetInt("mysql.max.idleconns"),
				MaxOpenConns: viper.GetInt("mysql.max.openconns"),
				MaxLifetime:  viper.GetDuration("mysql.max.lifetime"),
			},
			Secret: Secret{
				MysqlUsername: viper.GetString("secret.mysqlUsername"),
				MysqlPassword: viper.GetString("secret.mysqlPassword"),
			},
		}
	}
	return cfg
}
