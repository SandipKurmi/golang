package main

import (
	"fmt"
	"time"
)

// send
func processNum(numChannel chan int){
	for num := range numChannel {
		fmt.Println(num)
		time.Sleep(time.Second)
	}
}

// receive
func sum(result chan int, nums ...int)  {
	total := 0
	for _, num := range nums {
		total += num
	}
	result <- total
}

// goroutines synchronizing
func task(done chan bool){
	defer func() {done <- true}()
	fmt.Println("task is processing")
}

func emailSender(emailChannel chan string, emailsSent chan bool) {
	defer func() {emailsSent <- true}()

	for email := range emailChannel {
		fmt.Println("sending email to",email)
		time.Sleep(time.Second)
	}

}

func main() {
	// channels
	messageChannel := make(chan string)

	messageChannel <- "Hello"

	msg := <- messageChannel

	fmt.Println(msg)

	// multiple go routines
	numChannel := make(chan int)

	go processNum(numChannel)

	for i := 0; i < 10; i++ {
		numChannel <- i
	}


	// multiple return values
	resultChannel := make(chan int)

	go sum(resultChannel, 1, 2, 3, 4, 5)

	res := <- resultChannel 

	fmt.Println(res)

	// wait for task to finish

	done := make(chan bool)

	go task(done)

	<- done

	fmt.Println("task is done")

	// buffered channels

	emailChannel := make(chan string, 100)

	emailsSent := make(chan bool)

	go emailSender(emailChannel, emailsSent)

	for i := 0; i < 5; i++ {
		emailChannel <- fmt.Sprintf("%d@a.com", i)
	}

	//this is important to close the channel

	close(emailChannel)
	<- done
	fmt.Println("sending emails")

	// multiple senders and multiple receivers

	channel1 := make(chan int, 2)

	channel2 := make(chan string, 2)

	go func ()  {
		channel1 <- 1
	}()

	go func ()  {
		channel2 <- "hello"
		channel2 <- "world"
	}()


	for i := 0; i < 2; i++ {
		select {
		case msg1 := <- channel1:
			fmt.Println("received from channel1", msg1)
		case msg2 := <- channel2:
			fmt.Println("received from channel2", msg2)
		}
	}
}