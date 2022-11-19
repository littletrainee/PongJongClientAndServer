package EnumHolder

type Pattern int

const (
	CheckConnect Pattern = iota
	JoinAGame
	GetRoomUUID
	GetGameData
	MakeADecisionToDiscard
	DeclareRiiChi
	MakeRon
	MakePong
)
