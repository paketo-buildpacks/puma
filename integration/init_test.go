package integration_test

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/cloudfoundry/dagger"
	"github.com/paketo-buildpacks/occam"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
)

var (
	pumaBuildpack          string
	mriBuildpack           string
	bundlerBuildpack       string
	bundleInstallBuildpack string
)

func TestIntegration(t *testing.T) {
	Expect := NewWithT(t).Expect

	root, err := filepath.Abs("./..")
	Expect(err).ToNot(HaveOccurred())

	pumaBuildpack, err = dagger.PackageBuildpack(root)
	Expect(err).ToNot(HaveOccurred())

	// HACK: we need to fix dagger and the package.sh scripts so that this isn't required
	pumaBuildpack = fmt.Sprintf("%s.tgz", pumaBuildpack)

	mriBuildpack, err = dagger.GetLatestCommunityBuildpack("paketo-community", "mri")
	Expect(err).ToNot(HaveOccurred())

	bundlerBuildpack, err = dagger.GetLatestCommunityBuildpack("paketo-community", "bundler")
	Expect(err).ToNot(HaveOccurred())

	bundleInstallBuildpack, err = dagger.GetLatestCommunityBuildpack("paketo-community", "bundle-install")
	Expect(err).ToNot(HaveOccurred())

	defer func() {
		dagger.DeleteBuildpack(pumaBuildpack)
		dagger.DeleteBuildpack(mriBuildpack)
		dagger.DeleteBuildpack(bundlerBuildpack)
		dagger.DeleteBuildpack(bundleInstallBuildpack)
	}()

	SetDefaultEventuallyTimeout(10 * time.Second)

	suite := spec.New("Integration", spec.Parallel(), spec.Report(report.Terminal{}))
	suite("SimpleApp", testSimpleApp)
	suite.Run(t)
}

func ContainerLogs(id string) func() string {
	docker := occam.NewDocker()

	return func() string {
		logs, _ := docker.Container.Logs.Execute(id)
		return logs.String()
	}
}
