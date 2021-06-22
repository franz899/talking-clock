package main

import (
	"fmt"
	"net/http"

	"github.com/franz899/talking-clock/m/v2/clock"
)

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	var timeToParse string
	var timeQ = req.URL.Query().Get("time")

	if len(timeQ) > 0 {
		timeToParse = timeQ
	} else {
		timeToParse = clock.GetCurrentTime()
	}

	time, err := clock.Talk(timeToParse)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, time)
}
