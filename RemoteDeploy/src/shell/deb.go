package shell

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
	cmd := exec.Command("bash", "-c", cmdline)

	var buffout, bufferr bytes.Buffer

	cmd.Stdout = &buffout
	cmd.Stderr = &bufferr

	_ = cmd.Run()
	exitcode = cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	stdout = buffout.String()
	stderr = bufferr.String()
	return
}

func CommandContext(ctx context.Context, args ...string) (stdout, stderr string, exitcode int) {
	if args[0] == "bash" && args[1] == "-c" {
		args = args[2:]
	}
	cmdline := strings.Join(args, " ")
	cmd := exec.CommandContext(ctx, "bash", "-c", cmdline)

	var buffout, bufferr bytes.Buffer

	cmd.Stdout = &buffout
	cmd.Stderr = &bufferr

	_ = cmd.Run()
	exitcode = cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	stdout = buffout.String()
	stderr = bufferr.String()
	return
}
