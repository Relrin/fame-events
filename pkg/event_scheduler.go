package pkg

import "time"

type EventScheduler struct {
	Name        string
	CreatedAt   time.Time
	StartedAt   time.Time
	GroupStages []*GroupStage

	// Turn into config?
	// TiebreakResolver *TiebreakResolver
	// PlayOffOptimizer
	// GroupStageOptimizer
}

func (es *EventScheduler) HasTiebreak(gs *GroupStage) bool {
	return false
}
