package main

import (
	"fmt"
	"path/filepath"

	"github.com/paketo-buildpacks/packit"
)

//go:generate faux --interface Parser --output fakes/parser.go
type Parser interface {
	Parse(path string) (hasMri, hasPuma bool, err error)
}

type BuildPlanMetadata struct {
	Build  bool `toml:"build"`
	Launch bool `toml:"launch"`
}

func Detect(gemfileParser Parser) packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		hasMri, hasPuma, err := gemfileParser.Parse(filepath.Join(context.WorkingDir, "Gemfile"))
		if err != nil {
			return packit.DetectResult{}, fmt.Errorf("failed to parse Gemfile: %w", err)
		}

		if !hasMri || !hasPuma {
			return packit.DetectResult{}, packit.Fail
		}

		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{},
				Requires: []packit.BuildPlanRequirement{
					{
						Name: "gems",
						Metadata: BuildPlanMetadata{
							Launch: true,
						},
					},
					{
						Name: "bundler",
						Metadata: BuildPlanMetadata{
							Launch: true,
						},
					},
					{
						Name: "mri",
						Metadata: BuildPlanMetadata{
							Launch: true,
						},
					},
				},
			},
		}, nil
	}
}
