package src

import (
	"errors"
	"flag"
	"os"
)

func give_error(err error) (*os.File, string, *os.File, error) {
	return nil, "", nil, err
}

func ParseArgs() (*os.File, string, *os.File, error) {
	var wordlist_dir string
	flag.StringVar(&wordlist_dir, "wordlist_dir", "", "directory to generate dictionary from")
	flag.StringVar(&wordlist_dir, "w", "", "directory to generate dictionary from (shorthand)")
	var dictionary string
	flag.StringVar(&dictionary, "dictionary", "", "dictionary (csv) to read")
	flag.StringVar(&dictionary, "d", "", "dictionary (csv) to read (shorthand)")

	var hash string
	flag.StringVar(&hash, "hash", "", "hash value to get plain text of")
	flag.StringVar(&hash, "h", "", "hash value to get plain text of (shorthand)")

	flag.Parse()

	var generate_dict bool = len(wordlist_dir) != 0

	if len(hash) == 0 && !generate_dict {
		return give_error(errors.New("You gave this program nothing to do"))
	}
	if generate_dict && len(wordlist_dir) == 0 {
		return give_error(errors.New("If generating a dictionary, must give path to wordlists"))
	}
	if len(dictionary) == 0 && !generate_dict {
		return give_error(errors.New("Need to either give a dictionary or a directory of wordlists to turn into a dictionary"))
	}

	wordlist_dir, err := ExpandTilde(wordlist_dir)
	if err != nil {
		return give_error(err)
	}

	var dir *os.File
	if len(wordlist_dir) != 0 {
		dir, err = os.Open(wordlist_dir)
		if err != nil {
			return give_error(err)
		}
	}
	var dict *os.File
	if len(dictionary) != 0 {
		dict, err = os.Open(dictionary)
		if err != nil {
			return give_error(err)
		}
	}
	return dir, hash, dict, nil
}
