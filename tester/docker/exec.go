package docker

import (
	"bytes"
	"os/exec"
)

// Executes a command in the docker container specified and returns the stdout and stderr output
func Exec(container string, command []string) (string, string, error) {
	args := append([]string{"exec", container}, command...)

	cmd := exec.Command("docker", args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return stdout.String(), stderr.String(), err
	}
	return stdout.String(), stderr.String(), nil
}
