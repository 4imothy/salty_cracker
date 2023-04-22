# Salty Cracker
### A Rainbow Table to Get the Plaintext of Hashes
#### For CWRU Hacker Society 2023

Password breaker using go.

#### To Use
**Install**

[1]. Clone the code and run *go build -o=sc*, or (macos) download the sc.gz in releases, then use *gunzip* to decompress and *chmod +x* to mark it executable
Run with *-h* to get more info about flags.

**Create A Wordlist**

[1]. Compile lists of passwords and place them in a directory

[2]. Run *./sc -w="path to directory"* this will create a *dictionary.csv* file with the plain texts and there hash values

**To Get the Plaintext of A Hash**

If a matching hash exists in the dictionary it will be found

[1]. Create your own wordlist with the directions above or by copying the wordlists folder in the repo and running *./sc -w="path to directory"* this will create a *dictionary.csv*

[2]. Run *./sc* with *-h* to describe the hash value and *-d* to tell the executable which dictionary to check.
