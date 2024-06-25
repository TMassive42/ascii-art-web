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
		return "", fmt.Errorf("error opening banner file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) != 855 {
		return "", fmt.Errorf("the banner file %s has been modified", file.Name())
	}

	characters := [][]string{}
	for i := 0; i < len(lines); i += 9 {
		end := i + 9
		if end > len(lines) {
			end = len(lines)
		}
		characters = append(characters, lines[i:end])
	}

	input = strings.Replace(input, "\\n", "\n", -1)
	input = strings.Replace(input, "\\t", "    ", -1)

	var art strings.Builder
	handleLn(&art, input, characters)

	return art.String(), nil
}

func handleLn(art *strings.Builder, s string, b [][]string) {
	if s == "" {
		return
	}

	isAllNewline := true
	for _, char := range s {
		if char != '\n' {
			isAllNewline = false
			break
		}
	}
	if isAllNewline {
		count := strings.Count(s, "\n")
		for i := 0; i < count; i++ {
			art.WriteString("\n")
		}
		return
	}

	str := strings.Split(s, "\n")
	for _, char := range str {
		if char == "" {
			art.WriteString("\n")
			continue
		}
		printer(art, char, b)
	}
}

func printer(art *strings.Builder, s string, b [][]string) {
	for i := 1; i < 9; i++ {
		for _, char := range s {
			toPrint := char - 32
			if toPrint < 0 || int(toPrint) >= len(b) {
				art.WriteString("        ") // 8 spaces for unknown characters
				continue
			}
			art.WriteString(b[toPrint][i])
		}
		art.WriteString("\n")
	}
}
