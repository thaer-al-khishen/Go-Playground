package main

import (
	"fmt"
	"github.com/wagslane/go-tinytime"
	"time"
)

//Run go get REMOTE_REPO

func main() {
	tt := tinytime.New(1585750374)

	tt = tt.Add(time.Hour * 48)
	fmt.Println(tt)
	// prints 2020-04-03T14:12:54Z
}
