package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"speakbuddy-be/pkg/dao"
	"speakbuddy-be/pkg/dto"
	"speakbuddy-be/pkg/utils"
	"strconv"
	"strings"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// UploadAudio handles audio upload, conversion, and database storage
func UploadAudio(c *gin.Context) {
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

	// Check user and phrase validity
	if !isValidUser(userID) || !isValidPhrase(phraseID) {
		c.JSON(http.StatusOK, gin.H{
			"request_id": requestid.Get(c),
			"error":      "invalid user_id or phrase_id",
		})
		return
	}

	// Retrieve the audio file from the request
	file, err := c.FormFile("audio_file")
	if err != nil {
		log.Printf("[ERROR] retrieve the audio file from the request failed, err: %+v", err)
		c.JSON(http.StatusOK, gin.H{
			"request_id": requestid.Get(c),
			"error": "audio file is required",
		})
		return
	}

	tempDir := "./temp"
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
			log.Fatalf("[ERROR] failed to create temp directory: %+v", err)
		}
	}

	// Save the uploaded file temporarily
	tempFilePath := tempDir + "/" + file.Filename
	if err := c.SaveUploadedFile(file, tempFilePath); err != nil {
		log.Printf("[ERROR] save uploaded audio file failed, err: %+v", err)
		c.JSON(http.StatusOK, gin.H{
			"request_id": requestid.Get(c),
			"error": err.Error(),
		})
		return
	}

	storageDir := "./storage"
	if _, err := os.Stat(storageDir); os.IsNotExist(err) {
		if err := os.MkdirAll(storageDir, os.ModePerm); err != nil {
			log.Fatalf("[ERROR] failed to create storage directory: %+v", err)
		}
	}

	// Convert the file to WAV
	fileNameWithoutExt := strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename))
	storedFilePath := fmt.Sprintf("%s/%s_%d_%d.wav", storageDir, fileNameWithoutExt, userID, phraseID)
	if err := utils.ConvertMp3ToWav(tempFilePath, storedFilePath); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"request_id": requestid.Get(c),
			"error": err.Error(),
		})
		return
	}

	newAudiFile := &dto.AudioFile{
		UserId:   userID,
		PhraseId: phraseID,
		FilePath: storedFilePath,
	}
	err = dao.PostAudioFile(newAudiFile)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"request_id": requestid.Get(c),
			"error": "failed to save file details to database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"request_id": requestid.Get(c),
		"message": "audio file uploaded successfully",
	})
}

func isValidUser(userID int) bool {
	userData, err := dao.GetUserById(userID)
	if err != nil {
		log.Printf("[ERROR] get user by id failed, err: %+v", err)
		return false
	}
	return userData != nil
}

func isValidPhrase(phraseID int) bool {
	phraseData, err := dao.GetPhraseById(phraseID)
	if err != nil {
		log.Printf("[ERROR] get phrase by id failed, err: %+v", err)
		return false
	}
	return phraseData != nil
}
