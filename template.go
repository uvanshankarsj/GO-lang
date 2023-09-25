package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	name  string
	class string
}

type Testing int

var database []Item

func main() {
	testing := new(Testing)
	err := rpc.Register(testing)
	if err != nil {
		fmt.Println("error")
	}
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listener error", err)
	}
	http.Serve(listener, nil)

}

func (k *Testing) Getdata(item Item, reply *Item) error {
	fmt.Print("hi")
	return nil
}
