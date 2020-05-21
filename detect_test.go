package main_test

import (
	"os"
	"path/filepath"
	"testing"

	"io/ioutil"

	"github.com/paketo-buildpacks/packit"
	main "github.com/paketo-community/puma"
	"github.com/paketo-community/puma/fakes"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testDetect(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		workingDir    string
		gemfileParser *fakes.Parser
		detect        packit.DetectFunc
	)

	it.Before(func() {
		var err error
		workingDir, err = ioutil.TempDir("", "working-dir")
		Expect(err).NotTo(HaveOccurred())

		const GEMFILE_CONTENTS = `source 'https://rubygems.org'
ruby '~> 2.0'

gem 'puma'`
		err = ioutil.WriteFile(filepath.Join(workingDir, "Gemfile"), []byte(GEMFILE_CONTENTS), 0644)
		Expect(err).NotTo(HaveOccurred())

		gemfileParser = &fakes.Parser{}

		detect = main.Detect(gemfileParser)
	})

	it.After(func() {
		Expect(os.RemoveAll(workingDir)).To(Succeed())
	})

	it.Focus("returns a plan that provides gems", func() {
		result, err := detect(packit.DetectContext{
			WorkingDir: workingDir,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(result.Plan).To(Equal(packit.BuildPlan{
			Provides: []packit.BuildPlanProvision{},
			Requires: []packit.BuildPlanRequirement{
				{
					Name: "gems",
					Metadata: main.BuildPlanMetadata{
						Launch: true,
					},
				},
				{
					Name: "mri",
					Metadata: main.BuildPlanMetadata{
						Launch: true,
					},
				},
			},
		}))
	})

	context("when the Gemfile does not list puma", func() {
		it.Before(func() {
			gemfileParser.ParseCall.Returns.HasPuma = false
		})

		it("detect should fail with error", func() {
			_, err := detect(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).To(MatchError(packit.Fail))
		})
	})
	//
	// context("when the buildpack.yml parser fails", func() {
	// 	it.Before(func() {
	// 		gemfileParser.ParseVersionCall.Returns.Err = errors.New("some-error")
	// 	})
	//
	// 	it("returns an error", func() {
	// 		_, err := detect(packit.DetectContext{
	// 			WorkingDir: workingDir,
	// 		})
	// 		Expect(err).To(MatchError("failed to parse Gemfile: some-error"))
	// 	})
	// })
}
