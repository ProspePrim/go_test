package city

import (
	. "gopackages/wordz"
)

var RandomCity = []string{"Moscow", "Almaty", "Bishkek", "Tashkent", "Milan", "Paris"}

func City() string {
	Words = RandomCity
	return Random()
}
