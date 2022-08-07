package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/dannywolfmx/sintetizador/frequency"
	"github.com/dannywolfmx/sintetizador/keyevent"
	"github.com/eiannone/keyboard"
	"github.com/hajimehoshi/oto/v2"
)

func play(context *oto.Context, freq float64, duration time.Duration) oto.Player {
	p := context.NewPlayer(frequency.NewSineWave(freq, duration))
	p.Play()
	return p
}

var Keys = map[rune]keyevent.Keyevent{}

func run() error {
	c, ready, err := oto.NewContext(*frequency.SampleRate, *frequency.ChannelCount, *frequency.BitDepthInBytes)
	if err != nil {
		return err
	}
	<-ready

	var wg sync.WaitGroup
	var players []oto.Player

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	keyEvents, err := keyboard.GetKeys(10)

	fmt.Println("Preciona ESC para salir")
	KeyNote := map[rune]keyevent.Keyevent{
		'1': keyevent.NewKeyevent(c, frequency.A4),
		'2': keyevent.NewKeyevent(c, frequency.AB4),
		'3': keyevent.NewKeyevent(c, frequency.B4),
		'4': keyevent.NewKeyevent(c, frequency.C5),
		'5': keyevent.NewKeyevent(c, frequency.CB5),
		'6': keyevent.NewKeyevent(c, frequency.E5),
		'7': keyevent.NewKeyevent(c, frequency.F5),
		'8': keyevent.NewKeyevent(c, frequency.FB5),
		'9': keyevent.NewKeyevent(c, frequency.G5),
		'0': keyevent.NewKeyevent(c, frequency.GB5),
	}
	for {
		event := <-keyEvents
		if event.Err != nil {
			panic(event.Err)
		}

		fmt.Printf("rune %q, key %X\r\n", event.Rune, event.Key)
		if event.Key == keyboard.KeyEsc {
			break
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			key, ok := KeyNote[event.Rune]
			if ok {
				key.Press()
			}

		}()
	}

	wg.Wait()

	// Pin the players not to GC the players.
	runtime.KeepAlive(players)

	return nil
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		panic(err)
	}
}
