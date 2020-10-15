package puma

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type GemfileParser struct{}

func NewGemfileParser() GemfileParser {
	return GemfileParser{}
}

func (p GemfileParser) Parse(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, fmt.Errorf("failed to parse Gemfile: %w", err)
	}
	defer file.Close()

	quotes := `["']`
	pumaRe := regexp.MustCompile(fmt.Sprintf(`gem %spuma%s`, quotes, quotes))
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := []byte(scanner.Text())
		if pumaRe.Match(line) {
			return true, nil
		}
	}

	return false, nil
}
