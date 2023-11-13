package main

import (
	"github.com/UPSxACE/go-football-teams-strength-sim/game"
)

func main(){
	game.OpeningScreen()
	game.CreateTeams()
	game.SimulateTournament()
	game.EndingScreen()
}