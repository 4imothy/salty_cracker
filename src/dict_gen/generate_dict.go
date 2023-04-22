package dict_gen

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/csv"
	"fmt"
	"os"
)

// create a CSV of plain texts and values of hash functions
func GenerateDictionary(dir *os.File, cracking bool) (*os.File, error) {
	// Create a CSV file
	print("here")
	dict, err := os.Create("dictionary.csv")
	if err != nil {
		return nil, err
	}
	if !cracking {
		defer dict.Close()
	}

	dict_writer := csv.NewWriter(dict)
	defer dict_writer.Flush()

	// Write header row
	header := []string{"Plaintext"}
	for _, h := range []string{"MD5", "SHA1", "SHA256"} {
		header = append(header, h)
	}
	dict_writer.Write(header)
	seen_plaintexts := make(map[string]bool)
	files, err := dir.ReadDir(-1)
	if err != nil {
		return nil, err
	}

	i := 0
	total_hashes := 0
	for _, file := range files {
		f, err := os.Open(dir.Name() + "/" + file.Name())
		if err != nil {
			return nil, err
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		j := 0
		for scanner.Scan() {
			plaintext := scanner.Text()
			if seen_plaintexts[plaintext] {
				continue
			}
			seen_plaintexts[plaintext] = true
			md5sum := fmt.Sprintf("%x", md5.Sum([]byte(plaintext)))
			sha1sum := fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
			sha256sum := fmt.Sprintf("%x", sha256.Sum256([]byte(plaintext)))

			// Write row to CSV
			row := []string{plaintext, md5sum, sha1sum, sha256sum}
			dict_writer.Write(row)
			j++
			if j%1000 == 0 {
				fmt.Printf("File: %s, Line: %v\n", f.Name(), j)
			}
		}
		total_hashes += j
		i++
	}
	fmt.Printf("Total Hashes: %v\n", total_hashes)

	// default
	return dict, nil
}
