package configs

import (
	"bohemian/practice/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	environmentPath = "configs/environment"
	configFileName  = "dev"
	configType      = "yaml"
)

var (
	databaseConnection *gorm.DB
)

func InitializeConfig(engine *gin.Engine) {
	initializeViper()

	engine.Use(middleware.Logger())

	err := engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Panicf("failed to set trusted proxies. (err: %+v)", err)
	}

	initializeDatabase()
}

func initializeViper() {
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(environmentPath)
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Panicf("Config file not found; ignore error if desired. (err: %+v)", err)
		} else {
			log.Panicf("Config file was found but another error was produced. (err: %+v)", err)
		}
	}
}

func initializeDatabase() {
	var host = viper.GetString("database.host")
	var username = viper.GetString("database.username")
	var password = viper.GetString("database.password")
	var databaseName = viper.GetString("database.name")

	dsn := username + ":" + password + "@tcp(" + host + ")/" + databaseName + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("Database initialize fail. (err: %+v)", err)
	}

	databaseConnection = db
}
