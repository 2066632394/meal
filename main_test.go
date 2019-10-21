package main

import (
	"testing"
	"time"
	"fmt"
)

func TestTime(t *testing.T) {
	str := "2019-10-21"
	date ,_ := time.Parse("2006-01-02",str)
	fmt.Println("date",date)
}
