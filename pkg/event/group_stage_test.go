package event

import (
	"reflect"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestGroupStage_HandleGroupStageMatchResult(t *testing.T) {
	tests := map[string]struct {
		groupStage      *GroupStage
		matchResult     *GroupStageMatchResult
		groupStageAfter *GroupStage
	}{
		"updates all stats correctly after match completion": {
			groupStage: &GroupStage{
				TeamStats: []*TeamStats{
					{},
					{},
				},
			},
			matchResult: &GroupStageMatchResult{
				PointsPerTeam: map[uint32]uint32{
					0: 3,
					1: 1,
				},
				WonTeams: map[uint32]*TeamStats{
					0: {
						Kills:          6,
						Deaths:         5,
						Assists:        4,
						CombatScore:    3,
						ObjectiveScore: 2,
						SupportScore:   1,
					},
				},
				LostTeams: map[uint32]*TeamStats{
					1: {
						Kills:          1,
						Deaths:         2,
						Assists:        3,
						CombatScore:    4,
						ObjectiveScore: 5,
						SupportScore:   6,
					},
				},
			},
			groupStageAfter: &GroupStage{
				FinishedMatches: 1,
				TeamStats: []*TeamStats{
					{
						TotalMatches:   1,
						WonMatches:     1,
						LostMatches:    0,
						Points:         3,
						Kills:          6,
						Deaths:         5,
						Assists:        4,
						CombatScore:    3,
						ObjectiveScore: 2,
						SupportScore:   1,
					},
					{
						TotalMatches:   1,
						WonMatches:     0,
						LostMatches:    1,
						Points:         1,
						Kills:          1,
						Deaths:         2,
						Assists:        3,
						CombatScore:    4,
						ObjectiveScore: 5,
						SupportScore:   6,
					},
				},
			},
		},
		"updates only winners stats and points": {
			groupStage: &GroupStage{
				TeamStats: []*TeamStats{
					{},
					{},
				},
			},
			matchResult: &GroupStageMatchResult{
				PointsPerTeam: map[uint32]uint32{
					0: 3,
				},
				WonTeams: map[uint32]*TeamStats{
					0: {
						Kills:          6,
						Deaths:         5,
						Assists:        4,
						CombatScore:    3,
						ObjectiveScore: 2,
						SupportScore:   1,
					},
				},
				LostTeams: map[uint32]*TeamStats{},
			},
			groupStageAfter: &GroupStage{
				FinishedMatches: 1,
				TeamStats: []*TeamStats{
					{
						TotalMatches:   1,
						WonMatches:     1,
						LostMatches:    0,
						Points:         3,
						Kills:          6,
						Deaths:         5,
						Assists:        4,
						CombatScore:    3,
						ObjectiveScore: 2,
						SupportScore:   1,
					},
					{},
				},
			},
		},
		"updates only lost team stats and points": {
			groupStage: &GroupStage{
				TeamStats: []*TeamStats{
					{},
					{},
				},
			},
			matchResult: &GroupStageMatchResult{
				PointsPerTeam: map[uint32]uint32{
					1: 1,
				},
				WonTeams: map[uint32]*TeamStats{},
				LostTeams: map[uint32]*TeamStats{
					1: {
						Kills:          1,
						Deaths:         2,
						Assists:        3,
						CombatScore:    4,
						ObjectiveScore: 5,
						SupportScore:   6,
					},
				},
			},
			groupStageAfter: &GroupStage{
				FinishedMatches: 1,
				TeamStats: []*TeamStats{
					{},
					{
						TotalMatches:   1,
						WonMatches:     0,
						LostMatches:    1,
						Points:         1,
						Kills:          1,
						Deaths:         2,
						Assists:        3,
						CombatScore:    4,
						ObjectiveScore: 5,
						SupportScore:   6,
					},
				},
			},
		},
		"updates points per team and finished matches": {
			groupStage: &GroupStage{
				TeamStats: []*TeamStats{
					{},
					{},
				},
			},
			matchResult: &GroupStageMatchResult{
				PointsPerTeam: map[uint32]uint32{
					0: 3,
					1: 1,
				},
				WonTeams:  map[uint32]*TeamStats{},
				LostTeams: map[uint32]*TeamStats{},
			},
			groupStageAfter: &GroupStage{
				FinishedMatches: 1,
				TeamStats: []*TeamStats{
					{
						TotalMatches: 0,
						WonMatches:   0,
						LostMatches:  0,
						Points:       3,
					},
					{
						TotalMatches: 0,
						WonMatches:   0,
						LostMatches:  0,
						Points:       1,
					},
				},
			},
		},
		"updates only finished matches": {
			groupStage: &GroupStage{
				TeamStats: []*TeamStats{
					{},
					{},
				},
			},
			matchResult: &GroupStageMatchResult{
				PointsPerTeam: map[uint32]uint32{},
				WonTeams:      map[uint32]*TeamStats{},
				LostTeams:     map[uint32]*TeamStats{},
			},
			groupStageAfter: &GroupStage{
				FinishedMatches: 1,
				TeamStats: []*TeamStats{
					{},
					{},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.groupStage.FinishedMatches, uint32(0), "Should be 0 matches before the update")

			test.groupStage.HandleGroupStageMatchResult(test.matchResult)

			assert.Equal(t, test.groupStage.FinishedMatches, test.groupStageAfter.FinishedMatches, "Invalid FinishedMatches counter after handling match update")
			if !reflect.DeepEqual(test.groupStage.TeamStats, test.groupStageAfter.TeamStats) {
				t.Fatalf("returned %v; expected %v", test.groupStage, test.groupStageAfter)
			}
		})
	}
}

func TestGroupStage_GetTeamPlacements(t *testing.T) {
	tests := map[string]struct {
		instance GroupStage
		expected []*TeamPlacement
	}{
		"returns four team placements in ascending order": {
			instance: GroupStage{
				Teams: []*Team{
					{Name: "Team A"},
					{Name: "Team B"},
					{Name: "Team C"},
					{Name: "Team D"},
				},
				TeamStats: []*TeamStats{
					{TotalMatches: 3, WonMatches: 3, LostMatches: 0, Points: 9},
					{TotalMatches: 3, WonMatches: 0, LostMatches: 3, Points: 0},
					{TotalMatches: 3, WonMatches: 1, LostMatches: 2, Points: 3},
					{TotalMatches: 3, WonMatches: 2, LostMatches: 1, Points: 6},
				},
			},
			expected: []*TeamPlacement{
				{Name: "Team A", Matches: 3, Wins: 3, Loses: 0, Points: 9},
				{Name: "Team D", Matches: 3, Wins: 2, Loses: 1, Points: 6},
				{Name: "Team C", Matches: 3, Wins: 1, Loses: 2, Points: 3},
				{Name: "Team B", Matches: 3, Wins: 0, Loses: 3, Points: 0},
			},
		},
		"returns team placements in ascending order with 2 tied teams": {
			instance: GroupStage{
				Teams: []*Team{
					{Name: "Team A"},
					{Name: "Team B"},
					{Name: "Team C"},
					{Name: "Team D"},
				},
				TeamStats: []*TeamStats{
					{TotalMatches: 3, WonMatches: 3, LostMatches: 0, Points: 9},
					{TotalMatches: 3, WonMatches: 0, LostMatches: 3, Points: 0},
					{TotalMatches: 3, WonMatches: 2, LostMatches: 1, Points: 6},
					{TotalMatches: 3, WonMatches: 2, LostMatches: 1, Points: 6},
				},
			},
			expected: []*TeamPlacement{
				{Name: "Team A", Matches: 3, Wins: 3, Loses: 0, Points: 9},
				{Name: "Team C", Matches: 3, Wins: 2, Loses: 1, Points: 6},
				{Name: "Team D", Matches: 3, Wins: 2, Loses: 1, Points: 6},
				{Name: "Team B", Matches: 3, Wins: 0, Loses: 3, Points: 0},
			},
		},
		"returns team placements in ascending order with all tied teams": {
			instance: GroupStage{
				Teams: []*Team{
					{Name: "Team A"},
					{Name: "Team B"},
					{Name: "Team C"},
				},
				TeamStats: []*TeamStats{
					{TotalMatches: 3, WonMatches: 1, LostMatches: 2, Points: 3},
					{TotalMatches: 3, WonMatches: 1, LostMatches: 2, Points: 3},
					{TotalMatches: 3, WonMatches: 1, LostMatches: 2, Points: 3},
				},
			},
			expected: []*TeamPlacement{
				{Name: "Team A", Matches: 3, Wins: 1, Loses: 2, Points: 3},
				{Name: "Team B", Matches: 3, Wins: 1, Loses: 2, Points: 3},
				{Name: "Team C", Matches: 3, Wins: 1, Loses: 2, Points: 3},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := test.instance.GetTeamPlacements()
			expected := test.expected

			equal := slices.EqualFunc(result, expected, func(first *TeamPlacement, other *TeamPlacement) bool {
				if first == nil || other == nil {
					return false
				}

				return reflect.DeepEqual(first, other)
			})
			if !equal {
				t.Fatalf("returned %+v; expected %+v", result, expected)
			}
		})
	}
}
