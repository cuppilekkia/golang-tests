package utils

func BSort(els []int) []int {
	k := true

	for k {
		k = false

		for i := 0; i < len(els)-1; i++ {
			if els[i] > els[i+1] {
				els[i], els[i+1] = els[i+1], els[i]
				k = true
			}
		}
	}

	return els
}
