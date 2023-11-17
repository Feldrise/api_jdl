package group

import (
	"context"
	"net/http"

	"feldrise.com/jdl/config"
	"feldrise.com/jdl/models"
)

var GroupCtxKey = contextKey{"group"}

type contextKey struct {
	name string
}

func Middleware(c *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			groupCode := r.Header.Get("JDLGroupCode")

			if groupCode == "" {
				next.ServeHTTP(w, r)
				return
			}

			var group models.Group
			c.Database.Model(&models.Group{}).Where("code=?", groupCode).First(&group)

			if group.ID == 0 {
				http.Error(w, "invalid group", http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), GroupCtxKey, &group.ID)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *uint {
	raw, _ := ctx.Value(GroupCtxKey).(*uint)

	return raw
}
