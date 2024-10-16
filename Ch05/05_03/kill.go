package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func killServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	defer file.Close()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return errors.Wrap(err, "bad process ID")
	}

	// Simulate kill
	fmt.Printf("killing server with pid=%d\n", pid)

	if err := os.Remove(pidFile); err != nil {
		log.Printf("warning: cannot remove pid file - %s", err)
	}

	return nil
}

func main() {
	if err := killServer("server.pid"); err != nil {
		log.Fatalf("error: %s", err)
	}
}
