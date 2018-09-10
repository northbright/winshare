package winshare_test

import (
	"github.com/northbright/winshare"
	"log"
)

func ExampleCancelConnection() {
	remoteName := `\\server\IPC$`
	localName := ``

	if err := winshare.CancelConnection(remoteName, localName); err != nil {
		log.Printf("CancelConnection() error: %v", err)
		return
	}
	// Output:
}
