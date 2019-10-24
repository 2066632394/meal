package utils

import (
	"testing"
	"fmt"
)

func TestJwt(t *testing.T) {
	str := RandomString(6)
	fmt.Println("STR",str)
}


