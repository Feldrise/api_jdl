package gamemode

import (
	"net/http"

	"feldrise.com/jdl/errors"
	"feldrise.com/jdl/group"
	"feldrise.com/jdl/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// CRUD

// GetGameMode godoc
//
// @Summary Get all game's modes
// @Descripton Get all game's modes
// @ID get-game-modes
// @Tags GameMode
// @Success 200 {object} []GameMode
// @Failure 404 {object} ErrResponse
// @Router /games/{gameid}/modes/ [get]
func (config *Config) GetAll(w http.ResponseWriter, r *http.Request) {
	gameID := chi.URLParam(r, "gameid")

	var gameModes []models.GameMode
	config.Database.Model(&models.GameMode{}).Where("game_id=?", gameID).Find(&gameModes)

	render.JSON(w, r, gameModes)
}

// CreateGameMode godoc
// @Summary Create a game mode
// @Description Create a game mode
// @ID create-game-mode
// @Tags GameMode
// @Param request body GameModePostPutPayload true "game mode's info"
// @Success 201 {object} GameMode
// @Failure 400 {object} ErrResponse
// @Router /games/{gameid}/modes/ [post]
func (config *Config) Create(w http.ResponseWriter, r *http.Request) {
	// We bind the data
	data := &models.GameModePostPutPayload{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	// We check authorization
	gameID := chi.URLParam(r, "gameid")
	group := group.ForContext(r.Context())
	if group == nil {
		render.Render(w, r, errors.ErrUnauthorized("you must specify the group"))
		return
	}

	var game models.Game
	config.Database.First(&game, gameID)

	if game.GroupID != *group {
		render.Render(w, r, errors.ErrUnauthorized("this game doesn't belong to your group"))
		return
	}

	// we create the mode
	gameMode := models.GameMode{
		GameID: game.ID,
		Name:   *data.Name,
	}

	config.Database.Create(&gameMode)

	render.Status(r, 201)
	render.JSON(w, r, gameMode)
}
