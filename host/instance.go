package main

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-ps"
)

func ensureSingleInstance() error {
	currentPID := os.Getpid()

	executable, err := os.Executable()
	if err != nil {
		return err
	}

	currentName := filepath.Base(executable)

	processes, err := ps.Processes()
	if err != nil {
		return err
	}

	for _, proc := range processes {
		if proc.Pid() == currentPID || proc.Executable() != currentName {
			continue
		}

		process, err := os.FindProcess(proc.Pid())
		if err != nil {
			continue
		}

		if err := process.Kill(); err != nil {
			continue
		}
	}

	return nil
}
