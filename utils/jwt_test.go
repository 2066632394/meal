package utils

import (
	"testing"
	"time"
	"fmt"
)

func TestJwt(t *testing.T) {
	var et EasyToken
	et.Expires = time.Now().Add(time.Hour * 10).Unix()
	et.Username = "test123"
	token,err := et.GetToken()
	if err != nil {
		panic(err)
	}
	fmt.Print("token",token)
	check,err := et.ValidateToken(token)
	if err != nil {
		panic(err)
	}
	if check {
		fmt.Println("PASS")
	}

}
