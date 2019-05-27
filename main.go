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
	engineers := make([]engineer, 0)

	go register(&listeners, clickerClosure())
	go register(&listeners, xpEngineersClosure(&engineers))
	go register(&listeners, makeEngineerClosure(&engineers))

	ticker := time.NewTicker(100 * time.Millisecond)

	go tick(ticker, &listeners)

	time.Sleep(3000 * time.Millisecond)

	ticker.Stop()
	fmt.Println("Ticker stopped")

	fmt.Println(len(engineers))
	fmt.Println(listeners)

	for _, e := range engineers {
		fmt.Println(e)
	}

}

func clickerClosure() func() {
	fmt.Println("creating the clicker closure")
	return func() {
		fmt.Println("Clickety-click")
	}
}

func makeEngineerClosure(e *[]engineer) func() {
	fmt.Println("creating the engineer closure")
	return func() {
		*e = append(*e, engineer{})
	}
}

func xpEngineersClosure(e *[]engineer) func() {
	fmt.Println("creating the xp closure")
	return func() {
		for i := range *e {
			(*e)[i].ticks++
		}
	}
}

func register(listeners *[]chan time.Time, f func()) {
	fmt.Println("registering a function")
	c := make(chan time.Time)
	*listeners = append(*listeners, c)

	for {
		<-c
		f()
	}
}

func tick(t *time.Ticker, listeners *[]chan time.Time) {
	for t := range t.C {
		fmt.Println("tick")
		for i, l := range *listeners {
			l <- t
			fmt.Println(i)
		}
	}
}
