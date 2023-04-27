package shell

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	"strings"
)



func Run(ctx context.Context, command string) error {

	commandParsed := strings.Split(command, " ")
	cmd := exec.CommandContext(ctx, commandParsed[0], commandParsed[1:]...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return errors.New(stderr.String())
	}
	return nil
}

func Output(ctx context.Context, command string) (string, error) {
	commandParsed := strings.Split(command, " ")
	cmd := exec.CommandContext(ctx, commandParsed[0], commandParsed[1:]...)

	var stdout, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return stdout.String(), errors.New(stderr.String())
	}

	return stdout.String(), nil
}
