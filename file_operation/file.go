package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("file operation")
	file, err := os.Create("testing.txt")
	if err != nil {
		fmt.Println(err)
	}
	content := "hi i am uvan shankar"
	io.WriteString(file, content)
	err = os.Chmod("testing.txt", 0777)
}
