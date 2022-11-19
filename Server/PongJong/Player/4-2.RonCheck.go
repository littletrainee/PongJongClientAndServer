package Player

import (
	"sync"
)

// Check player is Ron
func (p *Player) RonCheck(otherplayer *Player, wg *sync.WaitGroup) {
	defer wg.Done()
	// declare
	var temphand []string = p.Hand.Get()

	// append meld to temphand
	if len(p.Meld.Get()) != 0 {
		for _, v := range p.Meld.Get() {
			temphand = append(temphand, v...)
		}
	}

	// append river last one to temphand
	temphand = append(temphand, otherplayer.River.Get()[len(otherplayer.River.Get())-1])

	// check is win
	p.Iswin.Set(iswin(temphand))
}
