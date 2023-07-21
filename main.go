package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Print("Enter Commit Message: ")
	reader := bufio.NewReader(os.Stdin)
	cmdString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	cmd1 := exec.Command("git", "add", ".")
	cmd2 := exec.Command("git", "commit", "-m", strings.TrimSpace(cmdString))
	cmd3 := exec.Command("git", "push", "origin", "dev_rafiq")
	cmd4 := exec.Command("git", "checkout", "master")
	cmd5 := exec.Command("git", "merge", "dev_rafiq")
	cmd6 := exec.Command("git", "push", "origin", "master")
	cmd7 := exec.Command("git", "checkout", "dev_rafiq")
	cmd1.Stdout = os.Stdout
	cmd1.Stderr = os.Stderr
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	cmd3.Stdout = os.Stdout
	cmd3.Stderr = os.Stderr
	cmd4.Stdout = os.Stdout
	cmd4.Stderr = os.Stderr
	cmd5.Stdout = os.Stdout
	cmd5.Stderr = os.Stderr
	cmd6.Stdout = os.Stdout
	cmd6.Stderr = os.Stderr
	cmd7.Stdout = os.Stdout
	cmd7.Stderr = os.Stderr
	cmd1.Run()
	cmd2.Run()
	cmd3.Run()
	cmd4.Run()
	cmd5.Run()
	cmd6.Run()
	cmd7.Run()

}
