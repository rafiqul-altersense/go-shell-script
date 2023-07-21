```
func main() {
	reader := bufio.NewReader(os.Stdin)
	cmdString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	println("cmdString---->", cmdString)
	// cmdString = strings.TrimSuffix(cmdString, "\n")
	// cmd := exec.Command(cmdString)
	// cmd.Stderr = os.Stderr
	// cmd.Stdout = os.Stdout
	// cmd.Run()
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		cmdString = strings.TrimSuffix(cmdString, "\n")
		cmd := exec.Command(cmdString)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

```
