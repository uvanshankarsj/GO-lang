package main
import (
	"fmt"
	//"time"
)

func main(){
    go fmt.Println("Go routine")
	//time.Sleep(time.Second)
    fmt.Println("Main Routine")
	
}