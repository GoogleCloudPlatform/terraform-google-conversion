package main

import (
	"flag"
	"log"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/cmd"
)

func main() {
	// Workaround for "ERROR: logging before flag.Parse"
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	_ = fs.Parse([]string{})
	flag.CommandLine = fs

	log.Printf("test")
	cmd.Execute()
}
