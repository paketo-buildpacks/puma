package puma

import (
	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/scribe"
)

func Build(logger scribe.Logger) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		logger.Title("%s %s", context.BuildpackInfo.Name, context.BuildpackInfo.Version)

		command := "bundle exec puma --bind tcp://0.0.0.0:${PORT:-9292}"

		logger.Process("Assigning launch processes")
		logger.Subprocess("web: %s", command)
		logger.Break()

		return packit.BuildResult{
			Launch: packit.LaunchMetadata{
				Processes: []packit.Process{
					{
						Type:    "web",
						Command: command,
						Default: true,
					},
				},
			},
		}, nil
	}
}
