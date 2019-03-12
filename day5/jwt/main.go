package main

import (
	"net/http"
	"strconv"
	"time"
	"log"
	"os"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func main() {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		SigningAlgorithm:	"RS256",
		PrivKeyFile: 		"private-key.pem",
		PubKeyFile: 		"public-key.pem",
		Timeout:				time.Hour,
		Authenticator: func(c *gin.Context)(interface{}, error){
			var login struct{
				Username string `json:"username"`
				Password string `json:"password"`
			}
			if err := c.ShouldBind(&login); err != nil{
				return "", jwt.ErrMissingLoginValues
			}
			if login.Username == "admin" && login.Password == "admin"{
				return "", nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
	})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	r := gin.Default()
	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/add/:a/:b", authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		a, _ := strconv.Atoi(c.Param("a"))
		b, _ := strconv.Atoi(c.Param("b"))
		c.JSON(http.StatusOK, gin.H{
			"result": a + b,
		})
	})

	r.Run(":8000")
}
