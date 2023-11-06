package main

import (
	"fmt"
	"github.com/UPSxACE/go-football-teams-strength-sim/teams"
)

func main(){
	fmt.Println("Hello world")
	teams.Create("Real Madrid", 90)
	fmt.Println(teams.List())
}