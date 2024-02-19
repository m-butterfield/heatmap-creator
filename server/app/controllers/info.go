package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/heatmap-creator/server/app/lib"
	"time"
)

type InfoResponse struct {
	MaxQueries int `json:"maxQueries"`
	NumQueries int `json:"numQueries"`
}

func info(c *gin.Context) {
	token := c.Query("t")
	if token == "" {
		c.AbortWithStatus(403)
		return
	}

	accessToken, err := ds.GetAccessTokenByQueryToken(token)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	if accessToken == nil {
		c.AbortWithStatus(403)
		return
	}

	yesterday := time.Now().AddDate(0, 0, -1)
	queryCount, err := ds.GetQueryCountForUser(accessToken.UserID, &yesterday)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	c.JSON(200, InfoResponse{
		MaxQueries: accessToken.User.DailyQueries,
		NumQueries: queryCount,
	})
}
