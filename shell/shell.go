package shell

import (
	"bytes"
	"errors"
	"os/exec"
)

func Exec(binary string, command string) (string, error) {
	args := getArguments(command)

	cmd := exec.Command(binary, args...)

	cmdOutput := &bytes.Buffer{}
	cmdError := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Stderr = cmdError
	err := cmd.Run()

	if err != nil {
		errorOutput := string(cmdError.Bytes())
		return "", errors.New(errorOutput)
	}

	return string(cmdOutput.Bytes()), nil
}

func getArguments(command string) []string {
	args := []string{}
	arg := ""
	seenQuote := false

	for _, char := range command {
		c := string(char)

		if c == "\"" {
			seenQuote = !seenQuote
			c = ""
		}

		if c == " " && !seenQuote {
			args = append(args, arg)
			arg = ""
		} else {
			arg = arg + c
		}
	}

	if arg != "" {
		args = append(args, arg)
	}

	return args
}
