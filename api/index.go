package handler

import (
	"jting-resume/router"
	"net/http"
)

func VercelHandler(w http.ResponseWriter, r *http.Request) {
	router := router.Router()

	router.ServeHTTP(w, r)
}
