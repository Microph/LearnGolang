package gin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetGinRouter(r *gin.Engine) {
	r.GET("/add/:a/:b", AddGinHandler)
}

func AddGinHandler(c *gin.Context) {
	a, _ := strconv.Atoi(c.Param("a"))
	b, _ := strconv.Atoi(c.Param("b"))

	c.JSON(http.StatusOK, gin.H{
		"result": a + b,
	})
}
