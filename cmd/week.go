package cmd

import (
	"fmt"
	"os"
	"time"
)

func TimeTest(args []string) {

	var nodeId string

	for i, arg := range args {
		switch arg {
		case "--node":
			nodeId = args[i+1]
		case "--help":
			helpFile, _ := os.ReadFile("./help/send_file_help.txt")
			fmt.Printf("%+s\n", helpFile)
			return
		}
	}

	go runWithHourlyMessages(nodeId)

	// Sleep for a week to keep the main function alive
	time.Sleep(7 * 24 * time.Hour)
}

func runWithHourlyMessages(nodeId string) {
	// Create a ticker that ticks every hour
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Call the function to send a message every hour
			sendMessage(nodeId)
		}
	}
}

func sendMessage(nodeId string) {
	socketPath := "/tmp/control.sock"
	tlsCert := "/etc/receptor/tls/receptor.crt"
	tlsKey := "/etc/receptor/tls/receptor.key"

	MessageSender(socketPath, nodeId, tlsCert, tlsKey)
}
