package ascii

import (
	"log"
	"os"
	"strings"
)

var (
	content []byte
	err     error
	Replace map[rune]([]string)
	Lines   []string
)

func Populate() map[rune]([]string) {
	Replace = map[rune]([]string){}
	_, _, Banner = Input()
	content, err = os.ReadFile("ascii/Banner/" + Banner)
	if err != nil {
		log.Fatal("Error : couldn't read file ", err)
	}

	Char := 32
	noreturn := strings.ReplaceAll(string(content), "\r", "")
	Lines := strings.Split(noreturn, "\n")
	for i := 0; i < len(Lines); i += 9 {
		if i+9 <= len(Lines)-1 {
			Replace[rune(Char)] = Lines[i+1 : i+9]
		}
		if Char <= 126 {
			Char++
		}

	}
	return Replace
}
