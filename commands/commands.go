package commands

import (
	"log"
	"regexp"

	"github.com/leia/TetriFast_server/blocks"
)

func ProcessCommand(command string) string {
	result := ""
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	command = reg.ReplaceAllString(command, "")

	switch command {
	case "getblock":
		result = blocks.GenerateBlock()
	}

	return result
}
