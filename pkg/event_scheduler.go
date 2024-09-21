package pkg

type EventScheduler struct {
	GroupStages      []GroupStage
	TiebreakResolver *TiebreakResolver
}

func (es *EventScheduler) HasTiebreak(gs *GroupStage) bool {
	return false
}
