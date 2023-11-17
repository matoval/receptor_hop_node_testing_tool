package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

func ScalePayloads(args []string) {
	var nodeId string
	timeout := 5 * time.Minute

	for i, arg := range args {
		switch arg {
		case "--node":
			nodeId = args[i+1]
		case "--timeout":
			tempTimeout, err := strconv.Atoi(args[i+1])
			if err != nil {
				fmt.Println("INT required for timeout")
				helpFile, _ := os.ReadFile("./help/scale_payloads_help.txt")
				fmt.Printf("%+s\n", helpFile)
				return
			}
			timeout = time.Minute * time.Duration(tempTimeout)
		case "--help":
			helpFile, _ := os.ReadFile("./help/scale_payloads_help.txt")
			fmt.Printf("%+s\n", helpFile)
			return
		}
	}
	increasePayloads(nodeId, timeout)
}

func increasePayloads(nodeId string, timeout time.Duration) {
	var wg sync.WaitGroup
	keepAlive := true
	numberOfRuns := 1

	for keepAlive {
		for j := 0; j < numberOfRuns; j++ {
			wg.Add(1)
			go sendMessageScale(nodeId, numberOfRuns, &wg)
		}

		// Wait for replies or timeout
		if !waitWithTimeout(&wg, timeout) {
			fmt.Println("Timeout reached at %d number of runs.", numberOfRuns)
			keepAlive = false
		}
		numberOfRuns = numberOfRuns * 2
	}

}

func sendMessageScale(nodeID string, numberOfRuns int, wg *sync.WaitGroup) {
	defer wg.Done()
	socketPath := "/tmp/control.sock"
	tlsCert := "/etc/receptor/tls/receptor.crt"
	tlsKey := "/etc/receptor/tls/receptor.key"

	//sent with the -f option to get the returned string
	cmd := exec.Command("receptorctl", "--cert", tlsCert, "--key", tlsKey, "--socket", socketPath, "work", "submit", "cat", "--node", nodeID, "-l", strconv.FormatInt(int64(numberOfRuns), 10), "-f")

	fmt.Printf("CMD: %v\n", cmd)

	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		return
	}
	stdout = append(stdout, 0)
}

func waitWithTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()

	select {
	case <-c:
		return true
	case <-time.After(timeout):
		return false
	}
}
