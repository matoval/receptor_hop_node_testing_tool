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
	var tlsCert string
	var tlsKey string

	for i, arg := range args {
		switch arg {
		case "--socket":
			socketPath = args[i+1]
		case "--node":
			nodeId = args[i+1]
		case "--payload":
			fileName = args[i+1]
		case "--key":
			tlsKey = args[i+1]
		case "--cert":
			tlsCert = args[i+1]
		case "--help":
			helpFile, _ := os.ReadFile("./help/send_file_help.txt")
			fmt.Printf("%+s\n", helpFile)
			return
		}
	}
	//receptorctl --cert /etc/receptor/tls/receptor.crt --key /etc/receptor/tls/receptor.key --socket /tmp/control.sock work submit echopayload -p /tmp/test2 --node test1

	cmd := exec.Command("receptorctl", "--cert", tlsCert, "--key", tlsKey, "--socket", socketPath, "work", "submit", "echopayload", "--node", nodeId, "--payload", fileName)

	fmt.Printf("CMD: %v\n", cmd)

	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		return
	}

	fmt.Println(string(stdout))
}
