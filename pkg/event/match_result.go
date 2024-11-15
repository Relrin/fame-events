package event

type GroupStageMatchResult struct {
	PointsPerTeam map[uint32]uint32
	WonTeams      map[uint32]*TeamStats
	LostTeams     map[uint32]*TeamStats
}
