package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	currentHour := currentTime.Hour()
	currentMinute := currentTime.Minute()

	fmt.Printf("%d:%d\n", currentHour, currentMinute)
}
