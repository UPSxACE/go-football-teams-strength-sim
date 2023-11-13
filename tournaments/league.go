package tournaments

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"

	"github.com/UPSxACE/go-football-teams-strength-sim/teams"
	"github.com/UPSxACE/go-football-teams-strength-sim/utils"
)

type League struct {
	currentPhase int
	totalPhases  int
	currentRound int
	totalRounds  int
	started      bool
	isOver       bool
	winner       string
	participants []teams.Team
	leaderboard  map[teamId]LeaderboardData
	schedule     map[roundNumber]MatchDays
}

type teamId int

type LeaderboardData struct {
	teamPoints int
	gamesPlayed int
}
type roundNumber int

type MatchDays map[int][]Match

type Match struct {
	homeTeam TeamInMatch
	awayTeam TeamInMatch
	playedAlready bool
}

type TeamInMatch struct {
	teamId teamId
	teamScore int
}

type calendarHelper map[teamId]calendarHelperData

type calendarHelperData struct {
	homeMatches int
	awayMatches int
}

func (league *League) HasStarted() bool{
	return league.started;
}
func (league *League) IsOver() bool{
	return league.isOver;
}
func (league *League) GetWinner() string{
	return league.winner;
}
func (league *League) Init(){
	fmt.Println(utils.LineMessage("LEAGUE BUILDER"))
	fmt.Println("All teams created will participate in the tournament.")
	fmt.Println("In this simulated tournament, in each round, every team will face each other once.")
	fmt.Println("How many rounds the tournament will have, is up to you to choose.")
	fmt.Println("Usually tournaments of this type have 2 rounds, but if there is only a few teams, more rounds is recommended sometimes.")

	reader := bufio.NewReader(os.Stdin)
	validInput := false
	var rounds int
	
	for validInput != true {
		input, err := utils.Ask(reader, "How many rounds you want the tournament to have:")

		if err != nil {
			log.Fatalf("Unexpected error: %v", err)
		}

		rounds, err = strconv.Atoi(input)

		if err != nil || rounds < 1{
			fmt.Println("Invalid number of rounds!")
			continue;
		}

		validInput = true
	}

	league.totalRounds = rounds;
	league.participants = teams.List()
	league.leaderboard = make(map[teamId]LeaderboardData);
	league.schedule = make(map[roundNumber]MatchDays)
	
	for _, team := range league.participants {
		league.leaderboard[teamId(team.Id)] = LeaderboardData{0,0}
	}

	var matchesPerDay int
	if(len(league.participants) >= 4){
		matchesPerDay = int(math.Floor(float64(len(league.participants)) / 4))
	 } else {
		matchesPerDay = 1
	 }
	
	 calendarHelper := calendarHelper{}
	 for _,team := range league.participants{
		calendarHelper[teamId(team.Id)] = calendarHelperData{0,0}
	 }
	for round := roundNumber(1); round <= roundNumber(rounds); round++ {
		// if i > 2

		
		league.schedule[round] = MatchDays{}
		idealDayForMatch := 1

		initialPool := make([]teamId,0, len(league.participants))
		for _, team := range league.participants {
			initialPool = append(initialPool, teamId(team.Id))
		}

		// Shuffle the teams in the initial pool
		rand.Shuffle(len(initialPool), func(i, j int) {
			initialPool[i], initialPool[j] = initialPool[j], initialPool[i]
		})

		poolQueue := []([]teamId){initialPool}

		for len(poolQueue) > 0 {
			poolRemovedFromQueue := poolQueue[0]
			poolQueue = poolQueue[1:]

			if(len(poolRemovedFromQueue)<= 1){
				continue;
			}

			// Sort pool by descending order of who already had the most home games
			sort.SliceStable(poolRemovedFromQueue, func(i int, j int) bool{
				return calendarHelper[poolRemovedFromQueue[i]].homeMatches > calendarHelper[poolRemovedFromQueue[j]].homeMatches
			})

			// Now, we will make the pool elements to be slightly mixed

			// This is a very important step, because later the pool will be split in half into two subpools,
			// and teams on the same subpool won't face each other.
			// Since the teams are ordered by who has the most games at home, this will make sure that there is an equal
			// (or almost equal) number of teams with simillar HOMEvsAWAY number of matches in both subpools
			// and this will guarantee that in the end of everything, all teams will have an equal(or almost equal)
			// number of Home and Away matches.

			// With "almost equal", I mean that in the case a tournament is composed by an even number of teams,
			// half of the teams will have one extra home game, and the other half will have one extra away game,
			// because all of them will have to face an odd number of opponents, on an odd number of matches
			// e.j: tournament with 18 teams means each team faces 17 opponents
			reorderedPool := make([]teamId,0, len(poolRemovedFromQueue))

			for index, team := range poolRemovedFromQueue {
				if(index%2 == 0){
					reorderedPool = append(reorderedPool, team)
				}
			}
			for index, team := range poolRemovedFromQueue {
				if(index%2 != 0){
					reorderedPool = append(reorderedPool, team)
				}
			}

			middleOfSlice := int(math.Ceil(float64(len(reorderedPool)) / 2));

			leftSubpool := reorderedPool[:middleOfSlice]
			rightSubpool := reorderedPool[middleOfSlice:]

			// This loop needs to run once for each team in the left subpool
			// The left subpool will always be the bigger one in case there is an odd number of teams in the subpool
			for i := 0; i < len(leftSubpool); i++{
				// This inner loop will match each team in the right subpool with someone in the left subpool
				// Since this will run once for each team in the left subpool, it will make sure every team in the
				// left subpool will face all the teams in the right pool once
				for j:= 0; j <len(leftSubpool); j++{
					// The modulus operator helps to safely circulate through slices
					teamFromRightSubpoolIndex := (j + i) % len(leftSubpool);
					if(teamFromRightSubpoolIndex >= len(rightSubpool)){
						continue;
					}

					// By default, teams on left subpool have home priority
					teamWithHomePriority := leftSubpool[j]
					teamWithoutPriority := rightSubpool[teamFromRightSubpoolIndex]

					_, ok := league.schedule[round][idealDayForMatch]
					if(!ok){
						// Initializing field needed
						league.schedule[round][idealDayForMatch] = []Match{}
					}
					

					if(len(league.schedule[round][idealDayForMatch]) >= matchesPerDay){
						idealDayForMatch++;
						_, ok := league.schedule[round][idealDayForMatch]
						if(!ok){
							// Initializing field needed
							league.schedule[round][idealDayForMatch] = []Match{}
						}
					}

					matchDay := idealDayForMatch;
					bothFree := false;

					for bothFree != true {
						_, ok := league.schedule[round][matchDay]
						if(!ok){
							// Initializing field needed
							league.schedule[round][matchDay] = []Match{}
						}
						

						if(len(league.schedule[round][matchDay]) >= matchesPerDay){
							_, ok := league.schedule[round][matchDay]
							if(!ok){
								// Initializing field needed
								league.schedule[round][matchDay] = []Match{}
							}
						}

						var teamsThatPlayThatDay []teamId;

						for _, team := range league.schedule[round][matchDay]{
							teamsThatPlayThatDay = append(teamsThatPlayThatDay, teamId(team.homeTeam.teamId))
							teamsThatPlayThatDay = append(teamsThatPlayThatDay, teamId(team.awayTeam.teamId))
						}

						notFree := false;
						// Check if one of the teams play that day
						for _, id := range teamsThatPlayThatDay {
							if(id == teamWithHomePriority || id == teamWithoutPriority){
								notFree = true;
								break;
							}
						}

						if(notFree){
							// If one of them plays that day, the match will have to be in another day
							matchDay++;
							continue;
						}

						// If none plays that day, both are free, and match will happen that day (quit the loop)
						bothFree = true

					}

					team1helper := calendarHelper[teamWithHomePriority]
					team2helper := calendarHelper[teamWithoutPriority]
					if(team2helper.homeMatches + team2helper.awayMatches >
						team1helper.homeMatches + team1helper.awayMatches){
							// If team without priority has more matches played, they will be the ones with home priority
							teamWithHomePriority, teamWithoutPriority = teamWithoutPriority, teamWithHomePriority
							team1helper, team2helper = team2helper, team1helper
					}

					home1 := team1helper.homeMatches
					away1 := team1helper.awayMatches
					home2 := team2helper.homeMatches
					away2 := team2helper.awayMatches
					
					situation := 0

					// smallest of the 4 has priority
					if(away2 <= home2 && away2 <= away1 && away2 <= home1){
						situation = 0;
					}
					if(home2 <= away2 && home2 <= away1 && home2 <= home1){
						situation = 1;
					}
					if(away1 <= away2 && away1 <= home2 && away1 <= home1){
						situation = 1;
					}
					if(home1 <= away2 && home1 <= home2 && home1 <= away1){
						situation = 0;
					}

					// if a team has 2 more home games than away games, it has priority for away games and vice-versa
					if(home1 - away1 > 2){
						situation = 1
					}
					if(away1 - home1 > 2){
						situation = 0
					}
					if(home2 - away2 > 2){
						situation = 0
					}
					if(away2 - home2 > 2){
						situation = 1
					}

					switch situation {
						case 0:
							match := Match{}
							match.homeTeam.teamId = teamWithHomePriority
							match.awayTeam.teamId = teamWithoutPriority
							league.schedule[round][matchDay] = append(league.schedule[round][matchDay], match)
							
							team1 := calendarHelper[teamWithHomePriority]
							team1.homeMatches += 1
							calendarHelper[teamWithHomePriority] = team1;

							team2 := calendarHelper[teamWithoutPriority]
							team2.awayMatches += 1
							calendarHelper[teamWithoutPriority] = team1;
							break;
						case 1:
							match := Match{}
							match.homeTeam.teamId = teamWithoutPriority
							match.awayTeam.teamId = teamWithHomePriority
							league.schedule[round][matchDay] = append(league.schedule[round][matchDay], match)
							
							team1 := calendarHelper[teamWithoutPriority]
							team1.homeMatches += 1
							calendarHelper[teamWithHomePriority] = team1;

							team2 := calendarHelper[teamWithoutPriority]
							team2.awayMatches += 1
							calendarHelper[teamWithoutPriority] = team1;
							break;
					}
					


				}
			}

			// Now make two new pools (not subpools) and add them to the queue
			// That means that half of teams on the left subpool will face other half of the teams in the left subpool,
			// and half of teams on the right subpool will face other half of the teams in the right subpool
			poolQueue = append(poolQueue, leftSubpool)
			poolQueue = append(poolQueue, rightSubpool)
		}
		



	}
}
func (league *League) NextPhase(){
}
func (league *League) Render(){
	fmt.Println(league.schedule)
	league.isOver = true
}