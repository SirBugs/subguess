# subguess
Generating Sub-Sub-Subdomain + validating all of them

# Idea
This tool is generating Subdomains, Sub-Subdomains, Sub-Sub-Subdomains and validating them, For ex:

help.example.com

timer.help.example.com

admin.timer.help.example.com

# Installation
Just need to install go, run:

```
▶ brew install go
▶ git clone https://github.com/SirBugs/subguess.git
```

or download from https://go.dev/dl/

# Usage:
```
▶ go run main.go -l domains.txt -w wordlist.txt

                  /$$$$$$            /$$        /$$$$$$
                 /$$__  $$          | $$       /$$__  $$
                | $$  \__/ /$$   /$$| $$$$$$$ | $$  \__/ /$$   /$$  /$$$$$$   /$$$$$$$ /$$$$$$$
                |  $$$$$$ | $$  | $$| $$__  $$| $$ /$$$$| $$  | $$ /$$__  $$ /$$_____//$$_____/
                 \____  $$| $$  | $$| $$  \ $$| $$|_  $$| $$  | $$| $$$$$$$$|  $$$$$$|  $$$$$$
                 /$$  \ $$| $$  | $$| $$  | $$| $$  \ $$| $$  | $$| $$_____/ \____  $$\____  $$
                |  $$$$$$/|  $$$$$$/| $$$$$$$/|  $$$$$$/|  $$$$$$/|  $$$$$$$ /$$$$$$$//$$$$$$$/
                 \______/  \______/ |_______/  \______/  \______/  \_______/|_______/|_______/
                                SubGuess Tool By @SirBugs
                             V: 1.0.1 Made With All Love <3
                       For Generating/Checking Sub-Sub-Sub Domains
                           Twitter@SirBagoza / GitHub@SirBugs
                   Run: go run main.go -l domains.txt -w wordlist.txt

[ ! ] Started Working !!
[ ! ] Wordlist: wordlist.txt !!
[ ! ] Targets:
        example.com
        example2.com

[ $ ] Subdomain Found: [ administrator.example.com ]
[ $ ] Subdomain Found: [ freight.example2.com ]
[ $ ] Subdomain Found: [ admin.eu-east-1.example.com ]
[ $ ] Subdomain Found: [ promo.example2.com ]

```
- -d --domain
- -l --list
- -w --wordlist
Note: You cannot use -d and -l together! Just one of them <3

Note: You can use your own list by submitting -w and if u didn't the tool is automatically using word.txt

# Credits
This tool was written in Golang 1.19.4, Made with all love in Egypt! <3
Twitter@SirBagoza , Github@SirBugs
