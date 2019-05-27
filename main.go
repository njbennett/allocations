package main

import (
	"fmt"
	"time"
)

type engineer struct {
	ticks int
}

func main() {
	clicker := make(chan time.Time)
	makeEngineer := make(chan time.Time)
	xpEngineers := make(chan time.Time)

	ticker := time.NewTicker(100 * time.Millisecond)

	go tick(ticker, clicker, makeEngineer, xpEngineers)

	go register(clicker, func() {
		fmt.Println("Clickety-click")
	})

	engineers := make([]engineer, 0)

	go register(makeEngineer, makeEngineerClosure(&engineers))

	go register(xpEngineers, xpEngineersClosure(&engineers))

	time.Sleep(3000 * time.Millisecond)

	ticker.Stop()
	fmt.Println("Ticker stopped")

	fmt.Println(len(engineers))

	for _, e := range engineers {
		fmt.Println(e)
	}
}

func makeEngineerClosure(e *[]engineer) func() {
	return func() {
		*e = append(*e, engineer{})
		fmt.Println("A new engineer! from a closure")
	}
}

func xpEngineersClosure(e *[]engineer) func() {
	return func() {
		for i := range *e {
			(*e)[i].ticks++
		}
		fmt.Println("Gave engineers some xp!")
	}
}

func register(c chan time.Time, f func()) {
	for {
		<-c
		f()
	}
}

func tick(t *time.Ticker, c, m, x chan time.Time) {
	for t := range t.C {
		c <- t
		m <- t
		x <- t
	}
}
