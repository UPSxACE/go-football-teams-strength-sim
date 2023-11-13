package game

import (
	"github.com/UPSxACE/go-football-teams-strength-sim/tournaments"
)

func SimulateTournament() {
	var tournament tournaments.Tournament
	tournament = &tournaments.League{}

	tournament.Init()

	for tournament.IsOver() != true {
		tournament.Render()
		// print stuff and ask if next [also, check notes]
	}
}