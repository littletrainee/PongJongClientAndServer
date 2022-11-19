package Player

import "github.com/littletrainee/PongJong_Client_And_Server/Server/PongJong/Wall"

// draw card from wall to hand
func (p *Player) DrawCard(wall *Wall.Wall) {
	// declare temp vairable
	tempwall := wall.Wall.Get()
	temphand := p.Hand.Get()
	// append tempwall first element to temphand
	temphand = append(temphand, tempwall[0])
	// remove first element
	tempwall = tempwall[1:]
	// reset to each slice
	p.Hand.Set(temphand)
	wall.Wall.Set(tempwall)
}
