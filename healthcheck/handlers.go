package healthcheck

import (
	"github.com/gin-gonic/gin"
	"github.com/remoteview/service-teams/version"
)

// Handler - route
// @Summary Health check
// @Description Health check
// @Accept  json
// @Produce  json
// @Router /_health [get]
func Handler(c *gin.Context) {
	version := version.GetVersion()
	c.JSON(200, HealthCheck{Status: "Ok", Version: version})
}
