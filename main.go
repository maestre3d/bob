package main

import (
	"github.com/maestre3d/bob/bin"
)

func main() {
	// Get bootstrapper
	bootstrapper := new(bin.Bootstrap)

	// Init CLI
	bootstrapper.InitCLI()
}
