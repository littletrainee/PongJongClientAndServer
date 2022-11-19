package EnumHolder

type Action int

const (
	None Action = iota
	CheckTsumo
	IsTsumo
	CheckCanDeclareRiiChi
	CanDeclareRiiChi
	MakeDeclareRiiChiAndDiscard
	DisCard
	IsDisCard
	CheckRon
	CanRon
	IsRon
	NotRon
	CheckMeld
	CanPong
	IsPong
	NotPong
	DrawCard

	PrintWinnerAndScore
)
