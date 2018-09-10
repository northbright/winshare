package winshare

import (
	//"bytes"
	"fmt"
	"os/exec"
)

var (
	cmdPath = `C:\Windows\System32\cmd.exe`
)

func CancelConnection(remoteName, localName string) error {
	cmd := exec.Command(cmdPath, "/c", "NET", "USE", localName, remoteName, "/DELETE")
	output, err := cmd.CombinedOutput()
	fmt.Printf("out: %s\n", output)
	if err != nil {
		return err
	}
	return nil
}

func AddConnection(remoteName, localName, userName, password string) error {
	return nil
}

func ListSharedDisksAndPrinters(remotePC string) ([]string, []string, error) {
	return nil, nil, nil
}
