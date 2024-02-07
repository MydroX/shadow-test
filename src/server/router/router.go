package router

import (
	"MydroX/shadow-technical-test/src/controller"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		teams := v1.Group("/teams")
		{
			// teams.GET("/", controller.GetTeams)
			// teams.GET("/:id", controller.GetTeam)
			teams.POST("/create", controller.CreateTeam)
			// teams.PUT("/:id/update", controller.UpdateTeam)
			// teams.DELETE("/:id/delete", controller.DeleteTeam)
		}
		// players := v1.Group("/players")
		// {
		// players.GET("/", controller.GetPlayers)
		// players.GET("/:id", controller.GetPlayer)
		// players.POST("/create", controller.CreatePlayer)
		// players.PUT("/:id/update", controller.UpdatePlayer)
		// players.DELETE("/:id/delete", controller.DeletePlayer)
		// }

		v1.GET("/health", controller.HealthEndpoint)
	}
	return r
}
