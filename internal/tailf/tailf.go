// This package implements utilities to stream tail -f output with golang
package tailf

import (
	"bufio"
	"os/exec"
)

// Streams tail -f output
func Tailf(path string, onReceive func(string)) error {
	cmd := exec.Command("tail", "-f", path)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	err = cmd.Start()

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		onReceive(line)
	}

	cmd.Wait()
	return nil
}
