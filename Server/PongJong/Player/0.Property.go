package Player

import "github.com/littletrainee/GetSet"

type Player struct {
	Name        GetSet.Type[string]
	Hand        GetSet.Type[[]string]
	Meld        GetSet.Type[[][]string]
	River       GetSet.Type[[]string]
	CodeNumber  GetSet.Type[uint8]
	Bookmaker   GetSet.Type[bool]
	Iswin       GetSet.Type[bool]
	IsRiiChi    GetSet.Type[bool]
	HasPongMeld GetSet.Type[bool]
}
