package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"MydroX/shadow-technical-test/pkg/log"
	teamsRepo "MydroX/shadow-technical-test/src/repository/teams"
)

type createTeamRequest struct {
	Name    string `json:"name" binding:"required"`
	TagName string `json:"tag_name" binding:"required"`
}

func CreateTeam(c *gin.Context) {
	req := createTeamRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "There is an invalid or missing input"})
		return
	}

	uuid := fmt.Sprintf("team-%s", uuid.New())

	err = teamsRepo.CreateTeam(uuid, req.Name, req.TagName)
	if err != nil {
		log.Logger.Error("Error creating team", zap.Error(err))
		c.JSON(500, gin.H{"error": "There was an error creating the team"})
		return
	}

	c.JSON(200, gin.H{"message": "Team created"})
}

func GetTeam(c *gin.Context) {
	uuid := c.Param("id")

	team, err := teamsRepo.GetTeamFromUUID(uuid)
	if err != nil {
		log.Logger.Error("Error getting team", zap.Error(err))
		c.JSON(500, gin.H{"error": "There was an error getting the team"})
		return
	}

	c.JSON(200, team)
}

func GetTeams(c *gin.Context) {
	teams, err := teamsRepo.GetTeams()
	if err != nil {
		log.Logger.Error("Error getting teams", zap.Error(err))
		c.JSON(500, gin.H{"error": "There was an error getting the teams"})
		return
	}

	c.JSON(200, teams)
}

func UpdateTeam(c *gin.Context) {
	uuid := c.Param("id")
	req := createTeamRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "There is an invalid or missing input"})
		return
	}

	err = teamsRepo.UpdateTeam(uuid, req.Name, req.TagName)
	if err != nil {
		log.Logger.Error("Error updating team", zap.Error(err))
		c.JSON(500, gin.H{"error": "There was an error updating the team"})
		return
	}

	c.JSON(200, gin.H{"message": "Team updated"})
}

func DeleteTeam(c *gin.Context) {
	uuid := c.Param("id")

	err := teamsRepo.DeleteTeam(uuid)
	if err != nil {
		log.Logger.Error("Error deleting team", zap.Error(err))
		c.JSON(500, gin.H{"error": "There was an error deleting the team"})
		return
	}

	c.JSON(200, gin.H{"message": "Team deleted"})
}
