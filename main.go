package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"net/http"
	"time"
	"strings"
)

func logo() { // This func is for printing the tool logo
	fmt.Println("\n\t\t  /$$$$$$            /$$        /$$$$$$                                        ")
	fmt.Println("\t\t /$$__  $$          | $$       /$$__  $$                                       ")
	fmt.Println("\t\t| $$  \\__/ /$$   /$$| $$$$$$$ | $$  \\__/ /$$   /$$  /$$$$$$   /$$$$$$$ /$$$$$$$")
	fmt.Println("\t\t|  $$$$$$ | $$  | $$| $$__  $$| $$ /$$$$| $$  | $$ /$$__  $$ /$$_____//$$_____/")
	fmt.Println("\t\t \\____  $$| $$  | $$| $$  \\ $$| $$|_  $$| $$  | $$| $$$$$$$$|  $$$$$$|  $$$$$$ ")
	fmt.Println("\t\t /$$  \\ $$| $$  | $$| $$  | $$| $$  \\ $$| $$  | $$| $$_____/ \\____  $$\\____  $$")
	fmt.Println("\t\t|  $$$$$$/|  $$$$$$/| $$$$$$$/|  $$$$$$/|  $$$$$$/|  $$$$$$$ /$$$$$$$//$$$$$$$/")
	fmt.Println("\t\t \\______/  \\______/ |_______/  \\______/  \\______/  \\_______/|_______/|_______/ ")
	fmt.Println("\t\t                SubGuess Tool By @SirBugs")
	fmt.Println("\t\t             V: 1.0.1 Made With All Love <3")
	fmt.Println("\t\t       For Generating/Checking Sub-Sub-Sub Domains")
	fmt.Println("\t\t           Twitter@SirBagoza / GitHub@SirBugs")
	fmt.Println("\t\t   Run: go run main.go -l domains.txt -w wordlist.txt\n")
}

var TheWholeCycle []string; // Storing All The Subdomain, Sub-SubDomains, Sub-Sub-SubDomains
var CreatedSubDomains []string; // Storing SubDomains
var CreatedSubDomains2 []string; // Storing Sub-SubDomains
var CreatedSubDomains3 []string; // Storing Sub-Sub-SubDomains

func submaker(Value string, fName string, wordList string) { // This func is for making the first SubDomains
	if Value == "domain" {
		fileRead, err := os.Open(wordList) // word.txt is alr downloaded with the tool
		if err != nil {
			log.Fatal(err)
		}
		myScannerr := bufio.NewScanner(fileRead) // Reading it
		for myScannerr.Scan() {
			subDomain := myScannerr.Text()
			CreatedSubDomains = append(CreatedSubDomains, subDomain+"."+fName)
		}
	} else {
		myFile, err := os.Open(fName) // Opening this target file
		if err != nil {
			log.Fatal(err)
		}
		myScanner := bufio.NewScanner(myFile) // Reading it
		for myScanner.Scan() {
			domainName := myScanner.Text()
			fmt.Println("\t"+domainName)
			fileRead, err := os.Open(wordList) // word.txt is alr downloaded with the tool
			// If u wanna submit your own list, just send it with param "-w : --wordlist"
			if err != nil {
				log.Fatal(err)
			}
			myScannerr := bufio.NewScanner(fileRead) // Reading it
			for myScannerr.Scan() {
				subDomain := myScannerr.Text()
				// fmt.Println(subDomain+"."+domainName)
				CreatedSubDomains = append(CreatedSubDomains, subDomain+"."+domainName) // appending all made subdomains to slice called "CreatedSubDomains"
			}
			// fmt.Println(CreatedSubDomains)
		}
	}
}

func cycle(wordlist string) { // This func is for repeating the main cycle 1 more time! Sub-SubDomains
	theFile, err := os.Open(wordlist)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(theFile)
	for fileScanner.Scan() {
		NewSub := fileScanner.Text()
		for _, elm := range CreatedSubDomains {
			CreatedSubDomains2 = append(CreatedSubDomains2, NewSub+"."+elm)
		}
	}
}

func cycle2(wordlist string) { // This func is for repeating the main cycle 1 more time! Sub-Sub-SubDomains
	theFile, err := os.Open(wordlist)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(theFile)
	for fileScanner.Scan() {
		NewSub := fileScanner.Text()
		for _, elm := range CreatedSubDomains2 {
			CreatedSubDomains3 = append(CreatedSubDomains3, NewSub+"."+elm) // Adding all of them to slice called CreatedSubDomains3
		}
	}
}

func AddToTheWholeSubs() { // This func is for adding all the SubDomain, Sub-SubDomains, Sub-Sub-SubDomains, To a slice called TheWholeCycle
	for _, elm := range CreatedSubDomains {
		TheWholeCycle = append(TheWholeCycle, elm) // Adding the SubDomains to slice called TheWholeCycle
	}
	for _, elm := range CreatedSubDomains2 {
		TheWholeCycle = append(TheWholeCycle, elm) // Adding the Sub-SubDomains to slice called TheWholeCycle
	}
	for _, elm := range CreatedSubDomains3 {
		TheWholeCycle = append(TheWholeCycle, elm) // Adding the Sub-Sub-SubDomains to slice called TheWholeCycle
	}
	// fmt.Println(TheWholeCycle)
}

func req(subDomain string) { // This func is for requesting the generated subdomains

	client := http.Client { // Creating the new client
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("GET", "https://"+subDomain, nil) // requesting the URL
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/109.0") // Submitting real User-Agent
	resp, err := client.Do(req) // Doing the request
	if resp == nil {
		fmt.Printf("") // [ - ] Request didn't go through [ "+subDomain+" ]
	} else {
		fmt.Println("[ $ ] Subdomain Found: [ "+subDomain+" ]") // If the Url Is Valid and Subdomain or Sub-Sub or Sub-Sub-Sub are fine and available
	}

}

func main() { // The main func !! <3

	// Identifing the arguments
	MyArgsTxt := strings.Join(os.Args[1:], " ") // Joining All the args input together
	// fmt.Println(MyArgsTxt)

	domain := ""
	list := ""
	wordlist := ""
	value := "list"
	target := ""
	fmt.Println(domain, list) // Just for ignoring {err:"declared but not used"}

	if strings.Contains(MyArgsTxt, "-d") { X := strings.Split(strings.Split(MyArgsTxt,"-d ")[1]," "); domain = X[0] } // Arg -d for domain
	if strings.Contains(MyArgsTxt, "-l") { X := strings.Split(strings.Split(MyArgsTxt,"-l ")[1]," "); list = X[0] } // Arg -l for targets list
	if strings.Contains(MyArgsTxt, "-w") { X := strings.Split(strings.Split(MyArgsTxt,"-w ")[1]," "); wordlist = X[0] } else { wordlist = "word.txt"} // Arg -w for custom wordlist

	if strings.Contains(MyArgsTxt, "-d") && strings.Contains(MyArgsTxt, "-l") { // Checking if the user is using -l and -d together, timeout!
		fmt.Println("[ ! ] Error: You Cannot Use -d : domain , -l : list, Just use one of them!") // output
		os.Exit(1) // timeout
	}
	if !strings.Contains(MyArgsTxt, "-d") && !strings.Contains(MyArgsTxt, "-l") && !strings.Contains(MyArgsTxt, "-w") { // Checking if the user didn't use any arg -d,-l,-w then timeout!
		fmt.Println("[ ! ] Error: Please run: go run main.go -d example.com / -l targets.txt [OPTION] -w word.txt") // output
		os.Exit(1) // timeout
	}

	logo() // Printing the logo func (line:13)

	fmt.Printf("[ ! ] Started Working !!\n")
	fmt.Printf("[ ! ] Wordlist: "+wordlist+" !!\n")
	fmt.Printf("[ ! ] Targets:\n") // If you are using -d it's gonna be empty, If you are using -l it's gonna show all the domains/subdomains

	if strings.Contains(MyArgsTxt, "-d") { value = "domain"; target = domain} else { value = "list"; target = list}

	submaker(value, target, wordlist) // Calling submaker (line:34)
	cycle(wordlist) // Calling cycle (line:70)
	cycle2(wordlist) // Calling cycle2 (line:84)
	AddToTheWholeSubs() // Calling submaker (line:98)

	fmt.Println("") // Empty line <3

	for _, element := range TheWholeCycle {
		go req(element)
		time.Sleep(50 * time.Millisecond)
	}
	time.Sleep(7 * time.Second) // Sleeping 7 Seconds till all the goroutines are done!
}
