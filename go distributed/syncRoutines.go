package main

import (
	"fmt"
	"sync"
)
var a=0

func main() {
	var wg sync.WaitGroup
	var mut sync.Mutex
	wg.Add(4)
	go foo(&wg,&mut)
	go foo1(&wg,&mut)
	go foo2(&wg,&mut)
	go foo3(&wg,&mut)
	go foo4(&wg,&mut)
	wg.Wait()
	fmt.Println("Main Routine")

}


func foo(wg *sync.WaitGroup,mut *sync.Mutex) {
	fmt.Println("New Routine 1")
	mut.Lock()
	a=2
	mut.Unlock()
	wg.Done()
}
func foo1(wg *sync.WaitGroup,mut *sync.Mutex) {
	fmt.Println("New Routine 2")
	mut.Lock()
	a=22
	mut.Unlock()
	wg.Done()
}
func foo2(wg *sync.WaitGroup,mut *sync.Mutex) {
	fmt.Println("New Routine 3")
	mut.Lock()
	a=24
	mut.Unlock()
	wg.Done()
}
func foo3(wg *sync.WaitGroup,mut *sync.Mutex) {
	fmt.Println("New Routine 4")
	mut.Lock()
	a=12
	mut.Unlock()
	wg.Done()
}
func foo4(wg *sync.WaitGroup,mut *sync.Mutex) {
	fmt.Println("New Routine 5")
	mut.Lock()
	a=287
	mut.Unlock()
	wg.Done()
}

