package gamecard

import (
	"net/http"

	"feldrise.com/jdl/errors"
	"feldrise.com/jdl/group"
	"feldrise.com/jdl/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// CRUD

// GetGameCard godoc
//
// @Summary Get all game's cards
// @Descripton Get all game's cards
// @ID get-game-cards
// @Tags GameCard
// @Success 200 {object} []GameCard
// @Failure 404 {object} ErrResponse
// @Router /games/{gameid}/cards/ [get]
func (config *Config) GetAll(w http.ResponseWriter, r *http.Request) {
	gameID := chi.URLParam(r, "gameid")

	var gameCards []models.GameCard
	config.Database.Model(&models.GameCard{}).Where("game_id=?", gameID).Find(&gameCards)

	render.JSON(w, r, gameCards)
}

// CreateGameCard godoc
// @Summary Create a game card
// @Description Create a game card
// @ID create-game-card
// @Tags GameCard
// @Param request body GameCardPostPutPayload true "game card's info"
// @Success 201 {object} GameCard
// @Failure 400 {object} ErrResponse
// @Router /games/{gameid}/cards/ [post]
func (config *Config) Create(w http.ResponseWriter, r *http.Request) {
	// We bind the data
	data := &models.GameCardPostPutPayload{}
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

	// we create the card
	gameCard := models.GameCard{
		GameID:  game.ID,
		Content: *data.Content,
	}

	config.Database.Create(&gameCard)

	render.Status(r, 201)
	render.JSON(w, r, gameCard)
}

// UpdateGameCard godoc
// @Summary update a game card
// @Description Update a game card
// @ID update-game-card
// @Tags GameCard
// @Param id path string true "The id of the card to update"
// @Param request body GameCardPostPutPayload true "game card's info"
// @Success 201 {object} GameCard
// @Failure 400 {object} ErrResponse
// @Router /games/{gameid}/cards/{id} [put]
func (config *Config) Update(w http.ResponseWriter, r *http.Request) {
	// We bind the data
	data := &models.GameCardPostPutPayload{}
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

	// we get the card
	cardID := chi.URLParam(r, "id")
	var gameCard models.GameCard
	config.Database.Find(&gameCard, cardID)

	if gameCard.ID == 0 {
		render.Render(w, r, errors.ErrNotFound())
		return
	}

	gameCard.Content = *data.Content

	config.Database.Save(gameCard)

	render.JSON(w, r, gameCard)
}
