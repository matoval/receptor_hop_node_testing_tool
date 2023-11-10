package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func SendFile(args []string) {
	var socketPath string
	var nodeId string
	var fileName string

	for i, arg := range args {
		switch arg {
		case "--socketPath":
			socketPath = args[i+1]
		case "--nodeId":
			nodeId = args[i+1]
		case "--fileName":
			fileName = args[i+1]
		case "--help":
			helpFile, _ := os.ReadFile("./help/send_file_help.txt")
			fmt.Printf("%+s\n", helpFile)
		}
	}

	cmd := exec.Command("receptorctl", fmt.Sprintf("--socket %v work submit echopayload --node %v --payload %v", socketPath, nodeId, fileName))

	fmt.Printf("CMD: %v\n", cmd)

	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		return
	}

	fmt.Println(string(stdout))
}
