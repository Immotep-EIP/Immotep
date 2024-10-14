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

			// TokenAuth godoc
			//
			//	@Summary		Authenticate user
			//	@Description	Authenticate user with email and password
			//	@Accept			json
			//	@Produce		json
			//	@Param			grant_type		formData	string			true	"password / refresh_token"
			//	@Param			username		formData	string			false	"User email"
			//	@Param			password		formData	string			false	"User password"
			//	@Param			refresh_token	formData	string			false	"Refresh token"
			//	@Success		200				{object}	TokenResponse	"Token data"
			//	@Failure		400				{object}	Error			"Invalid grant_type"
			//	@Failure		401				{object}	Error			"Unauthorized"
			//	@Failure		500
			//	@Router			/auth/token [post]
			auth.POST("/token", bServer.UserCredentials)
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
