package testing

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// GetTestDataPath returns the absolute path to the testdata directory
func GetTestDataPath() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), "..", "testdata")
}

// LoadTestFile loads a file from the testdata directory
func LoadTestFile(t *testing.T, relativePath string) []byte {
	t.Helper()

	path := filepath.Join(GetTestDataPath(), relativePath)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to load test file %s: %v", path, err)
	}
	return data
}

// CompareWithGolden compares output with a golden file
func CompareWithGolden(t *testing.T, got []byte, goldenPath string, update bool) {
	t.Helper()

	path := filepath.Join(GetTestDataPath(), "expected", goldenPath)

	if update {
		if err := os.WriteFile(path, got, 0644); err != nil {
			t.Fatalf("Failed to update golden file %s: %v", path, err)
		}
		return
	}

	expected, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read golden file %s: %v", path, err)
	}

	if string(got) != string(expected) {
		t.Errorf("Output doesn't match golden file %s", goldenPath)
	}
}
