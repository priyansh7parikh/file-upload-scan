//go:build dev
// +build dev

package swagger

import (
	"net/http"

	_ "github.com/priyansh7parikh/file-upload-scan/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Register(mux *http.ServeMux) {
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
}
