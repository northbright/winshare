package winshare

import (
	//"bytes"
	//"fmt"
	//"log"
	"os"
	"os/exec"
	"syscall"
)

var (
	cmdName = `cmd.exe`
)

// RunCmd runs a command and return exit code and combined output.
func RunCmd(cmd *exec.Cmd) (int, []byte, error) {
	var (
		ws syscall.WaitStatus
	)

	cmd.Env = append(os.Environ())
	output, err := cmd.CombinedOutput()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws = exitError.Sys().(syscall.WaitStatus)
			return ws.ExitStatus(), output, nil
		}
		return 0, output, err
	}
	ws = cmd.ProcessState.Sys().(syscall.WaitStatus)
	return ws.ExitStatus(), output, nil
}

// CancelConnection deletes the named connection.
func CancelConnection(name string) (int, []byte, error) {
	cmd := exec.Command(cmdName, "/c", "NET USE", name, "/DELETE")
	//cmd := exec.Command("cmd.exe", "/c", "net use", `\\server\IPC$`, `/DELETE`)
	return RunCmd(cmd)
}

func AddConnection(remoteName, localName, userName, password string) error {
	return nil
}

func ListSharedDisksAndPrinters(remotePC string) ([]string, []string, error) {
	return nil, nil, nil
}
