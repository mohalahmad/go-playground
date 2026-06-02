package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// for _, link := range links {
	// 	checkLink(link)
	// } // this is a blocking call, so we have to wait for each one to finish before moving on to the next one. seliarly, this is not ideal, especially if we have a lot of links to check. we can use goroutines to check the links concurrently.
	//
	// below is the same code as above but with go routines and channels
	// make requests concurrently, note that concurrency is not the same as parallelism, concurrency is when we have multiple tasks that can run at the same time, but they are not necessarily running at the same time, while parallelism is when we have multiple tasks that are running at the same time. in this case, we are using concurrency to check the links, but we are not necessarily running them in parallel.
	// if we used multiple cores, we could run them in parallel, but for now we will just use concurrency to check the links. we can use channels to communicate between the go routines and the main function, so that we can know when the checks are done.
	//

	c := make(chan string) // this is a channel that can be used to communicate between the go routines and the main function. we can send messages to the channel from the go routines, and we can receive messages from the channel in the main function.

	for _, link := range links {
		go checkLink(link, c) // this is a non-blocking call, so we can move on to the next one immediately. however, we have no way of knowing when the checks are done, so we need to use channels to communicate between the go routines and the main function.
	}
	// we can call println(<-c) here to receive a message from the channel, but we have no way of knowing how many messages will be sent to the channel, so we need to use a for loop to receive messages from the channel until we have received all the messages that we expect to receive. if we just call println(<-c) here, we will only receive one message from the channel, and we will not know when the checks are done, so we need to use a for loop to receive messages from the channel until we have received all the messages that we expect to receive.
	// fmt.Println(<-c) // this is a blocking call, so we have to wait for a message to be sent to the channel before we can move on. however, we have no way of knowing how many messages will be sent to the channel, so we need to use a for loop to receive messages from the channel until we have received all the messages that we expect to receive.

	// or we can use a for loop to receive messages from the channel until we have received all the messages that we expect to receive. we can use the length of the links slice to know how many messages we need to receive from the channel, since we are sending one message to the channel for each link that we check.

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	// blow will do same but infinite loop, we can use this approach if we don't know how many messages will be sent to the channel, but we need to make sure that we have a way to break out of the loop when we have received all the messages that we expect to receive, otherwise we will end up in an infinite loop.
	println("Now checking the links again...")
	//

	for _, link := range links {
		go checkLinkLink(link, c)
	}
	for l := range c { // infinite loop, we will keep receiving messages from the channel until we have received all the messages that we expect to receive, but we need to make sure that we have a way to break out of the loop when we have received all the messages that we expect to receive, otherwise we will end up in an infinite loop.
		// for  { // same loop also infinite loop, but not using the range keywork so not clear

		// time.Sleep(5 * time.Second) // wrong place for the pause while it will block the channels
		// go checkLinkLink(l, c) // just if did not use the literal below

		// go func() {
		// 	time.Sleep(1 * time.Second)
		// 	checkLinkLink(l, c)
		// }()
		// this will only work with Go 1.22+ while the loop variable have per-iteration scope
		// for older version use below
		// By the time the goroutine wakes up after 5 seconds, l has already been overwritten by the next iteration of the loop. All goroutines end up using whatever l was last assigned — likely the same URL repeated, or a garbage value after the loop ends.
		go func(link string) {
			time.Sleep(1 * time.Second)
			checkLinkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // while we don't care about the response, we want to know if there was an error
	if err != nil {
		log.Println(link, "might be down!")
		c <- "Might be down!" // send a message to the channel to indicate that the link is down, we can use this message in the main function to know when the checks are done.
		return
	}
	fmt.Println(link, "is up!")
	c <- "Yep, it's up!" // send a message to the channel to indicate that the link is up, we can use this message in the main function to know when the checks are done.
}

func checkLinkLink(link string, c chan string) {
	// time.Sleep(5 * time.Second) // this a valid place for the pasue but not logical while its expected make the call once the function is called, if we put the pause here, it will make the call after the pause, which is not what we want, we want to make the call immediately and then wait for the response, so we should put the pause after the call, but before we send the message to the channel, so that we can simulate a delay in the response time of the link, which is more realistic.
	_, err := http.Get(link) // while we don't care about the response, we want to know if there was an error
	if err != nil {
		log.Println(link, "might be down!")
		c <- link // send a message to the channel to indicate that the link is down, we can use this message in the main function to know when the checks are done.
		return
	}
	fmt.Println(link, "is up!")
	c <- link // send a message to the channel to indicate that the link is up, we can use this message in the main function to know when the checks are done.
}
