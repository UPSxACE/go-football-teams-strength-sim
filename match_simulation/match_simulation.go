package match_simulation

import (
	"log"
	"math/rand"

	"github.com/UPSxACE/go-football-teams-strength-sim/teams"
)

type MatchSimulation struct {
	homeTeam  *teams.Team
	awayTeam  *teams.Team
	scoreHome int
	scoreAway int
	minute    int
	over      bool
}

func MakeMatchSimulation(teamHome *teams.Team, teamAway *teams.Team) MatchSimulation {
	if teamHome == teamAway {
		log.Fatalf("match.Init: cannot initialize a match between the same team")
	}
	return MatchSimulation{teamHome, teamAway, 0, 0, 0, false}
}

func (match *MatchSimulation) Next() (scoreHome int, scoreAway int, isOver bool) {
	relativeStrength1 := (float64(match.homeTeam.Strength) / float64((match.homeTeam.Strength + match.awayTeam.Strength))) * 100
	relativeStrength2 := (float64(match.awayTeam.Strength) / float64((match.homeTeam.Strength + match.awayTeam.Strength))) * 100
	if match.minute < 90 {
		roundPower1 := rand.Intn(2000)
		roundPower2 := rand.Intn(2000)
		homeTeamActivated := roundPower1 <= int(relativeStrength1)
		awayTeamActivated := roundPower2 <= int(relativeStrength2)
		if homeTeamActivated && !awayTeamActivated ||
			(homeTeamActivated && awayTeamActivated && roundPower1 > roundPower2) {
			match.scoreHome++;
		}
		if awayTeamActivated && !homeTeamActivated ||
			(awayTeamActivated && homeTeamActivated && roundPower2 > roundPower1) {
			match.scoreAway++;
		}
		match.minute++
	}
	if match.minute >= 90 {
		match.over = true
	}
	return match.scoreHome, match.scoreAway, match.over
}

func (match *MatchSimulation) IsOver() bool {
	return match.over
}
