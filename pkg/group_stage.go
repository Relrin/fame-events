package pkg

type GroupStage struct {
	Name      string
	Teams     []*Team
	Standings map[string]*Standing
}

type Standing struct {
	// Generic stats
	PlayerMatches uint32
	Wins          uint32
	Loses         uint32
	Points        uint32

	// Team specific stats
	Kills          uint32
	Deaths         uint32
	Assists        uint32
	CombatScore    uint32
	ObjectiveScore uint32
	SupportScore   uint32
}
