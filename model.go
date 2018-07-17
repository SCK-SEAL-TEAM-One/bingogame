package bingogame

type Ticket struct {
	SizeX int
	SizeY int
	Grid  [][]State
}

type State struct {
	Number int
	Status bool
}

type Player struct {
	Name   string
	Ticker Ticket
}

type Game struct {
	Player    []Player
	NumberBox []int
}
