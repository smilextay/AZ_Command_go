package main

import (
	"log"

	"github.com/smilextay/az_command_go/src/cmd"
)

func main() {

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd Execute err:%v", err)
	}
}
