package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Name string
	Data string
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
	fmt.Print("listening at port 4040")
	http.Serve(listener, nil)

}

func (k *Testing) Createnewfile(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (k *Testing) Printfile(item Item, reply *Item) error {
	for idx, val := range database {
		if val.Name == item.Name {
			*reply = database[idx]
		}
	}
	return nil
}

func (k *Testing) Deletefile(item Item, reply *[]Item) error {
	for idx, val := range database {
		if val.Name == item.Name {
			database = append(database[:idx], database[idx+1:]...)
			*reply = database
		}
	}
	return nil
}

func (k *Testing) GetDB(empty string, reply *[]Item) error {
	*reply = database
	return nil
}
