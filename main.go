package main

import (
	"fmt"
	"time"
)

var nowFunc = time.Now

func NewNow() time.Time {
	return nowFunc()
}

func main(){
	fmt.Print(NewNow())
}

