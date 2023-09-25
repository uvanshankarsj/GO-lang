package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("working")
	r := mux.NewRouter()
	r.HandleFunc("/testing", greeter).Methods("POST")
	r.HandleFunc("/", serverhome).Methods("GET")
	r.HandleFunc("/2", serverhome2).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
func greeter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hey there mod user")
	requestbody := strings.NewReader("{'coursename':'uvan'}")
	response, err := http.Post("https://localhost:8009/post", "application/json", requestbody)
	if err != nil {
		fmt.Println("later")
		panic(err)
	}
	defer response.Body.Close()
}
func serverhome(w http.ResponseWriter, r *http.Request) {
	// w.req.body
	w.Write([]byte("<h1>welcom to HOme</h1>"))
	fmt.Printf("w: asd\n")
}
func serverhome2(w http.ResponseWriter, r *http.Request) {
	// w.req.body
	w.Write([]byte("<h1>welcom to HOme2</h1>"))
}
