package event

type GroupStageMatchResult struct {
	TeamPerformances []*TeamPerformanceEntry
}

type TeamPerformanceEntry struct {
	TeamId           string              // A unique team id assigned to the team for the entire tournament
	IsWinner         bool                // Did this team win within the one round group stage match
	Points           uint32              // Earned group stage points
	PlayerMatchStats []*PlayerMatchStats // Match stats per player
}

type PlayerMatchStats struct {
	PlayerId       string
	Kills          uint32
	Deaths         uint32
	Assists        uint32
	CombatScore    uint32
	ObjectiveScore uint32
	SupportScore   uint32
}
