package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// array of commands
var commands = []string{
	"git add .",
	"git commit -m",
	"git push origin gitpush",
	"git checkout main",
	"git pull origin main",
	"git merge gitpush",
	"git push origin main",
	"git checkout gitpush",
}

func main() {
	fmt.Print("Enter Commit Message: ")
	reader := bufio.NewReader(os.Stdin)
	cmdString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	cmdString = strings.Replace(cmdString, "\n", "", -1)
	commands[1] = commands[1] + " " + `"` + cmdString + `"`
	for _, cmd := range commands {
		fmt.Println(cmd)
		command := exec.Command("bash", "-c", cmd)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		if err := command.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}
