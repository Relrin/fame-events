package pkg

type TiebreakResolver interface {
	IsDeciderMatchRequired(gs *GroupStage) bool
	DetermineRanking(gs *GroupStage) []uint32
}

// ScoreTiebreakResolver is a tiebreak resolver that decides to advance
// teams based on the following rules:
// - Check win/loss matches for teams. If they are equal compare by the next rule.
// - Check and compare combat/support/objective scores. If they are equal compare by the next rule.
// - Check by KDA. If they are equal compare by the next rule.
// - Random selection as the last resort
type ScoreTiebreakResolver struct {
}

// IsDeciderMatchRequired returns a boolean value that indicates that a new match
// needs to be schedules as a decider match
func (s *ScoreTiebreakResolver) IsDeciderMatchRequired(gs *GroupStage) bool {
	return false
}

// DetermineRanking returns a final list of teams indices sorted by the given rules. The
// returns list is expected to be used as the main decider which teams should go further.
func (s *ScoreTiebreakResolver) DetermineRanking(gs *GroupStage) []uint32 {
	// TODO: Implement rules
	return []uint32{}
}
