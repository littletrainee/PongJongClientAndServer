package Player

// auto discard from hand to river
func (p *Player) AutoDiscard() {
	var (
		temphand  []string = p.Hand.Get()
		tempriver []string = p.River.Get()
	)
	tempriver = append(tempriver, temphand[len(temphand)-1])
	temphand = temphand[:len(temphand)-1]
	p.Hand.Set(temphand)
	p.River.Set(tempriver)
}
