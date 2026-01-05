package router

import "net/http"

func New() http.Handler {
	mux := http.NewServeMux()

	registerAPIRoutes(mux)

	registerSwagger(mux)

	return mux
}
