package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"unicode/utf8"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func Ask(reader *bufio.Reader, question string) (string, error) {
	fmt.Print(question + " ")

	input, err := reader.ReadString('\n')

	// convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)

	return strings.Replace(input, "\r", "", -1), err
}

func PressAnyKey() {
	fmt.Println("Press any key to continue...")

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		return true, nil
	})
}

func LineMessage(message string) string {
	var size int;
	size = utf8.RuneCountInString(message);

	if size >= 49 {
		return message;
	}

	dashes := 50-size
	leftDashes := dashes/2
	rightDashes := dashes-leftDashes
	result := strings.Repeat("-", leftDashes) + message + strings.Repeat("-", rightDashes)

	return result
}

func removeFromSlice(slice []interface{}, index int) (newSlice []interface{}, removedElement interface{}) {
	removed := slice[index]
    return append(slice[:index], slice[index+1:]...), removed
}

/*Remove from unordered slice

This function is greatly faster than removeFromSlice, because instead of moving every element to the left,
it just places the last element in the position of the removed element.
Because of that, it only makes sense to use it in a slice in which the order does not matter!
*/
func removeFromUnorderedSlice(slice []interface{}, index int) (newSlice []interface{}, removedElement interface{}) {
	removed := slice[index]
    slice[index] = slice[len(slice)-1]
    return slice[:len(slice)-1], removed
}

func initClear() func(){
	switch runtime.GOOS {
		case "linux":
			return func() { 
				cmd := exec.Command("clear")
				cmd.Stdout = os.Stdout
				cmd.Run()
			}
		

		case "windows":
			return func() {
				cmd := exec.Command("cmd", "/c", "cls")
				cmd.Stdout = os.Stdout
				cmd.Run()
			}

		default:
			log.Fatal("This OS is not supported :(")
			return func(){}
		}
		
}

var Clear = initClear()