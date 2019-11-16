package main

import (
	"OnBoardingPOC/internal/boot"
	"flag"
)

// This needs to be done as part of a bootstrap step when the process starts.
// The workers are supposed to be long running.

func main() {
	var mode string
	flag.StringVar(&mode, "m", "trigger", "Mode is worker or trigger.")
	flag.Parse()
	boot.Init(mode)
}
