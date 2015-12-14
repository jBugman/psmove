package main

import (
	"log"
	"time"

	"github.com/drbig/perlin"
	"github.com/jBugman/psmove"
)

func glowLEDs(move psmove.Move) {
	gr := perlin.NewGenerator(2, 2, 3, 100)
	gg := perlin.NewGenerator(2, 2, 3, 200)
	gb := perlin.NewGenerator(2, 2, 3, 300)

	f := func(g *perlin.Generator, t float64) int {
		return int(254.0 * (g.Noise1D(t) + 1.0) / 2.0)
	}
	t := 0.0
	for {
		t += 0.03
		r, g, b := f(gr, t), f(gg, t), f(gb, t)
		// log.Println(r, g, b)
		move.SetLEDs(r, g, b)
		move.UpdateLEDs()
		time.Sleep(15 * time.Millisecond)
	}
}

func main() {
	move, err := psmove.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer move.Disconnect()
	log.Println(move.ConnectionType())
	glowLEDs(move)
}
