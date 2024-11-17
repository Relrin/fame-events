package event

type TeamStats struct {
	// Group stage stats
	TotalMatches uint32
	WonMatches   uint32
	LostMatches  uint32
	Points       uint32

	// Team specific stats
	Kills          uint32
	Deaths         uint32
	Assists        uint32
	CombatScore    uint32
	ObjectiveScore uint32
	SupportScore   uint32
}

// Update returns latest team stats based on the given players information
func (teamStats *TeamStats) Update(playerStats []*PlayerStats) *TeamStats {
	if playerStats == nil {
		return teamStats
	}

	for _, stats := range playerStats {
		teamStats.Kills += stats.Kills
		teamStats.Deaths += stats.Deaths
		teamStats.Assists += stats.Assists
		teamStats.CombatScore += stats.CombatScore
		teamStats.ObjectiveScore += stats.ObjectiveScore
		teamStats.SupportScore += stats.SupportScore
	}

	return teamStats
}
