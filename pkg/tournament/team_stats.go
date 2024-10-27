package tournament

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
