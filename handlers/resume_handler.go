package handlers

import (
	"jting-resume/data"
	"jting-resume/render"
	"net/http"
)

func ResumeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, data.ResumeData)
}
