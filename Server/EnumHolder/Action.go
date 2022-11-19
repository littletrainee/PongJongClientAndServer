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
	IsDiscard
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
