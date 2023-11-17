package config

import (
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Constants struct {
	Port             string `yaml:"port"`
	JWTSecret        string `yaml:"jwtSecret"`
	ConnectionString string `yaml:"connectionString"`
	DataPath         string `yaml:"dataPath"`
	BaseURL          string `yaml:"baseURL"`
}

type Config struct {
	Constants
	Database *gorm.DB
}

func initViper() (Constants, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok && err != nil {
		return Constants{}, err
	}

	// At this point the only error would be a missing file
	if err != nil {
		err = initViperEnv()

		if err != nil {
			return Constants{}, err
		}
	}

	var constants Constants
	err = viper.Unmarshal(&constants)
	return constants, err
}

func initViperEnv() error {
	if !doesEnvExists("JWT_SECRET") ||
		!doesEnvExists("CONNECTION_STRING") ||
		!doesEnvExists("DATA_PATH") ||
		!doesEnvExists("BASE_URL") {
		return &MissingEnvVariableError{}
	}

	viper.SetDefault("Port", os.Getenv("PORT"))
	viper.SetDefault("JWTSecret", os.Getenv("JWT_SECRET"))
	viper.SetDefault("ConnectionString", os.Getenv("CONNECTION_STRING"))
	viper.SetDefault("DatePath", os.Getenv("DATA_PATH"))
	viper.SetDefault("BaseURL", os.Getenv("BASE_URL"))

	return nil
}

func doesEnvExists(name string) bool {
	_, exists := os.LookupEnv(name)

	return exists
}
func New() (*Config, error) {
	// Configuration
	config := Config{}
	constants, err := initViper()

	config.Constants = constants
	if err != nil {
		return &config, err
	}

	// Database
	databaseSession, err := gorm.Open(mysql.Open(config.Constants.ConnectionString), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	config.Database = databaseSession

	return &config, nil
}
