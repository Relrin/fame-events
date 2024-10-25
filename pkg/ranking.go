package pkg

import "sort"

// GetTeamsRankingByPerformance returns a sorted list of indices for teams based on the following rules:
// - Order by points (as normally we do)
// - Order by win/loss matches for teams. If they are equal compare by the next rule.
// - Order by combat/support/objective scores. If they are equal compare by the next rule.
// - Order by KDA between teams
func GetTeamsRankingByPerformance(stats []*TeamStats) []uint32 {
	aggregatedTeamData := make([]*AggregatedTeamData, 0)

	for index, teamStats := range stats {
		aggregatedTeamData = append(aggregatedTeamData, &AggregatedTeamData{
			TeamIndex: uint32(index),
			TeamStats: teamStats,
		})
	}

	sort.Slice(aggregatedTeamData, func(i, j int) bool {
		return aggregatedTeamData[i].TeamStats.Points > aggregatedTeamData[j].TeamStats.Points &&
			// Order by win/loss. Those who won more matches have an advantage
			aggregatedTeamData[i].TeamStats.WonMatches > aggregatedTeamData[j].TeamStats.WonMatches &&
			aggregatedTeamData[i].TeamStats.LostMatches < aggregatedTeamData[j].TeamStats.LostMatches &&

			// Then by Combat/Support/Objective scores. We prioritize playing into objective as the main
			// rule and then by support & combats characteristics
			aggregatedTeamData[i].TeamStats.ObjectiveScore > aggregatedTeamData[j].TeamStats.ObjectiveScore &&
			aggregatedTeamData[i].TeamStats.SupportScore > aggregatedTeamData[j].TeamStats.SupportScore &&
			aggregatedTeamData[i].TeamStats.CombatScore > aggregatedTeamData[j].TeamStats.CombatScore &&

			// And then by the overall KDA performance between the teams. Those teams which do
			// more kills and assists, fewer deaths have an advantage.
			aggregatedTeamData[i].TeamStats.Kills > aggregatedTeamData[j].TeamStats.Kills &&
			aggregatedTeamData[i].TeamStats.Assists > aggregatedTeamData[j].TeamStats.Assists &&
			aggregatedTeamData[i].TeamStats.Deaths < aggregatedTeamData[j].TeamStats.Deaths
	})

	ranking := make([]uint32, 0)
	for _, teamEntry := range aggregatedTeamData {
		ranking = append(ranking, teamEntry.TeamIndex)
	}

	return ranking
}
