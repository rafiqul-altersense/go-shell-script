package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Print("Enter command: ")
	reader := bufio.NewReader(os.Stdin)
	cmdString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	cmdString = strings.TrimSuffix(cmdString, "\n")
	cmd := exec.Command(cmdString) // exec.Command is for executing commands
	cmd.Stderr = os.Stderr         // os.Stderr is for error output, this coms from os package
	
	cmd.Stdout = os.Stdout         // os.Stdout is for standard output , this comes from os package
	var r_err = cmd.Run()                      // cmd.Run() is for running the command
	if r_err != nil {
		fmt.Fprintln(os.Stderr, r_err)
	}
}
