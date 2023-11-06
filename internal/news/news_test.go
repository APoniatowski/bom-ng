package news

import (
	"testing"
)

func TestCheckText(t *testing.T) {
	tests := []struct {
		text     string
		expected string // Assuming you have some way to capture the output to compare
	}{
		{
			text:     "Russia invades Belarus.",
			expected: "news message: russia invades belarus",
		},
		// ...other test cases...
	}

	for _, test := range tests {

		tc := TextChecker{}
		// Assuming you modify checkText to return a string for testing
		result := tc.CheckText(test.text)
		if result != test.expected {
			t.Errorf("checkText(%q) = %q; want %q", test.text, result, test.expected)
		}
	}
}
