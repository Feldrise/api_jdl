package game

import (
	"net/http"

	"feldrise.com/jdl/errors"
	"feldrise.com/jdl/group"
	"feldrise.com/jdl/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// CRUD

// GetAGame godoc
//
// @Summary Get one game
// @Descripton Get one game by its ID
// @ID get-a-game
// @Tags Game
// @Success 200 {object} Game
// @Failure 404 {object} ErrResponse
// @Router /games/{id} [get]
func (config *Config) Get(w http.ResponseWriter, r *http.Request) {
	group := group.ForContext(r.Context())
	if group == nil {
		render.Render(w, r, errors.ErrUnauthorized("you must specify the group"))
		return
	}

	id := chi.URLParam(r, "id")
	var game models.Game
	config.Database.Find(&game, id)

	if game.ID == 0 {
		render.Render(w, r, errors.ErrNotFound())
		return
	}

	if game.GroupID != *group {
		render.Render(w, r, errors.ErrUnauthorized("you are not in the right group"))
		return
	}

	render.JSON(w, r, game)
}

// GetGames godoc
//
// @Summary Get all games
// @Descripton Get all games
// @ID get-games
// @Tags Game
// @Success 200 {object} []Game
// @Failure 404 {object} ErrResponse
// @Router /games/ [get]
func (config *Config) GetAll(w http.ResponseWriter, r *http.Request) {
	group := group.ForContext(r.Context())
	if group == nil {
		render.Render(w, r, errors.ErrUnauthorized("you must specify the group"))
		return
	}

	var games []models.Game
	config.Database.Model(&models.Game{}).Where("group_id=?", group).Find(&games)

	render.JSON(w, r, games)
}

// CreateGame godoc
// @Summary Create a game
// @Description Create a game
// @ID create-game
// @Tags Game
// @Param request body GamePostPayload true "game's info"
// @Success 201 {object} Game
// @Failure 400 {object} ErrResponse
// @Router /games/ [post]
func (config *Config) Create(w http.ResponseWriter, r *http.Request) {
	// We bind the data
	data := &models.GamePostPayload{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	// We check authorization
	group := group.ForContext(r.Context())
	if group == nil {
		render.Render(w, r, errors.ErrUnauthorized("you must specify the group"))
		return
	}

	// We create the game
	game := models.Game{
		GroupID: *group,
		Name:    *data.Name,
		Type:    *data.Type,
	}

	config.Database.Create(&game)

	render.Status(r, 201)
	render.JSON(w, r, game)
}
