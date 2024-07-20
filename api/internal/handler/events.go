package handler

import (
	"log/slog"
	"net/http"

	"github.com/daanjo3/anweb-app/api/internal/anwb"
	"github.com/daanjo3/anweb-app/api/internal/service"
	"github.com/gin-gonic/gin"
)

// TODO bit less code duplication here would be nice
func ListJams(c *gin.Context) {
	document, exists := c.Get(KEY_DOCUMENT)
	if !exists {
		slog.Error("Document did not exist in child handler")
		c.JSON(http.StatusInternalServerError, "AWNB document handler failed")
		return
	}
	jams := service.ListJams(document.(anwb.Document))
	c.JSON(200, jams)
}

func ListRoadWorks(c *gin.Context) {
	document, exists := c.Get(KEY_DOCUMENT)
	if !exists {
		slog.Error("Document did not exist in child handler")
		c.JSON(http.StatusInternalServerError, "AWNB document handler failed")
		return
	}
	roadworks := service.ListRoadWorks(document.(anwb.Document))
	c.JSON(200, roadworks)
}

func ListRadars(c *gin.Context) {
	document, exists := c.Get(KEY_DOCUMENT)
	if !exists {
		slog.Error("Document did not exist in child handler")
		c.JSON(http.StatusInternalServerError, "AWNB document handler failed")
		return
	}
	radars := service.ListRadars(document.(anwb.Document))
	c.JSON(200, radars)
}
