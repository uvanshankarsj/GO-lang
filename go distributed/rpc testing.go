package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct{}
type Timeserver int64

func main() {
	timeserver := new(Timeserver)
	rpc.Register(timeserver)
	fmt.Printf("working")
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println(err)
	}

	http.Serve(listener, nil)
}

func (t *Timeserver) giverservertime(args *Args, reply *int64) {
	*reply = 1
}