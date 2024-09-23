package pkg

import (
	"sort"
)

type GroupStage struct {
	Name            string
	Teams           []*Team
	TeamStats       []*TeamStats
	TotalMatches    uint32
	FinishedMatches uint32
}

// HasFinishedAllMatches checks that N matches were finished before
// making a decision what teams must be advanced further.
func (gs *GroupStage) HasFinishedAllMatches() bool {
	return gs.TotalMatches != 0 && gs.TotalMatches == gs.FinishedMatches
}

// HandleGroupStageMatchResult updates team result based on the given result
func (gs *GroupStage) HandleGroupStageMatchResult(result *GroupStageMatchResult) {
	for teamIndex, teamStats := range result.WonTeams {
		if int(teamIndex) > len(gs.TeamStats) {
			continue
		}

		gs.TeamStats[teamIndex].TotalMatches += 1
		gs.TeamStats[teamIndex].WonMatches += 1

		gs.updateTeamStats(teamIndex, teamStats)
	}

	for teamIndex, teamStats := range result.LostTeams {
		if int(teamIndex) > len(gs.TeamStats) {
			continue
		}

		gs.TeamStats[teamIndex].TotalMatches += 1
		gs.TeamStats[teamIndex].LostMatches += 1

		gs.updateTeamStats(teamIndex, teamStats)
	}

	// Can convenient to do it like this, so it would be possible to
	// apply technical defeats and assign points separately
	for teamIndex, points := range result.PointsPerTeam {
		if int(teamIndex) > len(gs.TeamStats) {
			continue
		}

		gs.TeamStats[teamIndex].Points += points
	}

	gs.FinishedMatches += 1
}

func (gs *GroupStage) updateTeamStats(teamIndex uint32, stats *TeamStats) {
	if stats == nil {
		return
	}

	gs.TeamStats[teamIndex].Kills += stats.Kills
	gs.TeamStats[teamIndex].Deaths += stats.Deaths
	gs.TeamStats[teamIndex].Assists += stats.Assists
	gs.TeamStats[teamIndex].CombatScore += stats.CombatScore
	gs.TeamStats[teamIndex].ObjectiveScore += stats.ObjectiveScore
	gs.TeamStats[teamIndex].SupportScore += stats.SupportScore
}

// GetTeamPlacements returns current positions of teams within the
// group based on the earned points & win-loses
func (gs *GroupStage) GetTeamPlacements() []*TeamPlacement {
	teamPlacement := make([]*TeamPlacement, 0)
	for index, team := range gs.Teams {
		teamStats := gs.TeamStats[index]

		teamPlacement = append(teamPlacement, &TeamPlacement{
			Name:    team.Name,
			Matches: teamStats.TotalMatches,
			Wins:    teamStats.WonMatches,
			Loses:   teamStats.LostMatches,
			Points:  teamStats.Points,
		})
	}

	sort.Slice(teamPlacement, func(i, j int) bool {
		return teamPlacement[i].Points > teamPlacement[j].Points &&
			teamPlacement[i].Wins > teamPlacement[j].Wins &&
			teamPlacement[i].Loses < teamPlacement[j].Loses
	})

	return teamPlacement
}
