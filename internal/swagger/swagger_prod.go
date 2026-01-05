//go:build !dev
// +build !dev

package swagger

import "net/http"

func Register(_ *http.ServeMux) {
	// no-op
}
