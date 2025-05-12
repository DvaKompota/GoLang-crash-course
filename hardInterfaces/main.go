package main

import (
	"fmt"
	"io"
	"os"
)

// Use `go run .\main.go ../readme.md` to print out the README file
func main() {
	fp := os.Args[1]

	// This is a simpler solution, but it doesn't use interfaces:
	// data, err := os.ReadFile(fp)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(data))

	f, err := os.Open(fp)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
