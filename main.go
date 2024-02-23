package main

import "github.com/th3shadowbroker/sse-example/server"

func main() {
	server.Start(uint16(8080))
}
