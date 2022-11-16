package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/go-version"
)

func main() {
	if len(os.Args) != 3 {
		panic("wrong number of arguments, must supply a two semantic version strings to comparse")
	}

	v1, err := version.NewVersion(os.Args[1])

	if err != nil {
		panic(fmt.Sprintf("failed to parse version: %v", os.Args[1]))
	}

	v2, err := version.NewVersion(os.Args[2])

	if err != nil {
		panic(fmt.Sprintf("failed to parse version: %v", os.Args[2]))
	}

	fmt.Println(v1.LessThan(v2))
}
