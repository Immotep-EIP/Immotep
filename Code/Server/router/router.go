package router

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/maxzerbini/oauth"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"keyz/backend/controllers"
	_ "keyz/backend/docs" // mandatory import for swagger doc
	"keyz/backend/router/middlewares"
	"keyz/backend/router/validators"
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
		contact := v1.Group("/contact/")
		{
			contact.Use(mgin.NewMiddleware(limiter.New(memory.NewStore(), limiter.Rate{
				Period: 1 * time.Hour,
				Limit:  1,
			})))
			contact.POST("/", controllers.CreateContactMessage)
		}

		auth := v1.Group("/auth/")
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

			owner := root.Group("/owner/")
			registerOwnerRoutes(owner)

			tenant := root.Group("/tenant/")
			registerTenantRoutes(tenant)
		}
	}
}

func registerValidators() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		panic("Could not register validator")
	}
	_ = v.RegisterValidation("priority", validators.Priority)
	_ = v.RegisterValidation("reportType", validators.ReportType)
	_ = v.RegisterValidation("state", validators.State)
	_ = v.RegisterValidation("cleanliness", validators.Cleanliness)
	_ = v.RegisterValidation("roomType", validators.RoomType)
}

func Routes() *gin.Engine {
	var allowOrigins []string
	if gin.Mode() == gin.ReleaseMode {
		allowOrigins = []string{os.Getenv("WEB_PUBLIC_URL"), os.Getenv("SHOWCASE_PUBLIC_URL")}
		log.Println("Running in release mode")
	} else {
		allowOrigins = []string{"https://*", "http://*", "http://localhost:4242", "http://localhost:3002", "http://localhost:3000"}
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
	r.Use(mgin.NewMiddleware(limiter.New(memory.NewStore(), limiter.Rate{
		Period: 1 * time.Hour,
		Limit:  3000,
	})))

	registerValidators()

	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "Welcome to Keyz API") })
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	registerAPIRoutes(r, false)
	return r
}

func TestRoutes() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	registerValidators()
	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "Welcome to Keyz API") })
	registerAPIRoutes(r, true)
	return r
}
