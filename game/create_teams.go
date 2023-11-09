package game

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/UPSxACE/go-football-teams-strength-sim/teams"
	"github.com/UPSxACE/go-football-teams-strength-sim/utils"
)

func CreateTeams() {
	reader := bufio.NewReader(os.Stdin)
	selectedOption := ""

	for selectedOption != "n" && selectedOption != "N" {
		selectedOption = ""

		teamName, err := utils.Ask(reader, "Name of the team:")

		if err != nil {
			log.Fatalf("Unexpected error: %v", err)
		}

		teamStrength, err := utils.Ask(reader, "Strength of the team:")

		if err != nil {
			log.Fatalf("Unexpected error: %v", err)
		}

		teamStrengthInt, err := strconv.Atoi(teamStrength)

		if err != nil {
			fmt.Println("Invalid input on team name or strength.")
			err = nil
		} else {
			_, err = teams.Create(teamName, teamStrengthInt)

			if err != nil {
				fmt.Println("Invalid input on team name or strength.")
				err = nil
			}
		}

		if len(teams.List()) >= 2 {
			for selectedOption != "n" && selectedOption != "N" && selectedOption != "y" && selectedOption != "Y" {
				answer, err := utils.Ask(reader, "Do you want to create another team? (y/n):")

				if err != nil {
					log.Fatalf("Unexpected error: %v", err)
				}

				selectedOption = answer
			}
		}
	}

	utils.Clear()
}