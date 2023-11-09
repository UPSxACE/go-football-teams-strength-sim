package game

import (
	"fmt"
	"github.com/UPSxACE/go-football-teams-strength-sim/utils"
)

func OpeningScreen(){
	fmt.Println(utils.LineMessage("WELCOME TO FOOTBALL SIMULATOR"))
	fmt.Println("We will start by adding teams, and then we will simulate the outcome of a tournament with those teams.")
	utils.PressAnyKey()
	utils.Clear()
}