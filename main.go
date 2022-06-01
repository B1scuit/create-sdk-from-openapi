package main

import (
	"fmt"
	"os"

	"github.com/B1scuit/create-sdk-from-openapi/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
