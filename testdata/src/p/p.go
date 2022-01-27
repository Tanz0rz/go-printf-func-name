package p

import "log"

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
	validNestedInlineBacktickString := `{"lastFetchedPoolIndex":7,"pools":{"0x3911f80530595fbd01ab1516ab61255d75aeb066":{"coins":["0x6b175474e89094c44da98b954eedeac495271d0f","0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48","0xdac17f958d2ee523a2206206994597c13d831ec7"],"PRECISION_MUL":["1","1000000000000","1000000000000","1000000000000000000","1000000000000000000","1000000000000000000","1000000000000000000","1000000000000000000"],"IsCustomPool":true,"IsMetaPool":false},"0x4f6a43ad7cba042606decaca730d4ce0a57ac62e":{"coins":["0x8daebade922df735c38c80c7ebd708af50815faa","0x2260fac5e5542a773aa44fbcfedf7c193bc2c599","0xeb4c2781e4eba804ce9a9803c67d0893436bb27d","0xfe18be6b3bd88a2d2a7f928d00292e7a9963cfc6"],"PRECISION_MUL":["1","10000000000","10000000000","1","1000000000000000000","1000000000000000000","1000000000000000000","1000000000000000000"],"IsCustomPool":true,"IsMetaPool":false},"0xdec2157831D6ABC3Ec328291119cc91B337272b5":{"coins":["0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2","0x898bad2774eb97cf6b94605677f43b41871410b1"],"PRECISION_MUL":["1","1","1000000000000000000","1000000000000000000","1000000000000000000","1000000000000000000","1000000000000000000","1000000000000000000"],"IsCustomPool":true,"IsMetaPool":false},"0xa6018520eaacc06c30ff2e1b3ee2c7c22e64196a":{"coins":["0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2","0x0100546f2cd4c9d97f798ffc9755e47865ff7ee6","0x5e74c9036fb86bd7ecdcb084a0673efc32ea31cb"],"PRECISION_MUL":["1","1","1","1000000000000000000","1000000000000000000","1000000000000000000","1000000000000000000","1000000000000000000"],"IsCustomPool":true,"IsMetaPool":false},"0xc69ddcd4dfef25d8a793241834d4cc4b3668ead6":{"coins":["0xbc6da0fe9ad5f3b0d58160288917aa56653660e9","0x956f47f50a910163d8bf957cf5846d573e7f87ca","0x853d955acef822db058eb8505911ed77f175b99e","0x5f98805a4e8be255a32880fdec7f6728c6568ba0"],"PRECISION_MUL":["1","1","1","1","1000000000000000000","1000000000000000000","1000000000000000000","1000000000000000000"],"IsCustomPool":true,"IsMetaPool":false}}}`
	log.Printf(validNestedInlineBacktickString)

	invalidNestedInlineBacktickString := `{"0x930d001b7efb225613aC7F35911c52Ac9E111Fa9":{"coins":["0x23b891e5C62E0955ae2bD185990103928Ab817b3",//nUSD"0x6c3f90f043a72fa612cbac8115ee7e52bde6e490"//USD-LP],"underlying_coins":["0x056fd409e1d7a124bd7017459dfea2f387b6d5cd","0x6b175474e89094c44da98b954eedeac495271d0f","0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48","0xdac17f958d2ee523a2206206994597c13d831ec7"],"PRECISION_MUL":["1","1000000000000","1000000000000","1000000000000000000","1000000000000000000","1000000000000000000","1000000000000000000","1000000000000000000"],"PRECISION_MUL_META":["10000000000000000","1","1000000000000000000","1000000000000000000","1000000000000000000","1000000000000000000","1000000000000000000","1000000000000000000"],"isCustomPool":true,"isMetaPool":true,"CURVE_POOL_ADDRESS":"0xbebc44782c7db0a1a60cb6fe97d0b483032ff1c7","A_PRECISION":"100"}}` // want "string contains address with capital letters"
	log.Printf(invalidNestedInlineBacktickString)
}