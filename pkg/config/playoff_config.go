package config

// PlayOffConfig defines the configuration for the play off part of the event
type PlayOffConfig struct {
	Type                    PlayOffType
	TeamsPerMatch           int
	MaxWinnersPerMatchCount int
	TeamOptimizerType       PlayOffTeamOptimizerType
}

// PlayOffType defines the approach to handle the play off games
type PlayOffType int16

const (
	SingleElimination PlayOffType = 0
)

// PlayOffTeamOptimizerType defines the pre-processing strategy for teams before running any
// play off games
type PlayOffTeamOptimizerType int16

const (
	NoTeamOptimizer         PlayOffTeamOptimizerType = 0
	SerpentineTeamOptimizer PlayOffTeamOptimizerType = 1
)
