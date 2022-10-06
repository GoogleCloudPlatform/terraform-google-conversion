package e2etest

import (
	"os"
	"testing"
)

func TestComputeInstance(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
		return
	}

	tfFiles := []string{
		"full_compute_instance",
	}
	tmpDir := os.TempDir()
	data := initTestData()

	for _, name := range tfFiles {
		t.Run(name, func(t *testing.T) {
			roundtripTest(t, name, tmpDir, data)
		})
	}
}
