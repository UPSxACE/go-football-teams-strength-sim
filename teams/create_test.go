package teams

import (
	"testing"
)

func TestCreate(t *testing.T){
	teamName := "This is a valid name 123"
	teamStrength := 50
	teamList := List()

	// Check if team is not in the list already before the test(it should not be)
	for _, team := range teamList{
		if team.Name == teamName {
			t.Fatal("TestCreate: invalid teamName; team already exists")
		}
	}

	initialIndex := teamListIndex

	createdTeam, err := Create(teamName, teamStrength)

	expectedId := createdTeam.Id == initialIndex
	expectedName := createdTeam.Name == teamName
	expectedStrength := createdTeam.Strength == teamStrength
	indexUpdated := teamListIndex == initialIndex+1

	if err != nil || !expectedId || !expectedName || !expectedStrength || !indexUpdated {
		t.Fatalf(`Create(%q, %v) =  %v, %q, %v, %v, want %v, %q, %v, nil`,
			teamName, teamStrength, createdTeam.Id, createdTeam.Name, createdTeam.Strength, err, initialIndex, teamName, teamStrength)
	}

	teamList = List()
	isInList := false

	// Check if team is in the list now
	for _, team := range teamList{
		if team.Name == teamName {
			isInList = true
		}
	}

	if !isInList {
		t.Fatalf("TestCreate: team was not added to the list, %q", teamList)
	}

}

func TestCreateEmptyName(t *testing.T){
	teamName := ""
	teamStrength := 50

	initialIndex := teamListIndex

	_, err := Create(teamName, teamStrength)

	// It's also checking if the if the index updated (it should not update since the team was not created)
	indexUpdated := teamListIndex == initialIndex+1

	if err == nil || indexUpdated {
		t.Fatalf(`Create(%q, %v) = %v, %v, want %v, error`,
			teamName, teamStrength, teamListIndex, err, initialIndex)
	}
}

func TestCreateInvalidStrength(t *testing.T){
	teamName := "This is a valid name"

	testStrengthError := func(invalidStrength int){
		initialIndex := teamListIndex

		_, err := Create(teamName, invalidStrength)
	
		// It's also checking if the if the index updated (it should not update since the team was not created)
		indexUpdated := teamListIndex == initialIndex+1
	
		if err == nil || indexUpdated {
			t.Fatalf(`Create(%q, %v) = %v, %v, want %v, error`,
				teamName, invalidStrength, teamListIndex, err, initialIndex)
		}
	}
	
	testStrengthError(-50)
	testStrengthError(0)
	testStrengthError(100)	
}

func TestCreateInvalidName(t *testing.T){
	teamStrength := 50

	testInvalidName := func(invalidName string){
		initialIndex := teamListIndex
	
		_, err := Create(invalidName, teamStrength)
	
		// It's also checking if the if the index updated (it should not update since the team was not created)
		indexUpdated := teamListIndex == initialIndex+1
	
		if err == nil || indexUpdated {
			t.Fatalf(`Create(%q, %v) = %v, %v, want %v, error`,
			invalidName, teamStrength, teamListIndex, err, initialIndex)
		}
	}

	testInvalidName("invalid!because special characters")
	testInvalidName("invalid#")
	testInvalidName(")=#")
	testInvalidName("this is a name that is too big because its bigger than 50 characters")
}