package router

import (
	"evelp/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func corporation(c *gin.Context) {
	factionId, err := strconv.Atoi(c.Param("factionId"))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	corporations, err := model.GetCorporationsByFaction(factionId)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, corporations)
}