package router

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/maxzerbini/oauth"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"immotep/backend/controllers"
	_ "immotep/backend/docs" // mandatory import for swagger doc
)

func Routes() *gin.Engine {
	secretKey := os.Getenv("SECRET_KEY")
	bServer := oauth.NewOAuthBearerServer(
		secretKey,
		time.Hour*24,
		&TestUserVerifier{},
		nil)

	rate := limiter.Rate{
		Period: 1 * time.Hour,
		Limit:  3000,
	}

	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://*", "http://*", "http://localhost:4242"},
		// AllowOriginFunc:  func(origin string) bool { return origin == "https://github.com" },
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(mgin.NewMiddleware(limiter.New(memory.NewStore(), rate)))

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
