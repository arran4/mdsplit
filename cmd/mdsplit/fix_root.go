//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"os"
)

func main() {
	// Look for root.go in the current directory or cmd/mdsplit/root.go
	path := "root.go"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		path = "cmd/mdsplit/root.go"
		if _, err := os.Stat(path); os.IsNotExist(err) {
			// If not found in either, assume we are in cmd/mdsplit but the file is not there, or running from root
			// Let's try to be robust. If the file doesn't exist, we can't patch it.
			panic("root.go not found")
		}
	}

	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	content = bytes.Replace(content, []byte("err := Run("), []byte("err := mdsplit.Run("), -1)
	if err := os.WriteFile(path, content, 0644); err != nil {
		panic(err)
	}
}
