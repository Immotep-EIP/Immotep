package router

import (
	"log"
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

func registerAPIRoutes(r *gin.Engine, test bool) {
	secretKey := os.Getenv("SECRET_KEY")
	bServer := oauth.NewOAuthBearerServer(
		secretKey,
		time.Hour*24,
		&TestUserVerifier{},
		nil)

	v1 := r.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register/", controllers.RegisterOwner)
			auth.POST("/invite/:id/", controllers.RegisterTenant)
			if !test {
				auth.POST("/token/", controllers.TokenAuth(bServer))
			}
		}

		root := v1.Group("/")
		{
			if !test {
				root.Use(oauth.Authorize(secretKey, nil))
			} else {
				root.Use(middlewares.MockClaims())
			}
			root.Use(middlewares.CheckClaims())
			root.GET("/users/", controllers.GetAllUsers)
			root.GET("/user/:id/", controllers.GetUserByID)
			root.GET("/user/:id/picture/", controllers.GetUserProfilePicture)
			root.GET("/profile/", controllers.GetCurrentUserProfile)
			root.PUT("/profile/", controllers.UpdateCurrentUserProfile)
			root.GET("/profile/picture/", controllers.GetCurrentUserProfilePicture)
			root.PUT("/profile/picture/", controllers.UpdateCurrentUserProfilePicture)

			owner := root.Group("/owner")
			registerOwnerRoutes(owner)

			tenant := root.Group("/tenant")
			registerTenantRoutes(tenant)
		}
	}
}

func registerOwnerRoutes(owner *gin.RouterGroup) {
	owner.Use(middlewares.AuthorizeOwner())

	properties := owner.Group("/properties")
	{
		properties.POST("/", controllers.CreateProperty)
		properties.GET("/", controllers.GetAllProperties)
		properties.GET("/archived/", controllers.GetAllArchivedProperties)

		propertyId := properties.Group("/:property_id/")
		{
			propertyId.Use(middlewares.CheckPropertyOwnership("property_id"))
			propertyId.GET("/", controllers.GetPropertyById)
			propertyId.PUT("/", controllers.UpdateProperty)
			propertyId.PUT("/archive/", controllers.ArchiveProperty)
			propertyId.GET("/inventory/", controllers.GetPropertyInventory)
			propertyId.GET("/picture/", controllers.GetPropertyPicture)
			propertyId.PUT("/picture/", controllers.UpdatePropertyPicture)

			propertyId.POST("/send-invite/", controllers.InviteTenant)
			propertyId.DELETE("/cancel-invite/", middlewares.CheckPendingContract("property_id"), controllers.CancelInvite)

			contract := propertyId.Group("")
			{
				contract.Use(middlewares.CheckActiveContract("property_id"))
				contract.PUT("/end-contract/", controllers.EndContract)
				contract.GET("/documents/", controllers.GetPropertyDocuments)
			}

			rooms := propertyId.Group("/rooms")
			{
				rooms.POST("/", controllers.CreateRoom)
				rooms.GET("/", controllers.GetRoomsByProperty)
				rooms.GET("/archived/", controllers.GetArchivedRoomsByProperty)

				roomId := rooms.Group("/:room_id")
				{
					roomId.Use(middlewares.CheckRoomOwnership("property_id", "room_id"))
					roomId.GET("/", controllers.GetRoomByID)
					roomId.PUT("/archive/", controllers.ArchiveRoom)

					furnitures := roomId.Group("/furnitures")
					{
						furnitures.POST("/", controllers.CreateFurniture)
						furnitures.GET("/", controllers.GetFurnituresByRoom)
						furnitures.GET("/archived/", controllers.GetArchivedFurnituresByRoom)

						furnitureId := furnitures.Group("/:furniture_id")
						{
							furnitureId.Use(middlewares.CheckFurnitureOwnership("room_id", "furniture_id"))
							furnitureId.GET("/", controllers.GetFurnitureByID)
							furnitureId.PUT("/archive/", controllers.ArchiveFurniture)
						}
					}
				}
			}

			invReports := propertyId.Group("/inventory-reports")
			registerInvReportRoutes(invReports)
		}
	}
}

func registerTenantRoutes(tenant *gin.RouterGroup) {
	tenant.Use(middlewares.AuthorizeTenant())

	tenant.POST("/invite/:id/", controllers.AcceptInvite)
}

func registerInvReportRoutes(invReports *gin.RouterGroup) {
	invReports.POST("/", controllers.CreateInventoryReport)
	invReports.GET("/", controllers.GetInventoryReportsByProperty)

	invReports.GET("/:report_id/",
		middlewares.CheckInventoryReportOwnership("property_id", "report_id"),
		controllers.GetInventoryReportByID)

	invReports.POST("/summarize/", controllers.GenerateSummary)
	invReports.POST("/compare/:old_report_id/",
		middlewares.CheckInventoryReportOwnership("property_id", "old_report_id"),
		controllers.GenerateComparison)
}

func Routes() *gin.Engine {
	rate := limiter.Rate{
		Period: 1 * time.Hour,
		Limit:  3000,
	}

	var allowOrigins []string
	if gin.Mode() == gin.ReleaseMode {
		allowOrigins = []string{os.Getenv("WEB_PUBLIC_URL")}
		log.Println("Running in release mode")
	} else {
		allowOrigins = []string{"https://*", "http://*", "http://localhost:4242", "http://localhost:3002"}
		log.Println("Running in debug mode")
	}

	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins: allowOrigins,
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
	r.Use(gin.CustomRecovery(middlewares.PanicRecovery))
	r.Use(mgin.NewMiddleware(limiter.New(memory.NewStore(), rate)))

	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "Welcome to Immotep API") })
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	registerAPIRoutes(r, false)
	return r
}

func TestRoutes() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "Welcome to Immotep API") })
	registerAPIRoutes(r, true)
	return r
}
