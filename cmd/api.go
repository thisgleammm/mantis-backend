package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/thisgleammm/mantis-backend/internal/auth"
	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
	"github.com/thisgleammm/mantis-backend/internal/categories"
	"github.com/thisgleammm/mantis-backend/internal/products"
	"github.com/thisgleammm/mantis-backend/internal/users"
	_ "github.com/thisgleammm/mantis-backend/cmd/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID) // important for rate limiting
	r.Use(middleware.RealIP)    // impor for rate limiting and analytics and tracing
	r.Use(middleware.Logger)    // logging requests
	r.Use(middleware.Recoverer) // recover from panics

	r.Use(middleware.Timeout(60 * time.Second)) // timeout for requests

	// Routes
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})
	
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))

	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)

	categoriesService := categories.NewService(repo.New(app.db))
	categoriesHandler := categories.NewHandler(categoriesService)

	userService := users.NewService(repo.New(app.db))
	userHandler := users.NewHandler(userService)

	authService := auth.NewService(repo.New(app.db))
	authHandler := auth.NewHandler(authService)

	// API v1 Routing
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/products", func(r chi.Router) {
			r.Get("/", productHandler.ListProducts)
			r.Get("/{id:[0-9]+}", productHandler.FindProductByID)
			r.Get("/{slug}", productHandler.FindProductBySlug)
		})

		r.Route("/categories", func(r chi.Router) {
			r.Get("/", categoriesHandler.ListCategories)
			r.Get("/{id}", categoriesHandler.FindCategoryByID)
		})

		r.Route("/users", func(r chi.Router) {
			r.Get("/", userHandler.ListUsers)
			r.Get("/{id:[0-9]+}", userHandler.FindUserByID)
			r.With(auth.Middleware).Get("/me", userHandler.GetMe)
		})

		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", authHandler.Register)
			r.Post("/login", authHandler.Login)
		})


	})

	return r

}

// run
func (app *application) run(h http.Handler) error{
	srv := &http.Server{
		Addr : app.Config.addr,
		Handler : h,
		WriteTimeout : 30 * time.Second,
		ReadTimeout : 10 * time.Second,
		IdleTimeout : time.Minute,
	}

	slog.Info("server has started", "addr", app.Config.addr)

	return srv.ListenAndServe()
}

type application struct {
	Config config
	// logger
	db *pgx.Conn
}

type config struct {
	addr string
	db dbConfig
}

type dbConfig struct {
	dsn string
}