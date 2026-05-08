package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/jackc/pgx/v5/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "github.com/thisgleammm/mantis-backend/cmd/docs"
	repo "github.com/thisgleammm/mantis-backend/internal/adapters/postgresql/sqlc"
	"github.com/thisgleammm/mantis-backend/internal/env"
	"github.com/thisgleammm/mantis-backend/internal/handler"
	"github.com/thisgleammm/mantis-backend/internal/middleware"
	"github.com/thisgleammm/mantis-backend/internal/repository/postgresql"
	"github.com/thisgleammm/mantis-backend/internal/service"
)

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://mantis-marketplace.vercel.app", "http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(chiMiddleware.RequestID)    // important for rate limiting
	r.Use(chiMiddleware.RealIP)       // impor for rate limiting and analytics and tracing
	r.Use(chiMiddleware.Logger)       // logging requests
	r.Use(chiMiddleware.Recoverer)    // recover from panics
	r.Use(middleware.SecurityHeaders) // security headers

	r.Use(chiMiddleware.Timeout(60 * time.Second)) // timeout for requests

	// Routes
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})

	r.With(middleware.RelaxedSecurityHeaders).Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition
	))

	queries := repo.New(app.db)
	jwtSecret := env.RequiredString("JWT_SECRET")

	productRepo := postgresql.NewProductRepository(queries)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	categoryRepo := postgresql.NewCategoryRepository(queries)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	userRepo := postgresql.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	authService := service.NewAuthService(userRepo, jwtSecret)
	authHandler := handler.NewAuthHandler(authService)

	cartRepo := postgresql.NewCartRepository(queries)
	cartService := service.NewCartService(cartRepo)
	cartHandler := handler.NewCartHandler(cartService)

	orderRepo := postgresql.NewOrderRepository(queries)
	orderService := service.NewOrderService(orderRepo, cartRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	// API v1 Routing
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/products", func(r chi.Router) {
			r.Get("/", productHandler.ListProducts)
			r.Post("/", productHandler.CreateProduct)
			r.Get("/{slug}", productHandler.FindProductBySlug)
		})

		r.Route("/categories", func(r chi.Router) {
			r.Get("/", categoryHandler.ListCategories)
			r.Get("/{id}", categoryHandler.FindCategoryByID)
		})

		r.Route("/users", func(r chi.Router) {
			r.Get("/", userHandler.ListUsers)
			r.Get("/{id}", userHandler.FindUserByID)
			r.With(middleware.Middleware).Get("/me", userHandler.GetMe)
		})

		r.Route("/carts", func(r chi.Router) {
			r.Use(middleware.Middleware)
			r.Get("/", cartHandler.ListCarts)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/items", cartHandler.ListCartItems)
				r.Post("/items", cartHandler.AddItemToCart)
			})
			r.Route("/items", func(r chi.Router) {
				r.Patch("/{id}", cartHandler.UpdateItemQuantity)
				r.Delete("/{id}", cartHandler.RemoveItemFromCart)
			})
		})

		r.Route("/orders", func(r chi.Router) {
			r.Use(middleware.Middleware)
			r.Get("/", orderHandler.ListOrders)
			r.Post("/checkout", orderHandler.Checkout)
		})

		r.Route("/auth", func(r chi.Router) {
			r.Use(httprate.LimitByIP(5, 1*time.Minute)) // rate limit auth attempts
			r.Post("/register", authHandler.Register)
			r.Post("/login", authHandler.Login)
			r.Post("/refresh", authHandler.Refresh)
			r.Post("/logout", authHandler.Logout)
		})

	})

	return r

}

// run
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.Config.addr,
		Handler:      h,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	slog.Info("server has started", "addr", app.Config.addr)

	return srv.ListenAndServe()
}

type application struct {
	Config config
	// logger
	db *pgxpool.Pool
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
