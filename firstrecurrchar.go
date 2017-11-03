/*
Package contains solution to "Rirst Recurring Character"
coding interview problem.

The task is to find and print the first char in given string
which has more the one occurence, examples:

	"ABCA"	->	'A"
	"BCABA"	->	'B'
	"ABC"	->	''

Task source: https://www.youtube.com/watch?v=GJdiM-muYqc
*/
package firstrecurrchar

// Finds and returns first recrrent char in given string.
// For examples see package description.
func FindChar(s string) rune {
	chars := []rune(s)
	chSet := make(map[rune]int)

	for _, ch := range chars {
		_, has := chSet[ch]
		if has == false {
			chSet[ch] = 1
		} else {
			// found
			return ch
		}
	}

	// not found
	return '\x00'
}
