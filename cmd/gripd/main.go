package main

import (
	srv "github.com/codycollier/grip-go/pkg/server"
	"log"
)

func main() {

	log.Println("[gripd] Starting server...")
	srv.StartGripServer()
	log.Println("[gripd] Exited.")

}
