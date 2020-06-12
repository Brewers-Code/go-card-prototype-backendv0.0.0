package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 300)
	}
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %dm", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

var wg sync.WaitGroup

func printStuff() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	// c := make(chan string)
	// go boring("Charles is da man!!", c)

	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("You say: %q\n", <-c)
	// }
	// fmt.Println("You are boring; I'm leaving")
	go printStuff()
	fmt.Println("You are boring; I'm leaving")

}

// go to enable a concurrent op
// use channels to sync by passing and reciving data
