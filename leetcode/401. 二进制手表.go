package main

func readBinaryWatch(turnedOn int) []string {
	var r []string
	if turnedOn >= 9 {
		return r
	}
	for h := uint8(0); h < 12; h++ {
		h1 := bits.OnesCount8(h)
		for m := uint8(0); m < 60; m++ {
			m1 := bits.OnesCount8(m)
			if h1+m1 == turnedOn {
				r = append(r, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return r
}
