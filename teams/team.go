package teams

func Create(name string, strength int) Team{
	newTeam := Team{name,strength}
	listOfTeams = append(listOfTeams, newTeam)
	return newTeam
}
