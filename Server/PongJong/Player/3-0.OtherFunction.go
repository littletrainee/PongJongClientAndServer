package Player

// Check meld
func Ismeld(targetmeld []string) bool {
	first, second, third := targetmeld[0], targetmeld[1], targetmeld[2]
	if first == second && second == third {
		return true
	} else {
		return false
	}
}

// check target hand is forming
func iswin(targetslice []string) bool {
	// declare
	var tempplayer Player
	// sort targetslice
	tempplayer.Hand.Set(targetslice)
	tempplayer.SortHand()
	targetslice = tempplayer.Hand.Get()

	// check 3-set is all meld
	for i := 0; i < 9; i += 3 {
		if !Ismeld([]string{targetslice[i], targetslice[i+1], targetslice[i+2]}) {
			break
		}
		if i == 6 {
			return true
		}
	}
	return false
}
