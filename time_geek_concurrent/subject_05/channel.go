package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n", elem1)
}

func timeChannel() {
	chan1 := make(chan int, 1)
	chan2 := make(chan int, 1)
	chan3 := make(chan int, 1)
	chan4 := make(chan int, 1)
	go func() {
		for {
			fmt.Println("I'm goroutine 1")
			time.Sleep(1 * time.Second)
			chan2 <- 1 //I'm done, you turn
			<-chan1
		}
	}()
	go func() {
		for {
			<-chan2
			fmt.Println("I'm goroutine 2")
			time.Sleep(1 * time.Second)
			chan3 <- 1 //I'm done, you turn
		}
	}()
	go func() {
		for {
			<-chan3
			fmt.Println("I'm goroutine 3")
			time.Sleep(1 * time.Second)
			chan4 <- 1 //I'm done, you turn
		}
	}()
	go func() {
		for {
			<-chan4
			fmt.Println("I'm goroutine 4")
			time.Sleep(1 * time.Second)
			chan1 <- 1 //I'm done, you turn
		}
	}()

}
