package lib

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// StrStdin - Standard input as string
func StrStdin() (stringReturned string) {
	// input from cmd prompt.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput := scanner.Text()
	stringReturned = strings.TrimSpace(stringInput)
	return
}

// IntStdin - Standard input as int
func IntStdin() (int, error) {
	// str to int
	stringInput := StrStdin()
	return strconv.Atoi(strings.TrimSpace(stringInput))
}

// ChooseNom - standard input & return num
func ChooseNom(def int) int {
	nom, err := IntStdin()
	if err != nil {
		return def
	}
	return nom
}
