package digo

import (
	"fmt"
	"os"
	"testing"
)

func TestGetUserData(t *testing.T) {
	username := os.Getenv("DIUsername")
	password := os.Getenv("DIPassword")
	user := GetUserData(username, password)
	fmt.Println(user)
}
