package tournament

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
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
		expected   []uint32
	}{
		"decides by win-loss rounds": {
			instance: &TeamStatsTiebreakResolver{},
			groupStage: &GroupStage{
				TeamStats: []*TeamStats{
					{
						TotalMatches: 4,
						WonMatches:   4,
						LostMatches:  0,
						Points:       12,
					},
					{
						TotalMatches: 3,
						WonMatches:   1,
						LostMatches:  2,
						Points:       6,
					},
					{
						TotalMatches: 3,
						WonMatches:   2,
						LostMatches:  1,
						Points:       6,
					},
					{
						TotalMatches: 3,
						WonMatches:   2,
						LostMatches:  1,
						Points:       9,
					},
				},
			},
			expected: []uint32{0, 3, 2, 1},
		},
		"decides by win-loss & combat/support/objective scores": {
			instance: &TeamStatsTiebreakResolver{},
			groupStage: &GroupStage{
				TeamStats: []*TeamStats{
					{
						TotalMatches:   3,
						WonMatches:     1,
						LostMatches:    2,
						Points:         6,
						CombatScore:    750,
						ObjectiveScore: 750,
						SupportScore:   750,
					},
					{
						TotalMatches:   3,
						WonMatches:     1,
						LostMatches:    2,
						Points:         6,
						CombatScore:    500,
						ObjectiveScore: 500,
						SupportScore:   500,
					},
					{
						TotalMatches:   3,
						WonMatches:     3,
						LostMatches:    0,
						Points:         9,
						CombatScore:    1000,
						ObjectiveScore: 1000,
						SupportScore:   1000,
					},
				},
			},
			expected: []uint32{2, 0, 1},
		},
		"decides by win-loss, combat/support/objective scores and KDA between teams": {
			instance: &TeamStatsTiebreakResolver{},
			groupStage: &GroupStage{
				TeamStats: []*TeamStats{
					{
						TotalMatches:   3,
						WonMatches:     1,
						LostMatches:    2,
						Points:         6,
						Kills:          5,
						Deaths:         3,
						Assists:        1,
						CombatScore:    750,
						ObjectiveScore: 750,
						SupportScore:   750,
					},
					{
						TotalMatches:   3,
						WonMatches:     1,
						LostMatches:    2,
						Points:         6,
						Kills:          0,
						Deaths:         0,
						Assists:        0,
						CombatScore:    500,
						ObjectiveScore: 500,
						SupportScore:   500,
					},
					{
						TotalMatches:   3,
						WonMatches:     3,
						LostMatches:    0,
						Points:         9,
						Kills:          0,
						Deaths:         0,
						Assists:        0,
						CombatScore:    1000,
						ObjectiveScore: 1000,
						SupportScore:   1000,
					},
					{
						TotalMatches:   3,
						WonMatches:     1,
						LostMatches:    2,
						Points:         6,
						Kills:          5,
						Deaths:         2,
						Assists:        1,
						CombatScore:    750,
						ObjectiveScore: 750,
						SupportScore:   750,
					},
				},
			},
			expected: []uint32{2, 3, 0, 1},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := test.instance.DetermineRanking(test.groupStage)

			if reflect.DeepEqual(result, test.expected) {
				t.Fatalf("returned %+v; expected %+v", result, test.expected)
			}
		})
	}
}
