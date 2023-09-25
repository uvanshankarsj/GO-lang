package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Name  string
	Email string
}

func main() {
	var v string = "asd"
	var arr = [4]string{"uvan", "shankar", "s.j"}
	// arr[0] = "uvan"
	var slice = []string{}
	slice = append(slice, "uvan")

	maps := make(map[string]string)
	maps["uv"] = "uvan"
	maps["sh"] = "shankar"
	fmt.Print(maps)
	delete(maps, "sh")
	fmt.Print(maps)
	fmt.Print(slice)
	fmt.Print(arr)
	fmt.Println("New Routine 1", v)
	fmt.Printf("type: %T  ", v)
	reader := bufio.NewReader(os.Stdin)
	a, _ := reader.ReadString('\n')
	ans, err := strconv.ParseFloat(strings.TrimSpace(a), 64)
	if err == nil {
		fmt.Println(ans)
	}
	fmt.Printf("%T", a)
	presenttime := time.Now()
	fmt.Println(presenttime.Format("01-02-2006 Monday 15:04:05"))
	uvan := User{"uvan", "uvanshnakr02@gmail.com"}
	fmt.Printf("%+v", uvan)
	fmt.Printf("Email is %v", uvan.Email)

	no := rand.Intn(6) + 1
	switch no {
	case 1:
		fmt.Println("hi1")
		fallthrough
	case 2:
		fmt.Println("hi2")
	case 3:
		fmt.Println("hi3")
		fallthrough
	case 4:
		fmt.Println("hi4")
	case 5:
		fmt.Println("hi5")
	case 6:
		fmt.Println("hi6")
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			// goto ans
			continue
		}
		fmt.Print(i)
	}
	ansd, msgd := adder(1, 2, 3, 4)
	fmt.Println("the sum is ", ansd)
	fmt.Println("the sum is ", msgd)

	// ans:
	//
	//	fmt.Println("this is the no i am searching for")
	uvan.getstatus()
}

func adder(values ...int) (int, string) {
	total := 0
	for i := range values {
		total += i
	}
	return total, "i am uvan"
}

func (u User) getstatus() {
	defer fmt.Print(u.Name)
	fmt.Print("my name is")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
