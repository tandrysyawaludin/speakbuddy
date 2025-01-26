package main

import (
	"log"
	"speakbuddy-be/pkg/config"
	"speakbuddy-be/pkg/dao"
	"speakbuddy-be/pkg/server"
	"speakbuddy-be/pkg/speakbuddybeapp"
)

var cfg *config.Config
var err error

func init() {
	log.Print("Welcome to speakbuddybe api...")

	// get a config
	cfg, err = config.NewConfig()
	if err != nil {
		log.Fatal("Config init failed", err)
	}

	// migrate db
	if err = speakbuddybeapp.DbInit(cfg.DB); err != nil {
		log.Fatal("DB migration failed...")
	}

	dao.Init(cfg.DB)
}

func main() {
	server.Start(cfg)
}
