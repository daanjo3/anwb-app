package handler

import (
	"fmt"

	"github.com/daanjo3/anweb-app/api/internal/anwb"
	"github.com/gin-gonic/gin"
)

func ListJams(c *gin.Context) {
	// TODO make shared constant
	allJams := make([]anwb.RoadInfo, 0)
	document, exists := c.Get(KEY_DOCUMENT)
	if !exists {
		c.Status(500)
		fmt.Fprint(c.Writer, "AWNB document handler failed")
		return
	}
	for _, road := range document.(anwb.Document).Roads {
		for _, segment := range road.Segments {
			if len(segment.Jams) > 0 {
				allJams = append(allJams, segment.Jams...)
			}
		}
	}
	c.JSON(200, allJams)
}
