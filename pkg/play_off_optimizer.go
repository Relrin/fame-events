package pkg

type PlayOffOptimizer interface {
	PrepareTeams(teams []*Team) []*Team
}
