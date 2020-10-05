package router

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"[[.ModName]]/di"
	_ "[[.ModName]]/docs"
	"[[.ModName]]/infrastructure/db"
	"[[.ModName]]/presentation/middleware"
	defaultMiddleware "github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	chi.Router
}

func Init() *Router {
	r := chi.NewRouter()
	return &Router{r}
}

func (r *Router) Routes() {
	conn := db.New()
	meHandler, err := di.MeHandler(conn)
	if err != nil {
		log.Fatalln(err)
	}
	authMiddleware, err := di.AuthMiddleware(conn)
	if err != nil {
		log.Fatalln(err)
	}
	maintenanceMiddleware, err := di.MaintenanceMiddleware(conn)
	if err != nil {
		log.Fatalln(err)
	}
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Use(maintenanceMiddleware.All)
				r.Use(authMiddleware.Auth)
				r.Route("/me", func(r chi.Router) {
					r.Get("/", meHandler.Get)
				})
			})
		})
	})
}

func (r *Router) NotProd() {
	r.Use(middleware.CORS)
	r.Use(defaultMiddleware.Logger)
	r.Get("/api/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/api/swagger/doc.json"),
	))
}

func (r *Router) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
