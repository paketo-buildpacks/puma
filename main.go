package main

import (
	"os"

	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/scribe"
)

func main() {
	parser := NewGemfileParser()
	logger := scribe.NewLogger(os.Stdout)

	detect := Detect(parser)
	build := Build(logger)

	packit.Run(detect, build)
}
