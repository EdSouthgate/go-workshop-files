package fileutils

import "os"

// Public vars and functions are capitalised (pascal case) lower case are private (camel case)
func ReadTextFile(filename string) (string, error) { // in Go functions can return more than one value
	content, err := os.ReadFile(filename)
	if err != nil { // nil is like null in other languages
		// we couldn't read the file
		return "", err // returning a string here considered the 'lazy option'
	} else {
		// read operation successful
		return string(content), nil
	}
}
