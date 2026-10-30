package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawRectangle(height, width int, inner, outer string) {
	for row := range height {
		for col := range width {
			fmt.Print(selectCharacter(height, width, row, col, inner, outer))
		}
		fmt.Println()
	}
}

// Wählt anhand der Höhe, Breite und akuteller Position sowie
// der vorgegebenen Zeichen aus, welches Zeichen gerade gebraucht wird.
func selectCharacter(height, width, row, col int, inner, outer string) string {
	if amRand(height, width, row, col) {
		return outer
	}
	return inner
}

// Prüft, ob wir gerade am Rand sind.
//
// Wir sind am Rand, wenn die aktuelle Zeile 0 oder die Höhe-1 ist,
// oder wenn die aktuelle Spalte 0 oder die Breite-1 ist.
func amRand(height, width, row, col int) bool {
	return row == 0 || row == height-1 ||
		col == 0 || col == width-1
}
