package Player

import (
	"sync"

	"github.com/littletrainee/slices"
)

// check player has pong meld in hand
func (p *Player) CheckHasPongMeld(otherplayerriverlastone string, wg *sync.WaitGroup) {
	defer wg.Done()
	if slices.ContainsElement(p.Hand.Get(), otherplayerriverlastone) &&
		slices.CountNumber(p.Hand.Get(), otherplayerriverlastone) >= 2 {
		p.HasPongMeld.Set(true)
	}
}
