package gokeyword

import (
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	tests := []struct {
		dir     string
		details string
	}{
		{
			dir:     "go_no_details",
			details: "",
		},
		{
			dir:     "go_with_details",
			details: "details details details",
		},
		{
			dir:     "no_go",
			details: "",
		},
	}

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	for _, tc := range tests {
		tt := tc

		t.Run(tt.dir, func(t *testing.T) {
			td := filepath.Join(wd, "testdata", tt.dir)
			a := New()
			if tt.details != "" {
				_ = a.Flags.Set("details", tt.details)
			}
			analysistest.Run(t, td, a)
		})
	}
}
