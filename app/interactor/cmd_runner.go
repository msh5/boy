package interactor

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func runBytesAsCommand(b []byte, commandArgs []string) (int, error) {
	tempFile, err := writeToTempFile(b)
	if err != nil {
		return 0, err
	}

	defer os.Remove(tempFile.Name())

	if err := os.Chmod(tempFile.Name(), 0755); err != nil {
		return 0, err
	}

	command := exec.Command(tempFile.Name(), commandArgs...) //nolint:gosec
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err == nil {
		return 0, nil
	}

	if exitErr, ok := err.(*exec.ExitError); ok {
		return exitErr.ExitCode(), nil
	}

	return 0, err
}

func writeToTempFile(b []byte) (*os.File, error) {
	tempFile, err := ioutil.TempFile(os.TempDir(), "boy_*")
	if err != nil {
		return nil, err
	}

	_, err = tempFile.Write(b)

	tempFile.Close()

	if err != nil {
		os.Remove(tempFile.Name())
		return nil, err
	}

	return tempFile, nil
}
