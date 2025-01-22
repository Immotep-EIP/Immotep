package router

import (
	"net/http"
	"os"
	"time"

	"immotep/backend/controllers"
	_ "immotep/backend/docs" // mandatory import for swagger doc
	"immotep/backend/router/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/maxzerbini/oauth"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
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

	properties := owner.Group("/properties")
	{
		properties.POST("/", controllers.CreateProperty)
		properties.GET("/", controllers.GetAllProperties)

		propertyId := properties.Group("/:property_id")
		{
			propertyId.Use(middlewares.CheckPropertyOwnership("property_id"))
			propertyId.GET("/", controllers.GetPropertyById)
			propertyId.GET("/picture", controllers.GetPropertyPicture)
			propertyId.PUT("/picture", controllers.UpdatePropertyPicture)

			propertyId.POST("/send-invite", controllers.InviteTenant)
			propertyId.PUT("/end-contract", controllers.EndContract)

			rooms := propertyId.Group("/rooms")
			{
				rooms.POST("/", controllers.CreateRoom)
				rooms.GET("/", controllers.GetRoomsByProperty)

				roomId := rooms.Group("/:room_id")
				{
					roomId.Use(middlewares.CheckRoomOwnership("property_id", "room_id"))
					roomId.GET("/", controllers.GetRoomByID)
					roomId.DELETE("/", controllers.DeleteRoom)

					furnitures := roomId.Group("/furnitures")
					{
						furnitures.POST("/", controllers.CreateFurniture)
						furnitures.GET("/", controllers.GetFurnituresByRoom)

						furnitureId := furnitures.Group("/:furniture_id")
						{
							furnitureId.Use(middlewares.CheckFurnitureOwnership("room_id", "furniture_id"))
							furnitureId.GET("/", controllers.GetFurnitureByID)
							furnitureId.DELETE("/", controllers.DeleteFurniture)
						}
					}
				}
			}

			invReports := propertyId.Group("/inventory-reports")
			registerInvReportRoutes(invReports)
		}
	}
}

func registerInvReportRoutes(invReports *gin.RouterGroup) {
	invReports.POST("/", controllers.CreateInventoryReport)
	invReports.GET("/", controllers.GetInventoryReportsByProperty)

	invReports.GET("/:report_id",
		middlewares.CheckInventoryReportOwnership("property_id", "report_id"),
		controllers.GetInventoryReportByID)

	invReports.POST("/summarize", controllers.GenerateSummary)
	invReports.POST("/compare/:old_report_id",
		middlewares.CheckInventoryReportOwnership("property_id", "old_report_id"),
		controllers.GenerateComparison)
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
