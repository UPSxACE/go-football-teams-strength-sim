package tournaments

import "github.com/UPSxACE/go-football-teams-strength-sim/teams"

type League struct {
	currentPhase int
	totalPhases  int
	started      bool
	winner       string
	participants []teams.Team
	leaderboard map[int]int
}

func (league *League) HasStarted() bool{
	return false;
}
func (league *League) GetWinner() string{
	return "";
}
func (league *League) Init(){
}
func (league *League) NextPhase(){
}
func (league *League) Render(){
}