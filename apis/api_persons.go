package apis

import (
	"net/http"
	"vgo2/models"
	"github.com/gin-gonic/gin"
)

func PersonsIndex(c *gin.Context) {
	s := models.Person{Name: "sea",Age: "nnn"}
	c.JSON(http.StatusOK, gin.H{
		"message": "Hola"+s.Name,
	})
}