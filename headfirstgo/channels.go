package main

import (
	"fmt"
	"time"
)

func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wake up!")
}

func send(myChannel chan string) {
	reportNap("sending gorutine", 2)
	fmt.Println("***sending value***")
	myChannel <- "a"
	fmt.Println("***sending value***")
	myChannel <- "b"
}

func main() {
	//var myChannel chan float64
	//myChannel = make(chan float64)
	thisChannel := make(chan string)
	//go greeting(myChannel)
	//fmt.Println(<-myChannel)
	//receivedValue := <-myChannel
	//fmt.Println(receivedValue, "dude!")
	//fmt.Println(receivedValue)
	go send(thisChannel)
	reportNap("receiving gorutine", 5)
	fmt.Println(<-thisChannel)
	fmt.Println(<-thisChannel)

}
