package router

import (
	"jting-resume/handlers"
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	router.HandleFunc("/", handlers.ResumeHandler)

	return router
}
