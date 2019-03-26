package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, ""+
			`agenda-creator downloads and extrasts the list of fixtures.
Then creates, in google calendar, one event by fixture.

`)
		flag.PrintDefaults()
	}
	address := flag.String("address", "", "address of official site")
	number := flag.Int("number", 22, "fixtures count")
	team := flag.String("team", "", "specific team")
	flag.Parse()

	dom, err := extractDom(*address, *number)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	fixtures, err := extractFixtures(*team, dom)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	fmt.Printf("%v", fixtures)
}
