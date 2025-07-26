package main

import (
	"gitactivitytracker/cmd"
	"log"
)

func main() {
	log.Printf("Your cmd is getting executed from main func")
	cmd.Execute()
}
