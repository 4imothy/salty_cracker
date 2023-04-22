package cracker

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

func give_error(err error) (string, string, error) {
	return "", "", err
}

// dict is already open
func SearchDict(hash string, dict *os.File) (string, string, error) {
	// reset the file
	dict.Seek(0, io.SeekStart)
	reader := csv.NewReader(dict)
	header, err := reader.Read()
	if err != nil {
		return give_error(err)
	}

	// get the index of all columns
	plain_index := -1
	md5_index := -1
	sha1_index := -1
	sha256_index := -1

	for i, col := range header {
		switch col {
		case "Plaintext":
			plain_index = i
		case "MD5":
			md5_index = i
		case "SHA1":
			sha1_index = i
		case "SHA256":
			sha256_index = i
		}
	}

	rows, err := reader.ReadAll()
	if err != nil {
		give_error(err)
	}

	for i, row := range rows {
		if i%10000 == 0 {
			fmt.Printf("Searching Row %v of dictionary\n", i)
		}
		i++
		plain := row[plain_index]
		md5 := row[md5_index]
		sha1 := row[sha1_index]
		sha256 := row[sha256_index]
		switch hash {
		case md5:
			return plain, "md5", nil
		case sha1:
			return plain, "sha1", nil
		case sha256:
			return plain, "sha256", nil
		}
	}

	return "", "", errors.New("Didn't find a match")
}
