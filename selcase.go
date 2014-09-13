package main

import (
	"fmt"
	"time"
)

func func1(c chan string) {
	for {
		c <- "from 1"
		time.Sleep(time.Second * 2)
	}
}

func func2(c chan string) {
	for {
		c <- "from 2"
		time.Sleep(time.Second * 3)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	timer := time.After(10 * time.Second)
	go func1(c1)
	go func2(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case <-timer:
			fmt.Println("Time up!! \n")
			return
		}
	}
}
