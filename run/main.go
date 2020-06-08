package main

import (
	"os"

	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/scribe"
	"github.com/paketo-community/puma"
)

func main() {
	parser := puma.NewGemfileParser()
	logger := scribe.NewLogger(os.Stdout)

	packit.Run(
		puma.Detect(parser),
		puma.Build(logger),
	)
}
