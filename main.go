package main

import "github.com/paketo-buildpacks/packit"

func main() {
	parser := NewGemfileParser()

	detect := Detect(parser)
	build := Build()

	packit.Run(detect, build)
}
