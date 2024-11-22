package event

type TiebreakResolver interface {
	IsDeciderMatchRequired(gs *GroupStage) bool
	DetermineRanking(gs *GroupStage) []string
}

// TeamStatsTiebreakResolver is a tiebreak resolver that decides to advance
// teams based on the following rules:
// - Order by points (as normally we do)
// - Order by win/loss matches for teams. If they are equal compare by the next rule.
// - Order by combat/support/objective scores. If they are equal compare by the next rule.
// - Order by KDA between teams
type TeamStatsTiebreakResolver struct {
}

// IsDeciderMatchRequired returns a boolean value that indicates that a new match
// needs to be schedules as a decider match
func (t *TeamStatsTiebreakResolver) IsDeciderMatchRequired(gs *GroupStage) bool {
	return false
}

// DetermineRanking returns a final list of teams indices sorted by the given rules. The
// returns list is expected to be used as the main decider which teams should go further.
func (t *TeamStatsTiebreakResolver) DetermineRanking(gs *GroupStage) []string {
	return GetTeamsRankingByPerformance(gs.TeamStats)
}
