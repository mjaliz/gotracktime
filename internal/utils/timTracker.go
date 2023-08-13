package utils

import "time"

var timeChan = make(chan time.Duration)

func trackTime() {
	start := time.Now().UTC()
	lastPrint := time.Now().UTC()
	for {
		spent := time.Since(start)
		if time.Since(lastPrint) > 1*time.Second {
			timeChan <- spent
			lastPrint = time.Now().UTC()
		}
	}
}
