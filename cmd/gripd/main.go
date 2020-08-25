package main

import (
	"log"

	srv "github.com/codycollier/grip-go/pkg/server"
)

func main() {

	var addr = "127.0.0.1:8080"
	srv.StartGripServer(addr)
	log.Println("[gripd] Exited.")

}
