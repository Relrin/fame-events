package event

import "sort"

type PlayOffTeamOptimizer interface {
	PrepareTeams(config PlayOffTeamOptimizerConfig, teams []*Team, groupStageTeamStats []*TeamStats) []*PlayOffMatch
}

type PlayOffTeamOptimizerConfig struct {
	TotalMatches  int
	TeamsPerMatch int
}

type PlayOffMatch struct {
	Teams []*Team
}

// NoopPlayOffOptimizer is a default play off optimizer which doesn't apply any changes
// to the input and passed the data back as-is to the caller
type NoopPlayOffOptimizer struct{}

// PrepareTeams returns teams as-is with no changes to their ranking
func (n *NoopPlayOffOptimizer) PrepareTeams(config PlayOffTeamOptimizerConfig, teams []*Team, groupStageStats []*TeamStats) []*PlayOffMatch {
	matches := make([]*PlayOffMatch, 0)

	if config.TotalMatches == 0 || config.TeamsPerMatch == 0 {
		// TODO: Add logging
		return matches
	}

	for i := 0; i < config.TotalMatches; i++ {
		sliceStart := i * config.TeamsPerMatch
		sliceEnd := sliceStart + config.TeamsPerMatch

		matches = append(matches, &PlayOffMatch{
			Teams: teams[sliceStart:sliceEnd],
		})
	}

	return matches
}

// SerpentinePlayOffOptimizer is a custom play off optimizer, which seeds the teams in the play off
// based on the performance. It works in the following way:
// 1. We're ordering all teams which supposed to play in the tournament by the following criterias:
// - By points (as normally we do)
// - By win/loss matches for teams. If they are equal compare by the next rule.
// - By combat/support/objective scores. If they are equal compare by the next rule.
// - By KDA between teams
// 2. Seed the teams according to the serpentine system
// For more info: https://en.wikipedia.org/wiki/Serpentine_system
type SerpentinePlayOffOptimizer struct{}

// PrepareTeams organizes teams for the play off based on the serpentine system (i.e. snake seeding)
func (s *SerpentinePlayOffOptimizer) PrepareTeams(config PlayOffTeamOptimizerConfig, teams map[string]*Team, groupStageTeamStats map[string]*TeamStats) []*PlayOffMatch {
	teamsRanking := GetTeamsRankingByPerformance(groupStageTeamStats)
	teamToRankMapping := make(map[*Team]int)
	for index, teamId := range teamsRanking {
		team := teams[teamId]
		teamToRankMapping[team] = index
	}

	// Pre-pass: sort the teams based on the given ranking
	allTeams := make([]*Team, 0)
	for _, team := range teams {
		allTeams = append(allTeams, team)
	}

	sort.Slice(allTeams, func(i, j int) bool {
		teamA := allTeams[i]
		rankTeamA := teamToRankMapping[teamA]

		teamB := allTeams[j]
		rankTeamB := teamToRankMapping[teamB]

		return rankTeamA < rankTeamB
	})

	// Actual snake seeding implementation
	// Step 1: Spawn N matches that we sequentially fill in
	matches := make([]*PlayOffMatch, 0)
	for i := 0; i < config.TotalMatches; i++ {
		matches = append(matches, &PlayOffMatch{
			Teams: make([]*Team, 0),
		})
	}

	// Step 2: Put teams in a snake pattern.
	teamIndex := 0
	for i := 0; i < config.TeamsPerMatch; i++ {
		if i%2 == 0 {
			// Left to right
			for j := 0; j < config.TotalMatches; j++ {
				if teamIndex < len(teams) {
					matches[j].Teams = append(matches[j].Teams, allTeams[teamIndex])
					teamIndex += 1
				}
			}
		} else {
			// Right to left
			for j := config.TotalMatches - 1; j >= 0; j-- {
				if teamIndex < len(teams) {
					matches[j].Teams = append(matches[j].Teams, allTeams[teamIndex])
					teamIndex += 1
				}
			}
		}
	}

	return matches
}
