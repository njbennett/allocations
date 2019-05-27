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

	go func() {
		for t := range ticker.C {
			clicker <- t
			makeEngineer <- t
			xpEngineers <- t
		}
	}()

	register(clicker, func() {
		fmt.Println("Clickety-click")
	})

	engineers := make([]engineer, 0)

	register(makeEngineer, func() {
		engineers = append(engineers, engineer{})
		fmt.Println("A new engineer!")
	})

	register(xpEngineers, func() {
		for i := range engineers {
			engineers[i].ticks++
		}
		fmt.Println("Gave engineers some xp!")
	})

	time.Sleep(3000 * time.Millisecond)

	ticker.Stop()
	fmt.Println("Ticker stopped")

	fmt.Println(len(engineers))

	for _, e := range engineers {
		fmt.Println(e)
	}
}

func register(c chan time.Time, f func()) {
	go func() {
		for {
			<-c
			f()
		}
	}()
}
