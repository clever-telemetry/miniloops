package main

import (
	"github.com/clever-telemetry/miniloops/pkg/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		panic(err.Error())
	}
}


