package main

import (
	"fmt"
	"sync"
)

type Message struct {
	Data string
}

func sender(msgChan chan<- *Message, wg *sync.WaitGroup) {
	defer wg.Done()
	var inp string;
	fmt.Scan(&inp)
	msg := &Message{
		Data: "Hello, receiver!",
	}

	msgChan <- msg
	close(msgChan)
}

func receiver(msgChan <-chan *Message, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range msgChan {
		if msg.Data == "Nil" {
			fmt.Println("NIL")
		} else {
			fmt.Println("Received message:", msg.Data)
			msg.Data = "Nil"
		}
	}
}

func main() {
	msgChan := make(chan *Message)
	var wg sync.WaitGroup

	wg.Add(2)
	go sender(msgChan, &wg)
	go receiver(msgChan, &wg)

	wg.Wait()
	// receiver(msgChan, &wg)
}
