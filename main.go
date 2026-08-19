package main

import (
	"fmt"
	"os"
)

// identity_verifier - Verify user identity
func identity_verifier(path string) {
	fmt.Println("========================================")
	fmt.Println("  Identity-Verifier")
	fmt.Println("  Verify user identity")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	identity_verifier(path)
}
