package main_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/packit"
	main "github.com/paketo-community/puma"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testBuild(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		layersDir  string
		workingDir string
		cnbDir     string
		//		buffer     *bytes.Buffer

		build packit.BuildFunc
	)

	it.Before(func() {
		var err error
		layersDir, err = ioutil.TempDir("", "layers")
		Expect(err).NotTo(HaveOccurred())

		cnbDir, err = ioutil.TempDir("", "cnb")
		Expect(err).NotTo(HaveOccurred())

		workingDir, err = ioutil.TempDir("", "working-dir")
		Expect(err).NotTo(HaveOccurred())

		// buffer = bytes.NewBuffer(nil)

		build = main.Build()
	})

	it.After(func() {
		Expect(os.RemoveAll(layersDir)).To(Succeed())
		Expect(os.RemoveAll(cnbDir)).To(Succeed())
		Expect(os.RemoveAll(workingDir)).To(Succeed())
	})

	it("returns a result that provides a puma start command", func() {
		result, err := build(packit.BuildContext{
			WorkingDir: workingDir,
			CNBPath:    cnbDir,
			Stack:      "some-stack",
			BuildpackInfo: packit.BuildpackInfo{
				Name:    "Some Buildpack",
				Version: "some-version",
			},
			Plan: packit.BuildpackPlan{
				Entries: []packit.BuildpackPlanEntry{},
			},
			Layers: packit.Layers{Path: layersDir},
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(result).To(Equal(packit.BuildResult{
			Plan: packit.BuildpackPlan{
				Entries: nil,
			},
			Layers: nil,
			Processes: []packit.Process{
				{
					Type:    "web",
					Command: fmt.Sprintf(`BUNDLE_GEMFILE="%s" bundle exec puma`, filepath.Join(workingDir, "Gemfile")),
				},
			},
		}))

		// Expect(buffer.String()).To(ContainSubstring("Some Buildpack some-version"))
		// Expect(buffer.String()).To(ContainSubstring("Executing build process"))
		// Expect(buffer.String()).To(ContainSubstring("Configuring environment"))
	})

	// context("failure cases", func() {
	// 	context("when the layers directory cannot be written to", func() {
	// 		it.Before(func() {
	// 			Expect(os.Chmod(layersDir, 0000)).To(Succeed())
	// 		})
	//
	// 		it.After(func() {
	// 			Expect(os.Chmod(layersDir, os.ModePerm)).To(Succeed())
	// 		})
	//
	// 		it("returns an error", func() {
	// 			_, err := build(packit.BuildContext{
	// 				WorkingDir: workingDir,
	// 				CNBPath:    cnbDir,
	// 				Stack:      "some-stack",
	// 				BuildpackInfo: packit.BuildpackInfo{
	// 					Name:    "Some Buildpack",
	// 					Version: "some-version",
	// 				},
	// 				Plan: packit.BuildpackPlan{
	// 					Entries: []packit.BuildpackPlanEntry{
	// 						{
	// 							Name: "gems",
	// 						},
	// 					},
	// 				},
	// 				Layers: packit.Layers{Path: layersDir},
	// 			})
	// 			Expect(err).To(MatchError(ContainSubstring("permission denied")))
	// 		})
	// 	})
	//
	// 	context("when the layer directory cannot be removed", func() {
	// 		var layerDir string
	// 		it.Before(func() {
	// 			layerDir = filepath.Join(layersDir, bundle.LayerNameGems)
	// 			Expect(os.MkdirAll(filepath.Join(layerDir, "baller"), os.ModePerm)).To(Succeed())
	// 			Expect(os.Chmod(layerDir, 0000)).To(Succeed())
	// 		})
	//
	// 		it.After(func() {
	// 			Expect(os.Chmod(layerDir, os.ModePerm)).To(Succeed())
	// 			Expect(os.RemoveAll(layerDir)).To(Succeed())
	// 		})
	//
	// 		it("returns an error", func() {
	// 			_, err := build(packit.BuildContext{
	// 				WorkingDir: workingDir,
	// 				CNBPath:    cnbDir,
	// 				Stack:      "some-stack",
	// 				BuildpackInfo: packit.BuildpackInfo{
	// 					Name:    "Some Buildpack",
	// 					Version: "some-version",
	// 				},
	// 				Plan: packit.BuildpackPlan{
	// 					Entries: []packit.BuildpackPlanEntry{
	// 						{
	// 							Name: "gems",
	// 						},
	// 					},
	// 				},
	// 				Layers: packit.Layers{Path: layersDir},
	// 			})
	// 			Expect(err).To(MatchError(ContainSubstring("permission denied")))
	// 		})
	// 	})
	//
	// 	context("when install process returns an error", func() {
	// 		it.Before(func() {
	// 			installProcess.ExecuteCall.Returns.Error = errors.New("some-error")
	// 		})
	//
	// 		it("returns an error", func() {
	// 			_, err := build(packit.BuildContext{
	// 				WorkingDir: workingDir,
	// 				CNBPath:    cnbDir,
	// 				Stack:      "some-stack",
	// 				BuildpackInfo: packit.BuildpackInfo{
	// 					Name:    "Some Buildpack",
	// 					Version: "some-version",
	// 				},
	// 				Plan: packit.BuildpackPlan{
	// 					Entries: []packit.BuildpackPlanEntry{
	// 						{
	// 							Name: "gems",
	// 						},
	// 					},
	// 				},
	// 				Layers: packit.Layers{Path: layersDir},
	// 			})
	// 			Expect(err).To(MatchError(ContainSubstring("some-error")))
	// 		})
	// 	})
	// })
}
