package utils

import (
	"os"
	"syscall"
)

func KillProcess(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	err = syscall.Kill(-process.Pid, syscall.SIGKILL)
	if err != nil {
		return err
	}

	return nil
}
