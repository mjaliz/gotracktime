package main

import (
	"fmt"
	"time"
)

var timeChan = make(chan time.Duration)

func main() {
	go trackTime()
	for {
		select {
		case t := <-timeChan:
			fmt.Println(t)
		}
	}
}

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
