package keyevent

import (
	"time"

	"github.com/dannywolfmx/sintetizador/frequency"
	"github.com/hajimehoshi/oto/v2"
)

type Keyevent struct {
	ctx    *oto.Context
	freq   float64
	player oto.Player
}

func NewKeyevent(context *oto.Context, freq float64) Keyevent {
	return Keyevent{
		ctx:    context,
		freq:   freq,
		player: context.NewPlayer(frequency.NewSineWave(freq, 3*time.Second)),
	}

}

func (k *Keyevent) Press() {
	if k.player.IsPlaying() {
		return
	}

	k.player.Reset()
	k.player = k.ctx.NewPlayer(frequency.NewSineWave(k.freq, 3*time.Second))
	k.player.Play()
}

func (k *Keyevent) Realese() {
	k.player.Close()
}

var KeyNote = map[rune]float64{
	'1': frequency.A4,
	'2': frequency.AB4,
	'3': frequency.B4,
	'4': frequency.C5,
	'5': frequency.CB5,
	'6': frequency.E5,
	'7': frequency.F5,
	'8': frequency.FB5,
	'9': frequency.G5,
	'0': frequency.GB5,
}
