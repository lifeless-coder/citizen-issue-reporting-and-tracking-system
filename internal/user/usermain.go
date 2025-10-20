package user

import (
	"fmt"
	"net/http"
)

func UserDashboard(w http.ResponseWriter, r *http.Request) {

}
func UserMain() {
	var choice int
	fmt.Scanln(&choice)
	if choice == 1 {
		Register()
	} else {
		Login()
	}
}
