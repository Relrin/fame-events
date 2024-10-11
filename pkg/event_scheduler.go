package pkg

type EventScheduler struct {
	Name        string
	GroupStages []*GroupStage

	// Turn into config?
	TiebreakResolver *TiebreakResolver
	PlayOffOptimizer *PlayOffOptimizer
	// GroupStageOptimizer
}

func NewEventScheduler() *EventScheduler {
	return &EventScheduler{}
}

// WithGroupStages sets a list of group stages to handle for the event
func (es *EventScheduler) WithGroupStages(groupStages []*GroupStage) *EventScheduler {
	es.GroupStages = groupStages
	return es
}

// WithTiebreakResolver sets a new tiebreak resolver for the event
func (es *EventScheduler) WithTiebreakResolver(tiebreakResolver *TiebreakResolver) *EventScheduler {
	es.TiebreakResolver = tiebreakResolver
	return es
}

// WithPlayOffOptimizer sets a new playoff optimizer that runs after finishing all group stages
func (es *EventScheduler) WithPlayOffOptimizer(playOffOptimizer *PlayOffOptimizer) *EventScheduler {
	es.PlayOffOptimizer = playOffOptimizer
	return es
}

func (es *EventScheduler) Run() {
	// run an infinite loop where
	// 1. We pass through all group stages & schedule a next match whenever is possible within the group (it includes tiebreakers)
	// 2. Once we finished with all the matches, start a playoff by advancing N teams further. Also schedule a separate event for them
}
