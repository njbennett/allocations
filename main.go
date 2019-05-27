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

	go register(clicker, func() {
		fmt.Println("Clickety-click")
	})

	engineers := make([]engineer, 0)

	go register(makeEngineer, makeEngineerClosure(&engineers))
	go register(xpEngineers, xpEngineersClosure(&engineers))

	ticker := time.NewTicker(100 * time.Millisecond)
	listeners := []chan time.Time{clicker, makeEngineer, xpEngineers}
	go tick(ticker, listeners)

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

func tick(t *time.Ticker, listeners []chan time.Time) {
	for t := range t.C {
		for _, l := range listeners {
			l <- t
		}
	}
}
