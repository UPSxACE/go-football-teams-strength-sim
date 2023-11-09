package game

import (
	"fmt"

	"github.com/UPSxACE/go-football-teams-strength-sim/teams"
	"github.com/UPSxACE/go-football-teams-strength-sim/utils"
)

func EndingScreen() {
	fmt.Println(teams.List())
	fmt.Println(utils.LineMessage("THANK YOU!"))
}