package event

import (
	"github.com/rs/xid"
)

type Team struct {
	Name        string                  // Displayed name in the event
	Id          string                  // A unique id assigned to the team for the entire tournament
	BrandIndex  uint32                  // The team brand within the event
	Color       uint32                  // Assigned color through the entire event
	Players     []*PlayerInfo           // List of players assigned to the team
	PlayerStats map[string]*PlayerStats // Gathered stats per player for the entire event
}

type PlayerInfo struct {
	PlayerId string // A unique id for the given player
	PartyId  string // A unique id that refers to the existing (social) party
}

type PlayerStats struct {
	Kills          uint32
	Deaths         uint32
	Assists        uint32
	CombatScore    uint32
	ObjectiveScore uint32
	SupportScore   uint32
}

// InitTeam returns a new instance of the Team type
func InitTeam(players []*PlayerInfo) *Team {
	playerStats := make(map[string]*PlayerStats)
	for _, playerInfo := range players {
		playerStats[playerInfo.PlayerId] = &PlayerStats{}
	}

	return &Team{
		Name:        "",
		Id:          xid.New().String(),
		BrandIndex:  0,
		Color:       0,
		Players:     players,
		PlayerStats: playerStats,
	}
}

// WithName overrides the default name for the given team
func (team *Team) WithName(name string) *Team {
	team.Name = name
	return team
}

// WithBrandIndex overrides the default brand index for the given team
func (team *Team) WithBrandIndex(brandIndex uint32) *Team {
	team.BrandIndex = brandIndex
	return team
}

// WithColor overrides the default color for the given team
func (team *Team) WithColor(color uint32) *Team {
	team.Color = color
	return team
}
