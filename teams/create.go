package teams

func Create(name string, strength int) Team{
	teamId := index;
	index++;
	newTeam := Team{teamId,name,strength}
	listOfTeams = append(listOfTeams, newTeam)
	return newTeam
}
