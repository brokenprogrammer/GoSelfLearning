//###########################################
//# Understanding Concurrency in Go
//# Go Routines:
//# https://golang.org/doc/effective_go.html#concurrency
//# https://www.golang-book.com/books/intro/10#section1
//#
//# Channels:
//# http://2.bp.blogspot.com/-RlFbVyp66qw/TfFQDBF2JtI/AAAAAAAABIM/ArTxswQ0GeE/s400/channels.png
//# http://golangtutorials.blogspot.rs/2011/06/channels-in-go.html
//###########################################
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/*
	A Function that will take in a string which we will use to identify which goroutine is which
	and loops 3 times printing out the current loop it is on and then sleeping for a random time.
*/
func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, ":", i+1)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	fmt.Println("In Main")
	f("direct") //Normal direct call to the function

	//Making the calls as Go Routines
	go f("goroutine 1")
	go f("goroutine 2")
	go f("goroutine 3")
	go f("goroutine 4")

	//Creating a function and making it a Go Routine at the spot
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")

	//Channels are a communication gateway for goroutines to communicate with one another.
	var c chan string = make(chan string)

	go pinger(c)  //This function will send a string "ping" into the channel
	go ponger(c)  //This function will take turns with pinger to send "pong" into the channel
	go printer(c) //This function will forever print the content of the channel and then sleep for 1 second.
	/*
		Using a channel like this synchronizes the two goroutines.
		When pinger attempts to send a message on the channel it will wait until printer is ready to receive the message.
		(this is known as blocking)

		 Each side - the sender and the receiver - is communicating one item at a time,
		 and has to wait until the other side performs either a sending or a receiving correspondingly.

		 The code here runs within a single thread, line by line, successively.
		 The operation of writing to the channel (c <- 42) blocks the execution of the whole program because,
		 as we remember, writing operations on a synchronous channel can only succeed in case there is a receiver
		 ready to get this data.
	*/
	fmt.Scanln(&input)
	fmt.Println("Done Again")

	//Baseball factory experimental
	var bb chan string = make(chan string)

	go baseballFactory(bb)
	go baseballStore(bb)

	fmt.Scanln(&input)

	c1 := make(chan string)
	c2 := make(chan string)

	//Printing out "from 2" every 3 seconds
	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	//Printing out "from 2" every 3 seconds
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			/*
				Select is like a switch statement for channels
				This select statement below picks the first channel available and recieves from it.
				If noone is available we will timeout. We could also create a default case.
			*/
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("Timeout")
			}
		}
	}()

	fmt.Scanln(&input)
	fmt.Println("Done 3rd Time")
}

//A function taking in a channel can be restricted to for example only sending by adding <- to the chan keyword
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

//A function printing out the content of the channel c. adding <- infront of the chan keyword means it can only recieve.
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func baseballFactory(c chan string) {
	for i := 0; i < 10; i++ {
		baseballName := "Baseball nr: " + strconv.Itoa(i)
		fmt.Println("Creating baseball: ", baseballName)
		c <- baseballName
	}
}

func baseballStore(c chan string) {
	for i := 0; i < 10; i++ {
		s := <-c
		fmt.Println("Putting ", s, " on the shelf")
		time.Sleep(time.Second * 1)
	}
}
