package analysis_test

import (
	"testing"

	"github.com/just-hms/retguard/pkg/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestRetGuard(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analysis.RetGuard)
}
