package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	start := time.Now()

	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //starts a go routine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) //receives from ch
	}

	fmt.Printf("time taken: %f", time.Since(start).Seconds())

	channelLearning()

}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("while fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %s %s", secs, resp.Status, url)
}

//

// learning channel function with simplicity
func channelLearning() {
	ch := make(chan string) // create a channel for strings

	go sendMessage(ch) // start a goroutine to send a message

	msg := <-ch // receive the message from the channel
	fmt.Println("Received:", msg)
}

func sendMessage(ch chan<- string) {
	ch <- "Hello from the channel!" // send a message into the channel
}
