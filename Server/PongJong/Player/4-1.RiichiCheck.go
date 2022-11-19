package Player

import (
	PJ "github.com/littletrainee/PongJong_Client_And_Server/Server/PongJong"
	"github.com/littletrainee/slices"
)

// check player can riichi
func (p *Player) RiiChiCheck() []string {
	var (
		temphand           []string = p.Hand.Get()
		tempplayer         Player
		cloneforpopone     []string
		removeonefromclone []string
		probablywintile    []string
	)
	// append meld to temphand
	if len(p.Meld.Get()) != 0 {
		for _, v := range p.Meld.Get() {
			temphand = append(temphand, v...)
		}
	}

	// Sort temphand
	tempplayer.Hand.Set(temphand)
	tempplayer.SortHand()
	temphand = tempplayer.Hand.Get()

	// pop one by one to check is tenpai
	for i := range temphand {
		// reset cloneforpopone capacity and value
		cloneforpopone = make([]string, len(temphand))
		// clone from temphand to cloneforpopone[5]
		copy(cloneforpopone, temphand)
		// remove one from cloneforpopone array by temphand index [5->4]
		cloneforpopone = append(cloneforpopone[:i], cloneforpopone[i+1:]...)

		// try add one to array[4->5]
		for _, probablytarget := range PJ.Tile.Get() {
			// reset removeonefromclone capacity and value
			removeonefromclone = make([]string, len(cloneforpopone))
			// clone from cloneforpopone to removeoonefromclone[4]
			copy(removeonefromclone, cloneforpopone)
			removeonefromclone = append(removeonefromclone, probablytarget)
			if probablywin(removeonefromclone) {
				if !slices.ContainsElement(probablywintile, probablytarget) {
					probablywintile = append(probablywintile, probablytarget)
				}
			}
		}
	}
	return probablywintile
}
