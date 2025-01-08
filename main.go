package main

import (
	"fmt"
	"jting-resume/data"
	"jting-resume/render"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, data.ResumeData)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "styles.css")
	})

	fmt.Println("http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error:", err)
	}
}
