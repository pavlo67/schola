package e1

func SearchWithBisection(list []int, value int) bool {
	nMin := 0         // list index to start search
	nMax := len(list) // list index to stop  search

	for nMax > nMin {
		nMedium := nMin + (nMax-nMin)/2

		if list[nMedium] < value {
			nMin = nMedium + 1
		} else if list[nMedium] > value {
			nMax = nMedium
		} else { // if list[nMedium] == value
			return true
		}
	}

	return false
}
