package p

import "log"

const benignPackageConstString = "Not an package-level address"
const validPackageConstString = "0x1b175474e89094c44da98b954eedeac495271d0f"
const invalidPackageConstString = "0x1B175474e89094c44da98b954eedeac495271d0f" // want "string contains address with capital letters"

func printingAllTypesOfStrings(format string, args ...interface{}) {
	// Constants
	const benignConstString = "Not an address"
	const validConstString = "0x2b175474e89094c44da98b954eedeac495271d0f"
	const invalidConstString = "0x2B175474e89094c44da98b954eedeac495271d0f" // want "string contains address with capital letters"

	// Vars
	var benignVarString string
	benignVarString = "Again, not an address"
	log.Printf(benignVarString)

	var validVarString string
	validVarString = "0x3b175474e89094c44da98b954eedeac495271d0f"
	log.Printf(validVarString)

	var invalidVarString string
	invalidVarString = "0x3B175474e89094c44da98b954eedeac495271d0f" // want "string contains address with capital letters"
	log.Printf(invalidVarString)

	// In-line double quotes
	benignInlineString := "Yet again, not an address"
	log.Printf(benignInlineString)

	validInlineString := "0x4b175474e89094c44da98b954eedeac495271d0f"
	log.Printf(validInlineString)

	invalidInlineString := "0x4b175474E89094c44da98b954eedeac495271d0f" // want "string contains address with capital letters"
	log.Printf(invalidInlineString)

	// In-line backticks
	benignInlineBacktickString := `One final time: not an address`
	log.Printf(benignInlineBacktickString)

	validInlineBacktickString := `0x5b175474e89094c44da98b954eedeac495271d0f`
	log.Printf(validInlineBacktickString)

	invalidInlineBacktickString := `0x5b175474E89094c44da98b954eedeac495271d0f` // want "string contains address with capital letters"
	log.Printf(invalidInlineBacktickString)

	// Large multi-address strings
	// Normally, all of these would normally be nicely formatted, but limitations in the linter require the "want" clause to be appended to the first line of the variable assignment
	validNestedInlineBacktickString := `{"addresses":{"0x4911f80530595fbd01ab1516ab61255d75aeb066":{"addresses":["0x7b175474e89094c44da98b954eedeac495271d0f","0xb0b86991c6218b36c1d19d4a2e9eb0ce3606eb48","0xdac17f958d2ee523a2206206994597c13d831ec7"]}}}`
	log.Printf(validNestedInlineBacktickString)

	invalidNestedInlineBacktickString := `{"addresses":{"0x4911f80530595fBd01ab1516ab61255d75aeb066":{"addresses":["0x7b175474e89094c44da98b954eedeac495271d0f","0xb0b86991c6218b36c1d19d4a2e9eb0ce3606eb48","0xdac17f958d2ee523a2206206994597c13d831ec7"]}}}` // want "string contains address with capital letters"
	log.Printf(invalidNestedInlineBacktickString)
}