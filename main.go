package main

import (
	"fmt"
	"os"
	"time"

	"github.com/franz899/talking-clock/m/v2/clock"
)

func main() {
	var timeToParse string

	if len(os.Args) == 2 {
		timeToParse = os.Args[1]
	} else {
		timeToParse = getCurrentTime()
	}

	phrase, err := clock.Talk(timeToParse)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", phrase)
}

func getCurrentTime() string {
	currentTime := time.Now()
	currentHour := currentTime.Hour()
	currentMinute := currentTime.Minute()
	return fmt.Sprintf("%d:%d", currentHour, currentMinute)
}
