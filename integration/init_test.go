package integration_test

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/cloudfoundry/dagger"
	"github.com/paketo-buildpacks/occam"
	"github.com/paketo-buildpacks/packit/pexec"
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
	Expect(err).NotTo(HaveOccurred())

	pumaBuildpack, err = dagger.PackageBuildpack(root)
	Expect(err).NotTo(HaveOccurred())

	// HACK: we need to fix dagger and the package.sh scripts so that this isn't required
	pumaBuildpack = fmt.Sprintf("%s.tgz", pumaBuildpack)

	mriBuildpack, err = dagger.GetLatestCommunityBuildpack("paketo-community", "mri")
	Expect(err).NotTo(HaveOccurred())

	bundlerBuildpack, err = dagger.GetLatestCommunityBuildpack("paketo-community", "bundler")
	Expect(err).NotTo(HaveOccurred())

	bundleInstallBuildpack, err = dagger.GetLatestCommunityBuildpack("paketo-community", "bundle-install")
	Expect(err).NotTo(HaveOccurred())

	SetDefaultEventuallyTimeout(10 * time.Second)

	suite := spec.New("Integration", spec.Parallel(), spec.Report(report.Terminal{}))
	suite("SimpleApp", testSimpleApp)

	defer AfterSuite(t)
	suite.Run(t)
}

func AfterSuite(t *testing.T) {
	var Expect = NewWithT(t).Expect

	Expect(dagger.DeleteBuildpack(pumaBuildpack)).To(Succeed())
	Expect(dagger.DeleteBuildpack(mriBuildpack)).To(Succeed())
	Expect(dagger.DeleteBuildpack(bundlerBuildpack)).To(Succeed())
	Expect(dagger.DeleteBuildpack(bundleInstallBuildpack)).To(Succeed())
}

func ContainerLogs(id string) func() string {
	docker := occam.NewDocker()

	return func() string {
		logs, _ := docker.Container.Logs.Execute(id)
		return logs.String()
	}
}

func GetGitVersion() (string, error) {
	gitExec := pexec.NewExecutable("git")
	revListOut := bytes.NewBuffer(nil)

	err := gitExec.Execute(pexec.Execution{
		Args:   []string{"rev-list", "--tags", "--max-count=1"},
		Stdout: revListOut,
	})
	if err != nil {
		return "", err
	}

	stdout := bytes.NewBuffer(nil)
	err = gitExec.Execute(pexec.Execution{
		Args:   []string{"describe", "--tags", strings.TrimSpace(revListOut.String())},
		Stdout: stdout,
	})
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(strings.TrimPrefix(stdout.String(), "v")), nil
}
