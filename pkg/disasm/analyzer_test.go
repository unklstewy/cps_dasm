package disasm

import (
	"testing"

	cpstesting "github.com/unklstewy/cps_dasm/pkg/disasm/testing"
)

func TestIdentifyCPSType(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		expected CPSType
	}{
		{
			name:     "Baofeng CPS",
			filename: "CPS32UV.exe",
			expected: CPSTypeBaofeng,
		},
		{
			name:     "TYT CPS",
			filename: "TYT-MD-380-180317.exe",
			expected: CPSTypeTYT,
		},
		{
			name:     "Unknown CPS",
			filename: "something_else.exe",
			expected: CPSTypeUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IdentifyCPSType(tt.filename)

			if result != tt.expected {
				t.Errorf("IdentifyCPSType(%s) = %v, want %v", tt.filename, result, tt.expected)
			}
		})
	}
}

func TestParseCPSStrings(t *testing.T) {
	// This test is a placeholder - we need actual binary samples
	t.Skip("Requires real CPS binary samples")

	analyzer := NewCPSAnalyzer()

	// Sample binary data containing strings
	sampleData := cpstesting.LoadTestFile(t, "samples/cps_sample.bin")

	results := analyzer.ExtractStrings(sampleData)

	if len(results) == 0 {
		t.Error("Expected to extract some strings, got none")
	}

	// Check if results contain expected strings
	expectedStrings := []string{"Channel", "Frequency", "Timeout"}
	for _, str := range expectedStrings {
		found := false
		for _, result := range results {
			if result.Text == str {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Expected to find string %q in results, but didn't", str)
		}
	}
}
