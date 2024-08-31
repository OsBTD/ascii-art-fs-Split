package ascii

import (
	"log"
	"os"
	"strings"
)

var (
	Banner     string
	input      string
	inputsplit []string
)

func Input() ([]string, string, string) {
	args := os.Args[1:]

	if len(args) == 0 || len(args) > 2 {
		log.Fatal("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
	} else if len(args) == 1 {
		args = append(args, "standard.txt")
	}
	Banner = "standard.txt"
	if len(args) == 2 {
		Banner = args[1]
	} else if len(args) == 1 {
		args = append(args, Banner)
	}
	input = args[0]
	if !strings.HasSuffix(args[1], ".txt") {
		args[1] += ".txt"
	}
	if args[1] != "standard.txt" && args[1] != "thinkertoy.txt" && args[1] != "shadow.txt" {
		log.Fatal("Usage: this style is unavailable \nPlease choose one of the available styles \n1 : standard \n2 : thinkertoy \n3 : shadow")
	}

	inputsplit = strings.Split(input, "\\n")
	for _, line := range inputsplit {
		for _, c := range line {
			if c < 32 || c > 126 {
				log.Fatal("Error : input should only contain printable ascii characters")
			}
		}
	}

	return inputsplit, input, Banner
}
