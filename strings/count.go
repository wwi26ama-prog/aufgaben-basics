package strings

// Erwartet einen string s und zählt,
// wie oft der Buchstabe 'A' in s vorkommt.
func CountA(s string) int {
	zaehler := 0

	for _, element := range s {
		if element == 'A' {
			zaehler++
		}
	}

	return zaehler
}

// func CountA(s string) int {
// 	zaehler := 0
//
// 	for i := range s {
// 		if s[i] == 'A' {
// 			zaehler++
// 		}
// 	}
//
// 	return zaehler
// }
