package handlers

import (
	"fmt"
	"log"
	"net/http"
	"speakbuddy-be/pkg/dao"
	"speakbuddy-be/pkg/dto"
	"speakbuddy-be/pkg/utils"
	"strconv"

	"github.com/gin-contrib/requestid"

	"github.com/gin-gonic/gin"
)

func RetrieveAudio(c *gin.Context) {
	audioFormat := c.Param("audio_format")

	userID := 0
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		log.Printf("[ERROR] convert user_id failed, err: %+v", err)
		return
	}

	phraseID := 0
	phraseID, err = strconv.Atoi(c.Param("phrase_id"))
	if err != nil {
		log.Printf("[ERROR] convert phrase_id failed, err: %+v", err)
		return
	}

	if !isValidUser(userID) || !isValidPhrase(phraseID) {
		c.JSON(http.StatusOK, gin.H{
			"request_id": requestid.Get(c),
			"error":      "invalid user_id or phrase_id",
		})
		return
	}

	audioFileParam := &dto.AudioFile{
		UserId:   userID,
		PhraseId: phraseID,
	}
	audiFileData, err := dao.GetAudioFile(audioFileParam)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"request_id": requestid.Get(c),
			"error":      "audio file not found",
		})
		return
	}
	storedFilePath := audiFileData.FilePath

	requestedExtension := "." + audioFormat
	if requestedExtension != ".mp3" {
		c.JSON(http.StatusOK, gin.H{
			"request_id": requestid.Get(c),
			"error":      "only allow mp3 format",
		})
		return
	}

	tempFilePath := fmt.Sprintf("./original_file/user_%d_phrase_%d.%s", userID, phraseID, audioFormat)
	if err := utils.ConvertWavToMp3(storedFilePath, tempFilePath); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"request_id": requestid.Get(c),
			"error":      err.Error(),
		})
		return
	}

	c.File(storedFilePath)
}
