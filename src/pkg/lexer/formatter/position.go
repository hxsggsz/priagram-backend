package formatter

import (
	"math/rand"
	"time"
)

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func GeneratePosition() Position {
	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	return Position{
		X: randGenerator.Intn(400), Y: randGenerator.Intn(400),
	}
}
