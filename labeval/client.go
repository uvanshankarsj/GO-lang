package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item1 struct {
	Name string
	Data string
}

func main() {
	var reply Item1
	var db []Item1

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal(err)
	}

	a := Item1{"First", "first file"}
	b := Item1{"Second", "second file"}
	c := Item1{"Third", "third file"}

	client.Call("Testing.GetDB", "", &db)
	fmt.Println(db)

	client.Call("Testing.Createnewfile", a, &reply)
	fmt.Println(reply)

	client.Call("Testing.GetDB", "", &db)
	fmt.Println(db)

	err = client.Call("Testing.Createnewfile", b, &reply)
	fmt.Println(reply)

	client.Call("Testing.GetDB", "", &db)
	fmt.Println(db)

	client.Call("Testing.Createnewfile", c, &reply)
	fmt.Println(reply)

	client.Call("Testing.GetDB", "", &db)
	fmt.Println(db)

	client.Call("Testing.Printfile", Item1{"Second", "A new second item"}, &reply)
	fmt.Println(reply)

	go client.Call("Testing.Deletefile", b, &db)
	fmt.Println(db)

	client.Call("Testing.GetDB", "", &db)
	fmt.Println(db)

	client.Call("Testing.GetDB", "", &db)
	fmt.Println(db)

}
