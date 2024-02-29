package main

import (
	"flag"
	"os"
)

func main() {
	var port string

	flag.StringVar(&port, "port", "8080", "the port over which to serve")
	flag.Parse()

	os.Exit(0)
}
