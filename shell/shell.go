package shell

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func Exec(binary string, command string) (string, error) {
	args := strings.Fields(command)

	cmd := exec.Command(binary, args...)

	cmdOutput := &bytes.Buffer{}
	cmdError := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Stderr = cmdError
	err := cmd.Run()

	if err != nil {
		errorOutput := string(cmdError.Bytes())
		log.Fatal(errorOutput)
	}

	return string(cmdOutput.Bytes()), nil
}
