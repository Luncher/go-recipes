package mp

import (
	"fmt"
	"time"
)

type MP3Player struct {
	stat int
	progress int
}

func (p *MP3Player)Play(source string) {
	fmt.Println("Playing MP3 music", source)

	p.progress = 0

	for o.progress < 100 {
		time.Sleep(100 * time.Microsecond)
		fmt.Print(".")
		p.progress += 10
	}

	fmt.Println("\nFinisied playing", source)
}