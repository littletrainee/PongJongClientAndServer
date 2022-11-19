package Player

// check player is Tsumo
func (p *Player) TsumoCheck() {
	// declare
	var temphand []string = p.Hand.Get()

	// append meld to temphand
	if len(p.Meld.Get()) != 0 {
		for _, v := range p.Meld.Get() {
			temphand = append(temphand, v...)
		}
	}

	// check is win
	p.Iswin.Set(iswin(temphand))
}
