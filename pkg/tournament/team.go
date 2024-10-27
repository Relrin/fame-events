package tournament

type Team struct {
	Name       string        // Displayed name in the event
	Id         string        // A unique id assigned to the team for the entire tournament
	BrandIndex uint32        // The team brand within the event
	Color      uint32        // Assigned color through the entire event
	Players    []*PlayerInfo // List of players assigned to the team
}

type PlayerInfo struct {
	PlayerId string // A unique name
}
