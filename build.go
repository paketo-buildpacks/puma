package main

import (
	"fmt"
	"path/filepath"

	"github.com/paketo-buildpacks/packit"
)

func Build() packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {

		return packit.BuildResult{
			Processes: []packit.Process{
				{
					Type:    "web",
					Command: fmt.Sprintf(`BUNDLE_GEMFILE="%s" bundle exec puma`, filepath.Join(context.WorkingDir, "Gemfile")),
				},
			},
		}, nil
	}
}
