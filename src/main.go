package src

import (
	"fmt"
	"os"
	"salty_cracker/src/cracker"
	"salty_cracker/src/dict_gen"
)

func Run() {
	dir, hash, given_dict, err := ParseArgs()
	if err != nil {
		FatalError(err.Error())
	}

	if hash != "" {
		var dict *os.File
		if dir != nil {
			// result is an unclosed image file
			result, err := dict_gen.GenerateDictionary(dir, true)
			if err != nil {
				FatalError(err.Error())
			}
			dict = result
		} else {
			// known that this value is not null
			dict = given_dict
		}
		// get plain text
		plaintext, algo, err := cracker.SearchDict(hash, dict)
		if err != nil {
			FatalError(err.Error())
		}
		fmt.Printf("Plaintext found: %s\nHasd found used the %s algorithm \n", plaintext, algo)

	} else {
		result, err := dict_gen.GenerateDictionary(dir, false)
		if err != nil {
			FatalError(err.Error())
		}
		fmt.Printf("Created CSV Dictionary in: %s\n", result.Name())
	}
}
