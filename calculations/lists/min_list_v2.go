package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
func MinList2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Merke das erste Element als vorläufiges Minimum (min)
	minpos := 0

	// Laufe durch die Liste
	for i := 0; i < len(nums); i++ {
		// - Vergleiche jedes Element mit min
		if nums[i] < nums[minpos] {
			//   - Wenn das Element kleiner ist als min, update min
			minpos = i
		}
	}

	// Am Ende steht in min das kleines Element, dieses liefern wir als Ergebnis.
	return nums[minpos]
}
