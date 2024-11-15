package config

// GroupStageConfig defines the configuration for the group stage part of the event
type GroupStageConfig struct {
	Type                     GroupStageType
	TotalGroups              int32
	MaxWinnersPerGroupAmount int32
	TiebreakResolverType     GroupStageTiebreakResolverType
}

// GroupStageType defines a possible strategy of handling group stage part
type GroupStageType int16

const (
	Regular GroupStageType = 0
)

// GroupStageTiebreakResolverType defines a strategy of handling tiebreakers which
// can happen within the single group stage. It includes running additional games,
// resolving the qualified teams based on scores, etc.
type GroupStageTiebreakResolverType int16

const (
	NoTiebreakResolver        GroupStageTiebreakResolverType = 0
	TeamStatsTiebreakResolver GroupStageTiebreakResolverType = 1
)
