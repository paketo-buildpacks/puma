package puma

import (
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func Build(logger scribe.Emitter) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		logger.Title("%s %s", context.BuildpackInfo.Name, context.BuildpackInfo.Version)

		args := `bundle exec puma --bind tcp://0.0.0.0:${PORT:-9292}`
		processes := []packit.Process{
			{
				Type:    "web",
				Command: "bash",
				Args:    []string{"-c", args},
				Default: true,
				Direct:  true,
			},
		}
		logger.LaunchProcesses(processes)

		return packit.BuildResult{
			Launch: packit.LaunchMetadata{
				Processes: processes,
			},
		}, nil
	}
}
