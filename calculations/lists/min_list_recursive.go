package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
// ZUSATZBEDINGUNG: Diese Funktion darf keine Schleife verwenden.
func MinListRecursive(nums []int) int {

	// Basisfälle: nums ist leer oder Länge 1
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	// Bestimme das kleinste Element aus der restlichen Liste
	rest := nums[1:] // Die Liste ab Stelle 1
	minRest := MinListRecursive(rest)

	// Vergleiche das erste Element mit diesem kleinsten Element
	// aus der restlichen Liste
	if nums[0] < minRest {
		return nums[0]
	}
	return minRest
}

// REMARKS
// - Diese Funktion nennt man "rekursiv".
// - Rekursion ist ein größeres Thema, das in der Vorlesung separat behandelt wird.
//   Um die Denkweise frühzeitig zu üben, gibt es aber gelegentlich auch vorab Aufgaben dazu.
