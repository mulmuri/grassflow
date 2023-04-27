package shell

import (
	"context"
	"os/exec"
	"strings"
)



func run(ctx context.Context, command string) error {

	commandParsed := strings.Split(command, " ")
	cmd := exec.CommandContext(ctx, commandParsed[0], commandParsed[1:]...)
	return cmd.Run()
}

func output(ctx context.Context, command string) (string, error) {
	commandParsed := strings.Split(command, " ")
	cmd := exec.CommandContext(ctx, commandParsed[0], commandParsed[1:]...)
	output, err := cmd.Output()
	return string(output), err
}
