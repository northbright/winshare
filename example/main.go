package main

import (
	"github.com/northbright/winshare"
	"log"
)

func main() {
	name := `\\server\IPC$`

	exitCode, output, err := winshare.CancelConnection(name)
	log.Printf("CancelConnection() exit code: %v, error: %v, output: %s\n", exitCode, err, output)
}
