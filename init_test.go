package puma_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnitBundleInstall(t *testing.T) {
	suite := spec.New("puma", spec.Report(report.Terminal{}), spec.Parallel())
	suite("Build", testBuild)
	suite("Detect", testDetect)
	suite("GemfileParser", testGemfileParser)
	suite.Run(t)
}
