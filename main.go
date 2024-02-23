package main

import (
	"log"

	"github.com/th3shadowbroker/sse-example/client"
	"github.com/th3shadowbroker/sse-example/server"
)

func main() {
	var url = "http://localhost:8080/sse"

	go server.Start(8080)
	server.AwaitReadiness(url, 10, 1)

	var sse = client.NewSSEClient()
	go sse.Connect(url)

	for {
		message, ok := <-sse.Messages
		if !ok {
			return
		}
		log.Printf("Received message %s\n", message.Id)
	}
}
