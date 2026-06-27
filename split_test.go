package stringutil_test

import (
	// Built-in/core modules.
	"testing"

	// External modules.
	"github.com/stretchr/testify/assert"

	// Internal modules.
	"github.com/cuberat-go/stringutil"
)

func TestSplitOnDelims(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		delims     []string
		mergeDelim bool
		expected   []string
	}{
		{
			name:       "Empty input string",
			input:      "",
			delims:     []string{","},
			mergeDelim: false,
			expected:   []string{""},
		},
		{
			name:       "No delimiters in input string",
			input:      "hello world",
			delims:     []string{","},
			mergeDelim: false,
			expected:   []string{"hello world"},
		},
		{
			name:       "Single delimiter",
			input:      "hello,world",
			delims:     []string{","},
			mergeDelim: false,
			expected:   []string{"hello", "world"},
		},
		{
			name:       "Multiple delimiters",
			input:      "hello,world;this,is;a,test",
			delims:     []string{",", ";"},
			mergeDelim: false,
			expected:   []string{"hello", "world", "this", "is", "a", "test"},
		},
		{
			name:       "Multiple consecutive delimiters",
			input:      "hello,,world;;;this,,is;a,test",
			delims:     []string{",", ";"},
			mergeDelim: false,
			expected:   []string{"hello", "", "world", "", "", "this", "", "is", "a", "test"},
		},
		{
			name:       "Merge consecutive delimiters",
			input:      "hello,,world;;;this,,is;a,test",
			delims:     []string{",", ";"},
			mergeDelim: true,
			expected:   []string{"hello", "world", "this", "is", "a", "test"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stringutil.SplitOnDelims(tt.input, tt.delims,
				tt.mergeDelim)
			assert.Equal(t, tt.expected, result)
		})
	}
}
