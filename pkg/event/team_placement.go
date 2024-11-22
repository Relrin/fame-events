package event

type TeamPlacement struct {
	Name       string
	BrandIndex uint32
	Matches    uint32
	Wins       uint32
	Loses      uint32
	Points     uint32
}

type AggregatedTeamData struct {
	TeamId    string
	TeamStats *TeamStats
}
