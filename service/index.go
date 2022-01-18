package service

import (
	"fmt"

	"github.com/tradersclub/TCInterview/entiti"
	"github.com/tradersclub/TCInterview/repositories"
)

func ProcessUser(u string, p string) bool {

	if repositories.Process_user(p, u) {
		fmt.Print("here")
		return true
	}
	return false
}

func CreateUser(user entiti.User) bool {

	if repositories.CREATEUSER(user) {
		fmt.Print("here2")
		return true
	}
	return false
}
