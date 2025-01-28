package server

import (
	"fmt"
	"log"
	"net/http"
	"speakbuddy-be/pkg/config"
	"speakbuddy-be/pkg/handlers"
	"time"

	"github.com/gin-contrib/requestid"

	"github.com/gin-gonic/gin"
)

var Gcfg *config.Config // global config for server

func Start(cfg *config.Config) {
	r := gin.Default()
	r.Use(requestid.New())

	audioGroup := r.Group("/audio")
	{
		audioGroup.POST("/user/:user_id/phrase/:phrase_id", handlers.UploadAudio)
		audioGroup.GET("/user/:user_id/phrase/:phrase_id/:audio_format", handlers.RetrieveAudio)
	}
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	  })

	log.Printf("[INFO] server start with port: %s", cfg.ServerPort)

	r.Run(fmt.Sprintf(":%s", cfg.ServerPort))
}
