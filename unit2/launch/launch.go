package main

import "fmt"

func main() {
	launch := false

	launchText := fmt.Sprintf("%v ", launch)
	fmt.Println("Ready for launch:", launchText)
	var yesNo string

	if launch {
		yesNo = "Yes"
	} else {
		yesNo = "No"
	}

	fmt.Println("Ready for launch:", yesNo)
}
