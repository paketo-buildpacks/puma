package main

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

func (p GemfileParser) Parse(path string) (bool, bool, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, false, nil
		}
		return false, false, fmt.Errorf("failed to parse Gemfile: %w", err)
	}
	defer file.Close()

	quotes := `["']`
	mriRe := regexp.MustCompile(`^ruby .*`)
	pumaRe := regexp.MustCompile(fmt.Sprintf(`^gem %spuma%s`, quotes, quotes))

	hasMri := false
	hasPuma := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []byte(scanner.Text())
		if hasMri == false {
			hasMri = mriRe.Match(line)
		}
		if hasPuma == false {
			hasPuma = pumaRe.Match(line)
		}
	}

	return hasMri, hasPuma, nil
}
