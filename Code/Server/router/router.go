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
	"immotep/backend/router/middlewares"
)

func registerAPIRoutes(r *gin.Engine) {
	secretKey := os.Getenv("SECRET_KEY")
	bServer := oauth.NewOAuthBearerServer(
		secretKey,
		time.Hour*24,
		&TestUserVerifier{},
		nil)

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", controllers.RegisterOwner)
			auth.POST("/invite/:id", controllers.RegisterTenant)
			auth.POST("/token", controllers.TokenAuth(bServer))
		}

		root := v1.Group("/")
		{
			root.Use(oauth.Authorize(secretKey, nil))
			root.Use(middlewares.CheckClaims())
			root.GET("/users", controllers.GetAllUsers)
			root.GET("/user/:id", controllers.GetUserByID)
			root.GET("/user/:id/picture", controllers.GetUserProfilePicture)
			root.GET("/profile", controllers.GetCurrentUserProfile)
			root.PUT("/profile", controllers.UpdateCurrentUserProfile)
			root.GET("/profile/picture", controllers.GetCurrentUserProfilePicture)
			root.PUT("/profile/picture", controllers.UpdateCurrentUserProfilePicture)

			owner := root.Group("/owner")
			registerOwnerRoutes(owner)
		}
	}
}

func registerOwnerRoutes(owner *gin.RouterGroup) {
	owner.Use(middlewares.AuthorizeOwner())
	owner.POST("/properties", controllers.CreateProperty)
	owner.GET("/properties", controllers.GetAllProperties)

	property := owner.Group("/properties/:property_id")
	{
		property.Use(middlewares.CheckPropertyOwnership("property_id"))
		property.GET("/", controllers.GetPropertyById)
		property.GET("/picture", controllers.GetPropertyPicture)
		property.PUT("/picture", controllers.UpdatePropertyPicture)

		rooms := property.Group("/rooms")
		{
			rooms.POST("/", controllers.CreateRoom)
			rooms.GET("/", controllers.GetRoomsByProperty)
			rooms.GET("/:room_id", controllers.GetRoomByID)
			rooms.DELETE("/:room_id", controllers.DeleteRoom)
		}

		furnitures := property.Group("/rooms/:room_id/furnitures")
		{
			furnitures.Use(middlewares.CheckRoomExists("property_id", "room_id"))
			furnitures.POST("/", controllers.CreateFurniture)
			furnitures.GET("/", controllers.GetFurnituresByRoom)
			furnitures.GET("/:furniture_id", controllers.GetFurnitureByID)
			furnitures.DELETE("/:furniture_id", controllers.DeleteFurniture)
		}
	}

	owner.POST("/send-invite/:propertyId", controllers.InviteTenant)
}

func Routes() *gin.Engine {
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

	registerAPIRoutes(r)
	return r
}
