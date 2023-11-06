package teams

import (
	"errors"
	"regexp"
)

/*Create a team
	Name needs to use only normal characters and numbers
	Name cannot be empty
	Name cannot be bigger than 50 characters
	Strength needs to be a number between 1 and 99
*/
func Create(name string, strength int) (Team, error){
	if(name == ""){
		return Team{}, errors.New("teams.Create: empty name")
	}

	if(len(name) > 50){
		return Team{}, errors.New("teams.Create: name is bigger than 50 characters")
	}

	onlyLettersNumbersSpaces := regexp.MustCompile("^[a-zA-Z0-9\\s]*$")

	if(!onlyLettersNumbersSpaces.MatchString(name)){
		return Team{}, errors.New("teams.Create: invalid characters in name")
	}	

	teamId := teamListIndex;
	teamListIndex++;
	newTeam := Team{teamId,name,strength}
	listOfTeams = append(listOfTeams, newTeam)
	return newTeam, nil
}
