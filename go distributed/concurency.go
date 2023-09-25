package main

import (
	"fmt"
	"sync"
	"time"
)

func thread1(uv *sync.WaitGroup, ch chan int) {
	fmt.Println("das1")
	time.Sleep(2 * time.Second)
	ch <- 2
	uv.Done()
}

func thread2(uv *sync.WaitGroup, ch chan int) {
	fmt.Println("das2")
	ch <- 3
	ch <- 4
	x := <-ch
	fmt.Println("hi", x)
	uv.Done()
}

func thread3(uv *sync.WaitGroup, ch chan int) {
	fmt.Println("das3")
	ch <- 4
	uv.Done()
}

func thread4(uv *sync.WaitGroup, ch chan int) {
	fmt.Println("das4")
	ch <- 5
	uv.Done()
}
func main() {
	var uv sync.WaitGroup
	uv.Add(2)
	ch := make(chan int, 2)
	go thread1(&uv, ch)
	go thread2(&uv, ch)
	// go thread3(&uv, ch)
	// go thread4(&uv, ch)
	fmt.Print("waiting\n")
	// x = <-ch
	// fmt.Println("hi", x)
	// x = <-ch
	// fmt.Println("hi", x)
	// x = <-ch
	// fmt.Println("hi", x)
	uv.Wait()
}
