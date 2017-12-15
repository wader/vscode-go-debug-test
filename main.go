package main

import (
	"log"
)

func main() {
	// set a breakpoint below and do "Debug in docker"
	v := "test"
	log.Printf("%#v", v)
}
