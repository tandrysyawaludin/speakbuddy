package server

import (
	"fmt"
	"speakbuddy-be/pkg/config"
	"speakbuddy-be/pkg/handlers"

	"github.com/gin-gonic/gin"
)

var Gcfg *config.Config // global config for server

func Start(cfg *config.Config) {
	r := gin.Default()

	audioGroup := r.Group("/audio")
	{
		audioGroup.POST("/user/:user_id/phrase/:phrase_id", handlers.UploadAudio)
		audioGroup.GET("/user/:user_id/phrase/:phrase_id/:audio_format", handlers.RetrieveAudio)
	}

	r.Run(fmt.Sprintf(":%s", cfg.ServerPort))
}
