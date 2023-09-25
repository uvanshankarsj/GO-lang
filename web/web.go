package main

import (
	"fmt"
	// "log"
	"net/http"
	// "net/url"
)

const myurl = "https://lco.dev:3000/testing?coursename=computersecurity&name=uvan"

func main() {
	nodemain()
	// fmt.Println("web requeset")
	// response, err := http.Get(myurl)
	// if err != nil {
	// 	fmt.Print("error")
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%T", response)
	// defer response.Body.Close()
	// databyte, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(databyte))

	// result2, err := url.Parse(myurl)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(result2.Scheme))
	// fmt.Println(string(result2.Host))
	// fmt.Println(string(result2.Path))
	// fmt.Println(string(result2.RawQuery))
	// fmt.Println(string(result2.Port()))

	// quarams := result2.Query()
	// fmt.Printf("%T\n", quarams)
	// fmt.Print(quarams["name"])
	// fmt.Print(quarams["coursename"])

	// for a, val := range quarams {
	// 	fmt.Println("asd",a,val)
	// }

	// 	partofurl := &url.URL{
	// 		Scheme: "https",
	// 		Host:   "www.google.com",
	// 		Path:   "search",
	// 		// RawPath: "car",
	// 		RawQuery: "q=car",
	// 	}

	// fmt.Println(partofurl)
}

func nodemain() {
	// r := mux.NewRouter()
	response, err := http.Get("http://lco.dev")
	if err != nil {
		fmt.Print("srr")
		fmt.Println(err)
	}
	fmt.Println("status code=", response.StatusCode)
	fmt.Println("status=", response.Status)
	// requestbody := strings.NewReader("{'coursename':'uvan'}")
	// response, err = http.Post("https://localhost:8009/post", "application/json", requestbody)
	// if err != nil {
	// 	fmt.Println("later")
	// 	panic(err)
	// }
	defer response.Body.Close()
	fmt.Print(response.ContentLength)
	// content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))

	// log.Fatal(http.ListenAndServe(":8000", r))
}
