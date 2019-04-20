package handler

import (
	"net/http"
)

// Static creates hanlder for static assets
func Static() http.Handler {
	fs := http.FileServer(http.Dir("../../web/static"))
	return http.StripPrefix("/static/", fs)
}
