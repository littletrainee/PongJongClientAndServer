package PrintWinnerAndScore

func (ps *PrintWinnerAndScore) SetYakuAndPoint() {
	var oya, koya uint8

	ps.YakuType = ChineseNumber[ps.Color] + "色  " + ChineseNumber[ps.Kind] + "種"
	switch ps.Color {
	case 1:
		switch ps.Kind {
		case 1:
			oya = 200
			koya = 100
		case 2:
			oya = 60
			koya = 30
		case 3:
			oya = 60
			koya = 30
		}
	case 2:
		switch ps.Kind {
		case 1:
			oya = 80
			koya = 40
		case 2:
			oya = 30
			koya = 15
		case 3:
			oya = 10
			koya = 5
		}
	case 3:
		switch ps.Kind {
		case 1:
			oya = 60
			koya = 30
		case 2:
			oya = 10
			koya = 5
		case 3:
			oya = 30
			koya = 15
		case 9:
			oya = 6
			koya = 3
		}

	}
	ps.Score = []uint8{oya, koya}
}
