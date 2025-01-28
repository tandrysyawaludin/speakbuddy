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
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Print("welcome to speakbuddybe api")

	// get a config
	cfg, err = config.NewConfig()
	if err != nil {
		log.Printf("[ERROR] config init failed, err: %+v", err)
	}

	// migrate db
	if err = speakbuddybeapp.DbInit(cfg.DB); err != nil {
		log.Printf("[ERROR] db migration failed, err: %+v", err)
	}

	dao.Init(cfg.DB)
}

func main() {
	server.Start(cfg)
}
