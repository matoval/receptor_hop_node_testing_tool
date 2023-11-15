package cmd

import (
	"fmt"
	"os"
)

func KubeDeploy(args []string) {
	var socketPath string
	var tlsCert string
	var tlsKey string

	for i, arg := range args {
		switch arg {
		case "--socket":
			socketPath = args[i+1]
		case "--key":
			tlsKey = args[i+1]
		case "--cert":
			tlsCert = args[i+1]
		case "--help":
			helpFile, _ := os.ReadFile("./help/kube_deploy_help.txt")
			fmt.Printf("%+s\n", helpFile)
			return
		}
	}
	KubeLauncher(socketPath, tlsCert, tlsKey)
}
