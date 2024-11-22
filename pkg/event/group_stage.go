package event

import (
	"sort"
)

type GroupStage struct {
	Name            string
	Teams           map[string]*Team
	TeamStats       map[string]*TeamStats
	TeamPlacements  []*TeamPlacement
	TotalMatches    uint32
	FinishedMatches uint32
}

// HasFinishedAllMatches checks that N matches were finished before
// making a decision what teams must be advanced further.
func (gs *GroupStage) HasFinishedAllMatches() bool {
	return gs.TotalMatches != 0 && gs.TotalMatches == gs.FinishedMatches
}

// HandleGroupStageMatchResult updates team positions and stats based on the given match result
func (gs *GroupStage) HandleGroupStageMatchResult(matchResult *GroupStageMatchResult) {
	for _, entry := range matchResult.TeamPerformances {
		if entry == nil {
			// FIXME: Add logging
			continue
		}

		if team, exists := gs.Teams[entry.TeamId]; exists {
			gs.updatePlayerStats(team, entry)
		}

		if teamStats, exists := gs.TeamStats[entry.TeamId]; exists {
			gs.updateTeamStats(teamStats, entry)
		}
	}

	gs.FinishedMatches += 1
	gs.updateTeamPlacements()
}

// updatePlayerStats refreshes general player stats the based on the reported match results
func (gs *GroupStage) updatePlayerStats(team *Team, entry *TeamPerformanceEntry) {
	for _, playerMatchStats := range entry.PlayerMatchStats {
		playerStats, exists := team.PlayerStats[playerMatchStats.PlayerId]
		if !exists {
			// FIXME: Add logging
			return
		}

		playerStats.Kills += playerMatchStats.Kills
		playerStats.Deaths += playerMatchStats.Deaths
		playerStats.Assists += playerMatchStats.Assists
		playerStats.CombatScore += playerMatchStats.CombatScore
		playerStats.ObjectiveScore += playerMatchStats.ObjectiveScore
		playerStats.SupportScore += playerMatchStats.SupportScore
	}
}

// updateTeamStats refreshes team stats relevant only within the group stage
func (gs *GroupStage) updateTeamStats(teamStats *TeamStats, entry *TeamPerformanceEntry) {
	teamStats.TotalMatches += 1
	teamStats.Points += entry.Points

	if entry.IsWinner {
		teamStats.WonMatches += 1
	} else {
		teamStats.LostMatches += 1
	}
}

// updateTeamPlacements refrehes current positions of teams within the
// group based on the earned points & win-loses
func (gs *GroupStage) updateTeamPlacements() {
	teamPlacement := make([]*TeamPlacement, 0)
	for _, team := range gs.Teams {
		if teamStats, exists := gs.TeamStats[team.Id]; exists {
			teamPlacement = append(teamPlacement, &TeamPlacement{
				Name:       team.Name,
				BrandIndex: team.BrandIndex,
				Matches:    teamStats.TotalMatches,
				Wins:       teamStats.WonMatches,
				Loses:      teamStats.LostMatches,
				Points:     teamStats.Points,
			})
		}
	}

	sort.Slice(teamPlacement, func(i, j int) bool {
		return teamPlacement[i].Points > teamPlacement[j].Points &&
			teamPlacement[i].Wins > teamPlacement[j].Wins &&
			teamPlacement[i].Loses < teamPlacement[j].Loses
	})

	gs.TeamPlacements = teamPlacement
}
