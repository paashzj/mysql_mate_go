package api

import (
	"github.com/gin-gonic/gin"
	"mysql_mate_go/pkg/db"
	"net/http"
)

func GlobalStatusHandler(c *gin.Context) {
	status := db.GetGlobalStatus()
	c.JSON(http.StatusOK, status)
}

func SlaveStatusHandler(c *gin.Context) {
	status := db.GetSlaveStatus()
	c.JSON(http.StatusOK, status)
}
