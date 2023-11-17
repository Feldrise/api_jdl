package group

import (
	"net/http"

	"feldrise.com/jdl/errors"
	"feldrise.com/jdl/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// GetGroupFromCode godoc
//
// @Summary Get a group from its code
// @Descripton Get a group from its code
// @ID get-group-from-code
// @Tags Group
// @Param code path string true "The code of te group to get"
// @Success 200 {object} Group
// @Failure 404 {object} ErrResponse
// @Router /groups/code/{code} [get]
func (config *Config) GetFromCode(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	var group models.Group
	config.Database.Model(&models.Group{}).Where("code=?", code).First(&group)

	if group.ID == 0 {
		render.Render(w, r, errors.ErrNotFound())
		return
	}

	render.JSON(w, r, group)
}
