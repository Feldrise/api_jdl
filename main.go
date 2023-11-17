package main

import (
	"log"
	"net/http"

	"feldrise.com/jdl/config"
	"feldrise.com/jdl/database"
	"feldrise.com/jdl/game"
	"feldrise.com/jdl/gamecard"
	"feldrise.com/jdl/group"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Recoverer,
		group.Middleware(configuration),
		cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedHeaders:   []string{"*"},
			AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "DELETE"},
			AllowCredentials: true,
		}).Handler,
	)

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/games", game.New(configuration).Routes())
		r.Mount("/games/{gameid}/cards", gamecard.New(configuration).Routes())
		r.Mount("/groups", group.New(configuration).Routes())
	})

	fs := http.FileServer(http.Dir(configuration.Constants.DataPath + "/uploads"))
	router.Handle("/uploads/*", http.StripPrefix("/uploads/", fs))

	return router
}

func main() {
	// We initialize the project
	configuration, err := config.New()

	if err != nil {
		log.Panicln("Configuration error", err)
	}
	router := Routes(configuration)

	// Swagger configs
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	// We do the database migration
	database.Migrate(configuration.Database)

	// We show all the routes in the logs
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	// We serve the API
	log.Printf("connect to http://localhost:%s/swagger/index.html for documentation", configuration.Constants.Port)
	log.Fatal(http.ListenAndServe(":"+configuration.Constants.Port, router))
}
