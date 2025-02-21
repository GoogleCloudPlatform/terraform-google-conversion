package main

import (
	"flag"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/cmd"
)

func main() {
	// Workaround for "ERROR: logging before flag.Parse"
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	_ = fs.Parse([]string{})
	flag.CommandLine = fs

	cmd.Execute()
}
