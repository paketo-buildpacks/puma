package main

import (
	"fmt"
	"path/filepath"

	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/scribe"
)

func Build(logger scribe.Logger) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		logger.Title("%s %s", context.BuildpackInfo.Name, context.BuildpackInfo.Version)
		logger.Process("Writing start command")

		command := fmt.Sprintf(`BUNDLE_GEMFILE="%s" bundle exec puma`, filepath.Join(context.WorkingDir, "Gemfile"))

		logger.Subprocess(command)

		return packit.BuildResult{
			Processes: []packit.Process{
				{
					Type:    "web",
					Command: command,
				},
			},
		}, nil
	}
}
