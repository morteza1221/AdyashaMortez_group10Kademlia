package main

import (
	"fmt"
	"os"

	"kademlia"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s port\n", os.Args[0])
		os.Exit(1)
	}

	kademlia.Listen(os.Args[1])
}
