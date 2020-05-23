package xtime

import (
	"fmt"
	"time"
)

const TAIPEI = "Asia/Taipei"

func GetTme() time.Time {
	t := time.Now()
	localLocation, err := time.LoadLocation(TAIPEI)
	if err != nil {
		fmt.Println(err)
	}
	return t.In(localLocation)
}
