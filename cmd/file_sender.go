package cmd

import (
	"fmt"
	"os/exec"
	"time"
)

func FileSender(socketPath string, nodeId string, fileName string, tlsCert string, tlsKey string) {
	start := time.Now()
	cmd := exec.Command("receptorctl", "--cert", tlsCert, "--key", tlsKey, "--socket", socketPath, "work", "submit", "echopayload", "--node", nodeId, "--payload", fileName)

	fmt.Printf("CMD: %v\n", cmd)

	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		return
	}
	fmt.Printf("TIme to send: %v", time.Since(start))
	fmt.Println(string(stdout))
}

func MessageSender(socketPath string, nodeId string, tlsCert string, tlsKey string) {

	cmd := exec.Command("receptorctl", "--cert", tlsCert, "--key", tlsKey, "--socket", socketPath, "work", "submit", "echopayload", "--node", nodeId, "-l", "\"Hello world!\"")

	fmt.Printf("CMD: %v\n", cmd)

	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		return
	}

	fmt.Println(string(stdout))
}
