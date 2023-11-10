package tournaments

import "github.com/UPSxACE/go-football-teams-strength-sim/teams"

type League struct {
	currentPhase int
	totalPhases  int
	currentRound int
	totalRounds  int
	started      bool
	isOver       bool
	winner       string
	participants []teams.Team
	leaderboard  map[teamId]teamPoints
	schedule     map[roundNumber][]Match
}

type teamId int
type teamPoints int
type roundNumber int

type Match struct {
	homeTeam TeamInMatch
	awayTeam TeamInMatch
}

type TeamInMatch struct {
	teamId string
	teamScore int
}

func (league *League) HasStarted() bool{
	return false;
}
func (league *League) IsOver() bool{
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