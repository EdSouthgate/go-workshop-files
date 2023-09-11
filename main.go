package main

import (
	"fmt"
	"os"

	fileutils "frontendmasters.com/go/files/filemanagement"
)

func main() {
	rootPath, _ := os.Getwd() // _ used when we don't care about the value here it would be the error if we can't get the working directory
	filePath := "/data/text.txt"
	contents, err := fileutils.ReadTextFile(rootPath + filePath) // := is a shorthand way of doing variable assignment.
	if err == nil {
		fmt.Println(contents)
		newContent := fmt.Sprintf("Original: %v\n Double Orig8inal %v%v", contents, contents, contents)
		fmt.Println(newContent)
		fileutils.WriteToFile(filePath+".output.txt", newContent)
	} else {
		fmt.Printf("Error panic!! %v", err)
	}
}
