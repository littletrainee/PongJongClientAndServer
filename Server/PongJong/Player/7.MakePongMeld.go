package Player

import "github.com/littletrainee/slices"

// Make Pong Meld to Meld slice
func (p *Player) MakePongMeld(otherplayer *Player) {
	var (
		temphand         []string   = p.Hand.Get()
		tempmeldslices   [][]string = p.Meld.Get()
		tempmeld         []string
		otherplayerriver []string = otherplayer.River.Get()
		index            int
	)
	// Add three element to tempmeld
	for i := 0; i < 3; i++ {
		tempmeld = append(tempmeld, otherplayerriver[len(otherplayerriver)-1])
	}
	// Remove two element from temphand
	for i := 0; i < 2; i++ {
		index = slices.FindIndexOfElement(temphand, otherplayerriver[len(otherplayerriver)-1])
		temphand = append(temphand[:index], temphand[index+1:]...)
	}

	// remove last element from otherplayerriver
	otherplayerriver = otherplayerriver[:len(otherplayerriver)-1]

	// append tempmeld to tempmeldlist
	tempmeldslices = append(tempmeldslices, tempmeld)

	// Set each value
	p.HasPongMeld.Set(false)
	otherplayer.River.Set(otherplayerriver)
	p.Hand.Set(temphand)
	p.Meld.Set(tempmeldslices)
}
