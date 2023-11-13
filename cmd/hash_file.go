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
			return
		}
	}

	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("error reading file %v, error: %v\n", fileName, err)
		return
	}
	hashData := []byte(fileData)
	fmt.Printf("Hash: %x\n", md5.Sum(hashData))
}
