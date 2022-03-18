package main

import (
	"os"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
	"github.com/paketo-buildpacks/puma"
)

func main() {
	parser := puma.NewGemfileParser()
	logger := scribe.NewEmitter(os.Stdout)

	packit.Run(
		puma.Detect(parser),
		puma.Build(logger),
	)
}
