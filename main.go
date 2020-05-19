package main

import "github.com/paketo-buildpacks/packit"

func main() {
	detect := Detect()
	build := Build()

	packit.Run(detect, build)
}
