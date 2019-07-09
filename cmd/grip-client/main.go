package main

import (
	client "github.com/codycollier/grip-go/pkg/client"
	"log"
)

func main() {

	log.Println("[grip-client] Hello world!")

	// Get a grpc+grip client
	var addr = "127.0.0.1:8080"
	gripcl, conn := client.GetNewGripClient(addr)
	defer conn.Close()

	// ...
	client.CallEcho(gripcl, "foo", 0)
	client.CallCompute(gripcl, 3)

}
