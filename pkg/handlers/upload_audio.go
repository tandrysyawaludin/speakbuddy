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

// UploadAudio handles audio upload, conversion, and database storage
func UploadAudio(c *gin.Context) {
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

	// Check user and phrase validity
	if !isValidUser(userID) || !isValidPhrase(phraseID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id or phrase_id"})
		return
	}

	// Retrieve the audio file from the request
	file, err := c.FormFile("audio_file")
	if err != nil {
		log.Fatalf("retrieve the audio file from the request failed, err: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "audio file is required"})
		return
	}

	// Save the uploaded file temporarily
	tempFilePath := "./temp/" + file.Filename
	if err := c.SaveUploadedFile(file, tempFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer os.Remove(tempFilePath) // Clean up temp file

	// Convert the file to WAV
	storedFilePath := fmt.Sprintf("./storage/user_%d_phrase_%d.wav", userID, phraseID)
	if err := utils.ConvertMp3ToWav(tempFilePath, storedFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	newAudiFile := &dto.AudioFile{
		UserId:   userID,
		PhraseId: phraseID,
		FilePath: storedFilePath,
	}
	err = dao.PostAudioFile(newAudiFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file details to database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "audio file uploaded successfully", "file_path": storedFilePath})
}

func isValidUser(userID int) bool {
	userData, err := dao.GetUserById(userID)
	if err != nil {
		log.Fatalf("get user by id failed, err: %+v", err)
		return false
	}
	return userData != nil
}

func isValidPhrase(phraseID int) bool {
	phraseData, err := dao.GetPhraseById(phraseID)
	if err != nil {
		log.Fatalf("get phrase by id failed, err: %+v", err)
		return false
	}
	return phraseData != nil
}
