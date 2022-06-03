package digit

import (
	. "gopackages/wordz"
)

var RandomNumbers = []string{"one", "two", "three", "four", "five"}

func Digit() string {
	Words = RandomNumbers
	return Random()
}
