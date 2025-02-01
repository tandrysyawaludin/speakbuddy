package config

import (
	"fmt"
	"os"
	"speakbuddy-be/pkg/db"

	"gorm.io/gorm"
)

type Config struct {
	DbHost     string `yaml:"DbHost"`
	DbName     string `yaml:"DbName"`
	DbPass     string `yaml:"DbPass"`
	DbUser     string `yaml:"DbUser"`
	SftpHost   string `yaml:"SftpHost"`
	SftpPort   string `yaml:"SftpPort"`
	SftpPass   string `yaml:"SftpPass"`
	SftpUser   string `yaml:"SftpUser"`
	DB         *gorm.DB
	ServerPort string `yaml:"ServerPort"`
}

var (
	dbHost     = "localhost:3306"
	dbName     = "evergreen_speakbuddybe_db"
	dbPass     = "password"
	dbUser     = "root"
	serverPort = "8081"
	sftpHost   = ""
	sftpPort   = ""
	sftpPass   = ""
	sftpUser   = ""
)

func NewConfig() (*Config, error) {
	var err error
	cfg := &Config{
		DbHost:     dbHost,
		DbName:     dbName,
		DbPass:     dbPass,
		DbUser:     dbUser,
		ServerPort: serverPort,
		SftpHost:   sftpHost,
		SftpPort:   sftpPort,
		SftpUser:   sftpUser,
		SftpPass:   sftpPass,
	}

	// update config values from env, if any
	cfg.GETENVs()

	// init db conn
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbName)
	cfg.DB = db.InitDB(dsn)

	return cfg, err
}

func (c *Config) GETENVs() {
	if val, found := os.LookupEnv("CONFIG_DBHOST"); found {
		c.DbHost = val
	}

	if val, found := os.LookupEnv("CONFIG_DBNAME"); found {
		c.DbName = val
	}

	if val, found := os.LookupEnv("CONFIG_DBPASS"); found {
		c.DbPass = val
	}

	if val, found := os.LookupEnv("CONFIG_DBUSER"); found {
		c.DbUser = val
	}

	if val, found := os.LookupEnv("CONFIG_SERVER_PORT"); found {
		c.ServerPort = val
	}

	if val, found := os.LookupEnv("CONFIG_SFTPHOST"); found {
		c.SftpHost = val
	}

	if val, found := os.LookupEnv("CONFIG_SFTPPORT"); found {
		c.SftpPort = val
	}

	if val, found := os.LookupEnv("CONFIG_SFTPPASS"); found {
		c.SftpPass = val
	}

	if val, found := os.LookupEnv("CONFIG_SFTPUSER"); found {
		c.SftpUser = val
	}
}
