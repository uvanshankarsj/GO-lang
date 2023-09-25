package main

import (
	"fmt"
	"net/http"

	"main.go/router"
)

func main() {
	fmt.Println("mongodbapi")
	r := router.Router()
	fmt.Println("ver is getting started")
	http.ListenAndServe(":4000", r)
	fmt.Println("listening on port 4000")
}
