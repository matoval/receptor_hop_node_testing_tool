package cmd

import (
	"fmt"
	"os"
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
	FileSender(socketPath, nodeId, fileName, tlsCert, tlsKey)
}
