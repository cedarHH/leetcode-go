package concurrency

import (
	"fmt"
	"time"
)

func timeTimer() {
	timer := time.NewTimer(2 * time.Second)

	fmt.Println("Waiting for the timer...")

	<-timer.C
	fmt.Println("Timer expired")
}

func timeAfter() {
	fmt.Println("Waiting for 2 seconds...")

	<-time.After(2 * time.Second)

	fmt.Println("Time's up!")
}

func resetTimer() {
	timer := time.NewTimer(5 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("Timer expired")
	}()

	time.Sleep(2 * time.Second)
	if timer.Reset(3 * time.Second) {
		fmt.Println("Timer was reset")
	}

	time.Sleep(4 * time.Second)
}

func timerStop() {
	timer := time.NewTimer(5 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("Timer expired")
	}()

	time.Sleep(2 * time.Second)

	if timer.Stop() {
		fmt.Println("Timer stopped before expiration")
	}

	time.Sleep(3 * time.Second)
}

func timeTicker() {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func TimerExample() {
	timeTimer()
	timeAfter()
	resetTimer()
	timerStop()
	timerStop()
	timeTicker()
}
