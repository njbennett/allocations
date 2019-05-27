package main

import (
	"fmt"
	"time"
)

type engineer struct {
	ticks int
}

func main() {
	listeners := make([]chan time.Time, 0)

	go register(&listeners, func() {
		fmt.Println("Clickety-click")
	})

	engineers := make([]engineer, 0)

	go register(&listeners, xpEngineersClosure(&engineers))
	go register(&listeners, makeEngineerClosure(&engineers))

	ticker := time.NewTicker(100 * time.Millisecond)
	go tick(ticker, listeners)

	time.Sleep(3000 * time.Millisecond)

	ticker.Stop()
	fmt.Println("Ticker stopped")

	fmt.Println(len(engineers))
	fmt.Println(listeners)

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

func register(listeners *[]chan time.Time, f func()) {
	c := make(chan time.Time)
	*listeners = append(*listeners, c)

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
