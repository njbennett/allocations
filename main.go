package main

import "fmt"

func main() {
	history := make([]pool, 0)

	history = start(history)

	for i := 0; i < 52; i++ {
		history = increment(history)
	}

	stats := calculateStats(history)
	fmt.Printf("Weeks: %d Engineers: %d", stats.weeks, stats.engineers)
}

type pool struct {
	engineers []engineer
}

type engineer struct {
	power int
}

type stats struct {
	weeks     int
	engineers int
}

func start(history []pool) []pool {
	startingPool := pool{}
	history = append(history, startingPool)
	return history
}

func increment(history []pool) []pool {
	lastPool := history[len(history)-1]
	newPool := addEngineer(lastPool)

	history = append(history, newPool)
	return history
}

func addEngineer(p pool) pool {
	e := engineer{}
	p.engineers = append(p.engineers, e)
	return p
}

func calculateStats(history []pool) stats {
	weeks := len(history)
	engineers := len(history[len(history)-1].engineers)
	calculatedStats := stats{weeks: weeks, engineers: engineers}
	return calculatedStats
}
