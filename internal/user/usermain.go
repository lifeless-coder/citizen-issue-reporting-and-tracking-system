package user

import "fmt"

func UserMain() {
	var choice int
	fmt.Scanln(&choice)
	if choice == 1 {
		Register()
	} else {
		Login()
	}
}
