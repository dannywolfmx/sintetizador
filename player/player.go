package player

import (
	"time"

	"github.com/dannywolfmx/sintetizador/frequency"
	"github.com/hajimehoshi/oto/v2"
)

func Play(context *oto.Context, freq float64, duration time.Duration) oto.Player {
	p := context.NewPlayer(frequency.NewSineWave(freq, duration))
	p.Play()
	return p
}
