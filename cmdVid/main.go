package main

import (
	"fmt"
	"os"

	"github.com/opennota/phash"
)

func main() {
	h1, err := phash.VideoHash(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("h1: ", h1)
}
