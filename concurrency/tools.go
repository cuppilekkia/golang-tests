package main

import "fmt"

func main() {
	c := make(chan int)
	go helloWorld(c)
	<-c
}

func helloWorld(c chan int) {
	fmt.Println("Hello")
	c <- 1
}
