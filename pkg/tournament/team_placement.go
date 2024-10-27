package tournament

type TeamPlacement struct {
	Name    string
	Matches uint32
	Wins    uint32
	Loses   uint32
	Points  uint32
}

type AggregatedTeamData struct {
	TeamIndex uint32
	TeamStats *TeamStats
}
