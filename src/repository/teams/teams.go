package teams

import (
	"MydroX/shadow-technical-test/pkg/db/postgres"
	"MydroX/shadow-technical-test/src/models"
)

func CreateTeam(uuid string, name string, tag string) error {
	res := postgres.Conn.Create(&models.Team{
		UUID:    uuid,
		Name:    name,
		TagName: tag,
	})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func GetTeamFromUUID(uuid string) (*models.Team, error) {
	team := &models.Team{}
	res := postgres.Conn.Where("uuid = ?", uuid).First(team)

	if res.Error != nil {
		return nil, res.Error
	}

	return team, nil
}

func GetTeams() ([]models.Team, error) {
	var teams []models.Team
	res := postgres.Conn.Find(&teams)

	if res.Error != nil {
		return nil, res.Error
	}

	return teams, nil
}

func UpdateTeam(uuid string, name string, tag string) error {
	res := postgres.Conn.Model(&models.Team{}).Where("uuid = ?", uuid).Updates(models.Team{
		Name:    name,
		TagName: tag,
	})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func DeleteTeam(uuid string) error {
	res := postgres.Conn.Where("uuid = ?", uuid).Delete(&models.Team{})

	if res.Error != nil {
		return res.Error
	}

	return nil
}
