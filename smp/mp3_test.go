package mp

import (
	"testing"
)

func TestPlayMP3(t *testing.T) {
	p := MP3Player{}
	p.Play("haha.mp3")

	if p.progress != 100 {
		t.Error("MP3Player.Play failed.")
	}
}