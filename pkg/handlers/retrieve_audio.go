package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"speakbuddy-be/pkg/dao"
	"speakbuddy-be/pkg/dto"
	"speakbuddy-be/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RetrieveAudio retrieves and converts the stored audio file to the requested format
func RetrieveAudio(c *gin.Context) {
	audioFormat := c.Param("audio_format")

	userID := 0
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		log.Fatalf("convert user_id failed, err: %+v", err)
		return
	}

	phraseID := 0
	phraseID, err = strconv.Atoi(c.Param("phrase_id"))
	if err != nil {
		log.Fatalf("convert phrase_id failed, err: %+v", err)
		return
	}

	// Validate user and phrase
	if !isValidUser(userID) || !isValidPhrase(phraseID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id or phrase_id"})
		return
	}

	audioFileParam := &dto.AudioFile{
		UserId:   userID,
		PhraseId: phraseID,
	}
	audiFileData, err := dao.GetAudioFile(audioFileParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "audio file not found"})
		return
	}
	storedFilePath := audiFileData.FilePath

	// If the requested format is different from the stored format, convert the file
	requestedExtension := "." + audioFormat
	if requestedExtension != ".wav" { // Assuming stored format is WAV
		tempFilePath := fmt.Sprintf("./temp/user_%d_phrase_%d.%s", userID, phraseID, audioFormat)

		// Perform the conversion based on the requested format
		switch audioFormat {
		case "mp3":
			err = utils.ConvertMp3ToWav(storedFilePath, tempFilePath)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported audio format"})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "audio conversion failed"})
			return
		}

		defer os.Remove(tempFilePath) // Clean up temporary file after response
		storedFilePath = tempFilePath
	}

	// Serve the audio file
	c.File(storedFilePath)
}
