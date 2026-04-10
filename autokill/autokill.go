// Package autokill provides the function KillOtherInstance, which terminates all other
// running instances of the fastlauncher_next application, leaving only the current process.
package autokill

import (
	"errors"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/probeldev/fastlauncher_next/bash"
)

func KillOtherInstance() error {
	currentPid := os.Getpid()

	pids, err := getPidsOtherAllInstance()
	if err != nil {
		return err
	}

	for _, pid := range pids {
		if pid == strconv.Itoa(currentPid) {
			continue
		}

		killInstance(pid)
	}

	return nil
}

func getPidsOtherAllInstance() ([]string, error) {
	currentOs := runtime.GOOS
	switch currentOs {
	case "darwin":
		pids := []string{}

		command := "pgrep fastlauncher_next"
		output, err := bash.RunCommand(command)
		if err != nil {
			return pids, err
		}
		outputArr := strings.Split(output, "\n")

		for _, o := range outputArr {
			if o == "" {
				continue
			}

			pids = append(pids, o)
		}

		return pids, nil

	case "linux":
		// TODO: Implementation is needed
		return []string{}, nil
	}
	return nil, errors.New("OS is not support")
}

func killInstance(pid string) error {
	command := "kill " + pid
	_, err := bash.RunCommand(command)
	if err != nil {
		return err
	}

	return nil
}
