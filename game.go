package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/UPSxACE/go-football-teams-strength-sim/teams"
	"github.com/UPSxACE/go-football-teams-strength-sim/utils"
)




func main(){
	
	
	
	reader := bufio.NewReader(os.Stdin)
	selectedOption := ""
	
	fmt.Println(utils.LineMessage("WELCOME TO FOOTBALL SIMULATOR"))
	fmt.Println("We will start by adding teams, and then we will simulate the outcome of a tournament with those teams.")
	
	utils.PressAnyKey()
	utils.Clear()

	for (selectedOption != "n" && selectedOption != "N") {
		selectedOption = "";

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
	fmt.Println(teams.List())
	fmt.Println(utils.LineMessage("THANK YOU!"))
}