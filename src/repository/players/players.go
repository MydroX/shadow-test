package players

import (
	"MydroX/shadow-technical-test/pkg/db/postgres"
	"MydroX/shadow-technical-test/src/models"
	"time"
)

func CreatePlayer(uuid string, firstname string, surename string, dob time.Time, number uint8, position string, nationality string, teamID int) error {
	res := postgres.Conn.Create(&models.Player{
		UUID:        uuid,
		Firstname:   firstname,
		Surename:    surename,
		DateOfBirth: dob,
		Number:      number,
		Position:    position,
		Nationality: nationality,
		TeamID:      teamID,
	})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func GetPlayerFromUUID(uuid string) (*models.Player, error) {
	player := &models.Player{}
	res := postgres.Conn.Where("uuid = ?", uuid).First(player)

	if res.Error != nil {
		return nil, res.Error
	}

	return player, nil
}

func GetPlayers() ([]models.Player, error) {
	var players []models.Player
	res := postgres.Conn.Find(&players)

	if res.Error != nil {
		return nil, res.Error
	}

	return players, nil
}

func UpdatePlayer(uuid string, firstname string, surename string, dob time.Time, number uint8, position string, nationality string) error {
	res := postgres.Conn.Model(&models.Player{}).Where("uuid = ?", uuid).Updates(models.Player{
		Firstname:   firstname,
		Surename:    surename,
		DateOfBirth: dob,
		Number:      number,
		Position:    position,
	})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func DeletePlayer(uuid string) error {
	res := postgres.Conn.Where("uuid = ?", uuid).Delete(&models.Player{})

	if res.Error != nil {
		return res.Error
	}

	return nil
}
