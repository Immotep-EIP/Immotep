package router

import (
	"net/http"
	"os"
	"time"

	"immotep/backend/controllers"

	"github.com/gin-gonic/gin"
	"github.com/maxzerbini/oauth"
)

func Routes() *gin.Engine {
	secretKey := os.Getenv("SECRET_KEY")

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	bServer := oauth.NewOAuthBearerServer(
		secretKey,
		time.Hour*24,
		&TestUserVerifier{},
		nil)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Immotep API")
	})

	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.CreateUser)
		auth.POST("/token", bServer.UserCredentials)
	}

	api := r.Group("/")
	{
		api.Use(oauth.Authorize(secretKey, nil))
		api.GET("/users", controllers.GetAllUsers)
		api.GET("/user/:id", controllers.GetUserByID)
		api.GET("/profile", controllers.GetProfile)
	}

	return r
}
