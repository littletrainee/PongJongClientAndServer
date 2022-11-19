package Player

func probablywin(temp []string) bool {
	var tempplayer Player
	tempplayer.Hand.Set(temp)
	tempplayer.TsumoCheck()
	return tempplayer.Iswin.Get()
}
