package pkg

import (
	"testing"
)

func TestGroupStage_HasFinishedAllMatches(t *testing.T) {
	tests := map[string]struct {
		instance GroupStage
		expected bool
	}{
		"returns true for the finished group sage": {
			instance: GroupStage{
				TotalMatches:    3,
				FinishedMatches: 3,
			},
			expected: true,
		},
		"returns false for a new group stage": {
			instance: GroupStage{
				TotalMatches:    0,
				FinishedMatches: 0,
			},
			expected: false,
		},
		"returns false for the ongoing group stage": {
			instance: GroupStage{
				TotalMatches:    3,
				FinishedMatches: 1,
			},
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := test.instance.HasFinishedAllMatches()
			expected := test.expected

			if result != expected {
				t.Fatalf("returned %v; expected %v", result, expected)
			}
		})
	}
}
