package main

import (
	"fmt"
	srv "github.com/codycollier/grip-go/pkg/server"
)

func main() {

	fmt.Println("[gripd] Starting server...")
	srv.StartGripServer()
	fmt.Println("[gripd] Exited.")

}
