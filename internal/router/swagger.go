package router

import (
	"net/http"

	"github.com/priyansh7parikh/file-upload-scan/internal/swagger"
)

func registerSwagger(mux *http.ServeMux) {
	swagger.Register(mux)
}
