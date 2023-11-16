package menu

import (
	"fmt"
	"log"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

type Menu struct {
	options []Option
	LastInput string
	HadInputAlready bool
}

type Option struct {
	id string
	keyToListen string
	text string
}

func (menu *Menu) AddOption(optionId string, keyToListenString string, optionText string){
	menu.options = append(menu.options, Option{optionId, keyToListenString, optionText})
}

func (menu *Menu) Listen() string{
	if(len(menu.options) == 0){
		log.Fatalf("menu.Listen: Can't listen to a menu with zero options!")
	}

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		for _, option := range menu.options {
			if key.String() == option.keyToListen {
				menu.LastInput = option.id;
				menu.HadInputAlready = true;
				return true, nil
			}
		}
		
		return false, nil
	})

	return menu.LastInput;
}

func (menu *Menu) RenderMenu(){
	for _, option := range menu.options {
		fmt.Println(option.text)
	}
}