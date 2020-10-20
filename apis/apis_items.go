package apis

import (
	"net/http"
	"vgo2/models"
	"github.com/gin-gonic/gin"
)

func ItemsIndex(c *gin.Context) {
	s := models.Item{Title: "sea",Notes: "nnn"}
	c.JSON(http.StatusOK, gin.H{
		"message": "Hola"+s.Title,
	})
}