package router

import (
	"net/http"
	"os"
	"time"

	"immotep/backend/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/oauth"
)

func Routes() http.Handler {
	secretKey := os.Getenv("SECRET_KEY")

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	bServer := oauth.NewBearerServer(
		secretKey,
		time.Hour*24,
		&TestUserVerifier{},
		nil)

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Immotep API"))
	})

	router.Route("/auth", func(r chi.Router) {
		r.Post("/register", controllers.CreateUser)
		r.Post("/token", bServer.UserCredentials)
	})

	router.Route("/api", func(r chi.Router) {
		r.Use(oauth.Authorize(secretKey, nil))
		r.Get("/users", controllers.GetAllUsers)
		r.Get("/user/{id}", controllers.GetUserByID)
		r.Get("/profile", controllers.GetProfile)
	})

	return router
}
