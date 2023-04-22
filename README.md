# Salty Cracker
### For CWRU Hacker Society 2023

Password breaker using go.

#### To Use
Run with *-h* to get more info about flags.
**Create A Wordlist**
[1]. Clone the code and run *go build sc.go*
[2]. Compile lists of passwords and place them in a directory
[3]. Run *./sc -w="path to directory"* this will create a *dictionary.csv* file with the plain texts and there hash values

**To Get the Plaintext of A Hash**
If a matching hash exists in the dictionary it will be found
[1]. Clone the code and run *go build sc.go*
[2]. Create your own wordlist with directions above or 




**Crack a Password**

**To Get The WordLists:**
- curl https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/Common-Credentials/10-million-password-list-top-1000000.txt > ten_million_password.txt
- Go to this link: http://downloads.skullsecurity.org/passwords/rockyou.txt.bz2, and run bzip2 -d rockyou.txt.bz2, in the terminal
