package main

import (
	"log"

	"github.com/yann0917/go-tour-book/tour/cmd"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
