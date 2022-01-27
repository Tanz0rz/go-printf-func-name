package analyzer

import (
	"regexp"
)

var capitalLetterPattern = regexp.MustCompile("[A-F]")
var addressPattern = regexp.MustCompile("0x[a-fA-F0-9]{40}")

func validateString(valueRaw string) bool {

	// Looks for all addresses in the string
	addresses := addressPattern.FindAllString(valueRaw, -1)

	// If no addresses, return early
	if len(addresses) == 0 {
		return true
	}

	// For each address, look for any capital letters
	for _, address := range addresses {
		found := capitalLetterPattern.FindString(address)
		if len(found) == 0 {
			return true
		}

		return false
	}

	return true
}
