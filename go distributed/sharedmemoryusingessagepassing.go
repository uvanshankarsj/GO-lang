package main

import (
	"fmt"
	"sync"
)

type MemoryMessage struct {
	Operation string
	Address   int
	Value     int
	Result    chan<- int
}

func memoryManager(memoryChan <-chan *MemoryMessage, wg *sync.WaitGroup) {
	defer wg.Done()

	memory := make([]int, 100)

	for msg := range memoryChan {
		switch msg.Operation {
		case "read":
			msg.Result <- memory[msg.Address]
		case "write":
			memory[msg.Address] = msg.Value
		}
	}
}

func main() {
	memoryChan := make(chan *MemoryMessage)
	var wg sync.WaitGroup
	var inp int
	wg.Add(1)
	go memoryManager(memoryChan, &wg)
	fmt.Scan(&inp)
	writeMsg := &MemoryMessage{
		Operation: "write",
		Address:   0,
		Value:     inp,
	}
	memoryChan <- writeMsg

	readResult := make(chan int)
	readMsg := &MemoryMessage{
		Operation: "read",
		Address:   0,
		Result:    readResult,
	}
	memoryChan <- readMsg

	value := <-readResult
	fmt.Println("Read value from shared memory:", value)

	close(memoryChan)
	wg.Wait()
}