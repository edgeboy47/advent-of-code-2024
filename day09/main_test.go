package main

import (
	"slices"
	"strings"
	"testing"
)

func TestDiskMapToFileBlocks(t *testing.T) {
	t.Run("Convert disk map to file blocks", func(t *testing.T) {
		vals := map[string]string{
			"12345":               "0..111....22222",
			"2333133121414131402": "00...111...2...333.44.5555.6666.777.888899",
		}

		for id, val := range vals {
			expected := strings.Split(val, "")
			actual := diskmapToFileBlocks(id)

			if !slices.Equal(expected, actual) {
				t.Errorf("Expected %s, Got %s", expected, actual)
			}
		}
	})

	t.Run("Compact file blocks", func(t *testing.T) {
		vals := []struct {
			input, output string
		}{
			{"0..111....22222", "022111222......"},
			{"00...111...2...333.44.5555.6666.777.888899", "0099811188827773336446555566.............."},
		}

		for _, val := range vals {
			expected := strings.Split(val.output, "")
			actual := compactFileBlock(strings.Split(val.input, ""))

			if !slices.Equal(expected, actual) {
				t.Errorf("Error with compactFileBlock. Expected %s, Got %s", expected, actual)
			}
		}
	})

	t.Run("Calculate Checksum", func(t *testing.T) {
	  vals := map[string] int {
	    "0099811188827773336446555566..............": 1928,
	  }

	for id, val := range vals {
		expected := val
		actual := calculateChecksum(strings.Split(id, ""))

		if expected != actual {
	      t.Errorf("Error with compactFileBlock. Expected %d, Got %d", expected, actual)
		}
	}
	})
}
