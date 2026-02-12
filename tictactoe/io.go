package tictactoe

import (
	"fmt"
)

// Show zeigt den aktuellen Spielstand an.
func (g *Game) Show() {
	fmt.Println(g.board)
}

// Move fragt den Spieler nach einem Zug und führt diesen Zug aus, falls er erlaubt ist.
// Falls der Zug nicht erlaubt ist, wird der Spieler erneut nach einem Zug gefragt.
func (g *Game) Move(player string) {
	// Hinweis:
	// - Zeigen Sie zuerst das Spielfeld mit den nummerierten Zellen an, damit der Spieler weiß, welche Zahl welcher Zelle entspricht.
	// - Fragen Sie den Spieler dann nach seiner Eingabe.
	// - Für die Überprüfung der Eingabe können Sie sie als String einlesen und dann überprüfen, ob dieser
	//   String eine gültige Zahl zwischen "1" und "9" ist.
	//   Falls nicht, geben Sie eine Fehlermeldung aus und rufen Sie Move erneut auf.
	// - Wandeln Sie dann die eingegebene Zahl in die entsprechenden Zeilen- und Spaltenindizes um.
	// - Die Zeilen- und Spaltenindizes können Sie mit einfachen mathematischen Operationen berechnen.
	// - Überprüfen Sie auch diese Indizes erneut auf Gültigkeit, bevor Sie den Zug ausführen.
	//   Falls der Zug nicht erlaubt ist, geben Sie eine Fehlermeldung aus und rufen Sie Move erneut auf.

	// TODO
}
