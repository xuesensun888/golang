package shell

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func Command(args ...string) (stdout, stderr string, exitcode int, pid int) {
	if args[0] == "bash" && args[1] == "-c" {
		args = args[2:]
	}
	cmdline := strings.Join(args, " ")
	cmd := exec.Command("bash", "-c", cmdline)
	var buffout, bufferr bytes.Buffer
	cmd.Stdout = &buffout
	cmd.Stderr = &bufferr

	cmd.Start()

	pid = cmd.Process.Pid
	time.Sleep(2 * time.Second)

	process, _ := os.FindProcess(pid)
	err := process.Signal(os.Interrupt)
	if err != nil {
		log.Fatal(err)
	}
	_ = cmd.Wait()
	exitcode = cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	stdout = buffout.String()
	stderr = bufferr.String()
	return

}
