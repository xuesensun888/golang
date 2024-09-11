package internal

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
	"syscall"
)

func Command(args ...string) (stdout, stderr string, exitcode int) {
	if args[0] == "bash" && args[1] == "-c" {
		args = args[2:]
	}

	cmdline := strings.Join(args, " ")

	c := exec.Command("bash", "-c", cmdline)

	var buffout, bufferr bytes.Buffer

	c.Stdout = &buffout

	c.Stderr = &bufferr

	_ = c.Run()

	exitcode = c.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()

	stdout = buffout.String()

	stderr = bufferr.String()

	return
}

func CommandContext(ctx context.Context, args ...string) (stdout, stderr string, exitcode int) {
	if args[0] == "bash" && args[1] == "-c" {
		args = args[2:]
	}
	cmdline := strings.Join(args, " ")

	c := exec.CommandContext(ctx, "bash", "-c", cmdline)

	var buffout, bufferr bytes.Buffer

	c.Stdout = &buffout

	c.Stderr = &bufferr

	_ = c.Run()

	exitcode = c.ProcessState.ExitCode()

	stdout = buffout.String()

	stderr = bufferr.String()

	return
}
