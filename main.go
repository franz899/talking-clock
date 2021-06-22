package main

import (
	"fmt"
	"os"
	"time"

	"github.com/franz899/talking-clock/m/v2/clock"
)

func main() {
	currentTime := time.Now()
	currentHour := currentTime.Hour()
	currentMinute := currentTime.Minute()

	fmt.Printf("%d:%d\n", currentHour, currentMinute)

	time := fmt.Sprintf("%d:%d", currentHour, currentMinute)
	phrase, err := clock.Talk(time)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", phrase)
}
