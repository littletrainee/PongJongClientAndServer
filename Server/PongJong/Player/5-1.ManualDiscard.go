package Player

import (
	"log"
	"strconv"

	"github.com/littletrainee/GetSet"
)

func (p *Player) ManualDiscard(decision GetSet.Type[string]) {
	var (
		key       int
		temphand  []string = p.Hand.Get()
		tempriver []string = p.River.Get()
	)
	key, err := strconv.Atoi(decision.Get())
	if err != nil {
		log.Println(err)
	}

	tempriver = append(tempriver, temphand[key-1])
	temphand = append(temphand[:key-1], temphand[key:]...)
	p.Hand.Set(temphand)
	p.River.Set(tempriver)
}
