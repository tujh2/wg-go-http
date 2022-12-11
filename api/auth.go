package api

import (
	"net/http"

	"gopkg.in/macaron.v1"
)

func AuthMiddleware(secret string) macaron.Handler {
	return func(ctx *macaron.Context) {
		auth := ctx.Req.Header.Get("Authorization")

		if auth != secret {
			http.Error(ctx.Resp, "Not Authorized", http.StatusUnauthorized)
			return
		}
	}
}
