package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamStatsTiebreakResolver_IsDeciderMatchRequired(t *testing.T) {
	tests := map[string]struct {
		instance   *TeamStatsTiebreakResolver
		groupStage *GroupStage
		expected   bool
	}{
		"always returns false": {
			instance:   &TeamStatsTiebreakResolver{},
			groupStage: &GroupStage{},
			expected:   false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := test.instance.IsDeciderMatchRequired(test.groupStage)
			assert.Equal(t, result, test.expected, "IsDeciderMatchRequired returned wrong result")
		})
	}
}

func TestTeamStatsTiebreakResolver_DetermineRanking(t *testing.T) {
	tests := map[string]struct {
		instance   *TeamStatsTiebreakResolver
		groupStage *GroupStage
		expected   []uint32{}
	}{
		"decides by win and loss rounds": {
			instance:   &TeamStatsTiebreakResolver{},
			groupStage: &GroupStage{},
			expected:   []uint32{},
		},
		// decides by win/loss & combat/support/objective scores
		// decides by win/loss, combat/support/objective scores and KDA between teams
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := test.instance.IsDeciderMatchRequired(test.groupStage)
		})
	}
}
