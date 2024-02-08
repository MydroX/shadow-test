package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"MydroX/shadow-technical-test/pkg/log"
	playersRepo "MydroX/shadow-technical-test/src/repository/players"
)

type createPlayerRequest struct {
	Firstname   string `json:"firstname" binding:"required"`
	Surename    string `json:"surename" binding:"required"`
	DateOfBirth string `json:"dob" binding:"required"`
	Number      uint8  `json:"number" binding:"required"`
	Position    string `json:"position" binding:"required"`
	Nationality string `json:"nationality" binding:"required"`
	TeamID      int    `json:"team" binding:"required"`
}

func CreatePlayer(c *gin.Context) {
	req := createPlayerRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "There is an invalid or missing input"})
		return
	}

	uuid := fmt.Sprintf("player-%s", uuid.New())

	dob, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		c.JSON(400, gin.H{"error": "There is an invalid date of birth"})
		return
	}

	err = playersRepo.CreatePlayer(uuid, req.Firstname, req.Surename, dob, req.Number, req.Position, req.Nationality, req.TeamID)
	if err != nil {
		log.Logger.Error("Error creating player", zap.Error(err))
		c.JSON(500, gin.H{"error": "There was an error creating the player"})
		return
	}

	c.JSON(200, gin.H{"message": "Player created"})
}

func GetPlayer(c *gin.Context) {
	uuid := c.Param("id")

	player, err := playersRepo.GetPlayerFromUUID(uuid)
	if err != nil {
		log.Logger.Error("Error getting player", zap.Error(err))
		c.JSON(500, gin.H{"error": "There was an error getting the player"})
		return
	}

	c.JSON(200, player)
}

func GetPlayers(c *gin.Context) {
	players, err := playersRepo.GetPlayers()
	if err != nil {
		log.Logger.Error("Error getting players", zap.Error(err))
		c.JSON(500, gin.H{"error": "There was an error getting the players"})
		return
	}

	c.JSON(200, players)
}

func UpdatePlayer(c *gin.Context) {
	uuid := c.Param("id")
	req := createPlayerRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "There is an invalid or missing input"})
		return
	}

	dob, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		c.JSON(400, gin.H{"error": "There is an invalid date of birth"})
		return
	}

	err = playersRepo.UpdatePlayer(uuid, req.Firstname, req.Surename, dob, req.Number, req.Position, req.Nationality)
	if err != nil {
		log.Logger.Error("Error updating player", zap.Error(err))
		c.JSON(500, gin.H{"error": "There was an error updating the player"})
		return
	}

	c.JSON(200, gin.H{"message": "Player updated"})
}

func DeletePlayer(c *gin.Context) {
	uuid := c.Param("id")

	err := playersRepo.DeletePlayer(uuid)
	if err != nil {
		log.Logger.Error("Error deleting player", zap.Error(err))
		c.JSON(500, gin.H{"error": "There was an error deleting the player"})
		return
	}

	c.JSON(200, gin.H{"message": "Player deleted"})
}
