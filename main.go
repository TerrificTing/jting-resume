package main

import (
	"fmt"
	"jting-resume/router"
	"net/http"
)

func main() {
	router := router.Router()

	fmt.Println("http://localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("server error:", err)
	}
}
