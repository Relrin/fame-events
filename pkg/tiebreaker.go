package pkg

type TiebreakResolver interface {
	IsNewMatchRequired(gs *GroupStage) bool
	DetermineRanking(gs *GroupStage) []uint32
}

struct