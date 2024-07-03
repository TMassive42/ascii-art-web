package asciiart

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GenerateASCIIArt(input string, banner string) (string, error) {
	bannerFile := fmt.Sprintf("%s.txt", banner)
	file, err := os.Open(bannerFile)
	if err != nil {
		return "", ErrNotFound
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) != 855 {
		return "", ErrBadRequest
	}

	characters := [][]string{}
	for i := 0; i < len(lines); i += 9 {
		end := i + 9
		if end > len(lines) {
			end = len(lines)
		}
		characters = append(characters, lines[i:end])
	}

	var art strings.Builder

	str := strings.Split(input, "\n")
	for _, char := range str {

		printer(&art, char, characters)
	}

	return art.String(), nil
}

func printer(art *strings.Builder, s string, b [][]string) {
	for i := 1; i < 9; i++ {
		for _, char := range s {
			toPrint := char - 32
			// handleLn(&art, input, characters)
			if toPrint < 0 || int(toPrint) >= len(b) {
				art.WriteString("        ") // 8 spaces for unknown characters
				continue
			}
			art.WriteString(b[toPrint][i])
		}
		art.WriteString("\n")
	}
}
