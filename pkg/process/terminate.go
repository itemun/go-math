package process

import (
	"log"
	"os"
	"syscall"

	"go.uber.org/zap"
)

// Terminate its application gratefully.
func Terminate() {
	pid := os.Getpid()

	proc, err := os.FindProcess(pid)
	if err != nil {
		log.Fatal(
			"failed to find process ID for terminate myself",
			zap.Error(err),
		)
	}

	if err = proc.Signal(syscall.SIGTERM); err != nil {
		log.Fatal(
			"failed to send SIGTERM signal to myself",
			zap.Int("PID", pid),
			zap.Error(err),
		)
	}
}
