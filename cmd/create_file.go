package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

func CreateFile(args []string) {
	fileName := "test"
	fileSize := 100
	var fileMode fs.FileMode = 0666
	for i, arg := range args {
		switch arg {
		case "--fileName":
			fileName = args[i+1]
		case "--fileSize":
			fileSize, _ = strconv.Atoi(args[i+1])
		case "--fileMode":
			newFileMode, _ := strconv.ParseUint(args[i+1], 8, 32)
			fileMode = fs.FileMode(newFileMode)
		case "--help":
			helpFile, _ := os.ReadFile("./help/create_file_help.txt")
			fmt.Printf("%+s\n", helpFile)
		}
	}

	bigBuff := make([]byte, fileSize)
	err := os.WriteFile(fileName, bigBuff, fs.FileMode(fileMode))
	if err != nil {
		fmt.Printf("CreateFile returned error: %v", err)
	}

}
