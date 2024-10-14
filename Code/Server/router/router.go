package router

import (
	"net/http"
	"os"
	"time"

	"immotep/backend/controllers"
	_ "immotep/backend/docs"

	"github.com/gin-gonic/gin"
	"github.com/maxzerbini/oauth"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "Welcome to Immotep API") })
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", controllers.CreateUser)
			auth.POST("/token", controllers.TokenAuth(bServer))
		}

		root := v1.Group("/")
		{
			root.Use(oauth.Authorize(secretKey, nil))
			root.GET("/users", controllers.GetAllUsers)
			root.GET("/user/:id", controllers.GetUserByID)
			root.GET("/profile", controllers.GetProfile)
		}
	}

	return r
}
