package cmd

import (
	"fmt"
	"os"
	"strconv"
)

func OneHundredTest(args []string) {
	var nodeId string
	helpFile, _ := os.ReadFile("./help/onehundred_help.txt")

	for i, arg := range args {
		switch arg {
		case "--number":
			nodeId = args[i+1]
		case "--help":
			fmt.Printf("%+s\n", helpFile)
			return
		}
	}
	num, err := strconv.Atoi(nodeId)
	if err != nil {
		fmt.Println("NUMBER NOT AN INT")
		fmt.Printf("%+s\n", helpFile)
	}

	go oneHundredTestRuns(num)

}

func oneHundredTestRuns(num int) {
	for i := 0; i < num; i++ {
		for j := 0; j < 10; j++ {
			nodeID := fmt.Sprintf("exi-node%d", i)
			sendMessage(nodeID)
		}
	}
}
