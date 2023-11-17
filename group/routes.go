package group

import (
	"feldrise.com/jdl/config"
	"github.com/go-chi/chi"
)

func New(configuration *config.Config) *Config {
	return &Config{configuration}
}

func (config *Config) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/code/{code}", config.GetFromCode)

	return router
}
