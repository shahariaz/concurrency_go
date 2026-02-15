package main

import (
	"errors"
	"math/rand"
)

type FootBallPlayer struct {
	stamina int
	power   int
}

type CR7 struct {
	stamina   int
	power     int
	technique int
}
type Player interface {
	KickBall()
}

func (f FootBallPlayer) KickBall() {
	shot := f.stamina + f.power
	println("Shot power is:", shot)
}
func (c CR7) KickBall() {
	shot := c.stamina + c.power + c.technique
	println("CR7 Shot power is:", shot)
}

func TrowError() error {

	return errors.New("New  ERROR")
}

func main() {
	if err := TrowError(); err != nil {
		println("Error:", err.Error())
	}
	team := make([]Player, 11)

	for i := range team {
		if i == len(team)-1 {
			team[i] = CR7{

				stamina:   rand.Intn(10),
				power:     rand.Intn(10),
				technique: rand.Intn(10),
			}
			continue
		}
		team[i] = FootBallPlayer{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}
	for i := range team {
		team[i].KickBall()
	}

}
