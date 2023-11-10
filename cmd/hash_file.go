package cmd

import (
	"crypto/md5"
	"fmt"
	"os"
)

func HashFile(args []string) {
	var fileName string

	for i, arg := range args {
		switch arg {
		case "--fileName":
			fileName = args[i+1]
		case "--help":
			helpFile, _ := os.ReadFile("./help/hash_file_help.txt")
			fmt.Printf("%+s\n", helpFile)
		}
	}

	fileData, _ := os.ReadFile(fileName)
	hashData := []byte(fileData)
	fmt.Printf("Hash: %x\n", md5.Sum(hashData))
}
