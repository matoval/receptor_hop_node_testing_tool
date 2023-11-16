package main

import (
	"example/receptor/receptor_hop_node_testing_tool/cmd"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	switch {
	case strings.Contains(args[1], "create-file"):
		cmd.CreateFile(args)
	case strings.Contains(args[1], "hash-file"):
		cmd.HashFile(args)
	case strings.Contains(args[1], "send-file"):
		cmd.SendFile(args)
	case strings.Contains(args[1], "week-test"):
		cmd.TimeTest(args)
	case strings.Contains(args[1], "kube-deploy"):
		cmd.KubeDeploy(args)
	case strings.Contains(args[1], "onehundred"):
		cmd.OneHundredTest(args)
	case strings.Contains(args[1], "scalepayloads"):
		cmd.ScalePayloads(args)
	default:
		helpFile, _ := os.ReadFile("./help/general_help.txt")
		fmt.Printf("%+s\n", helpFile)
	}
}
