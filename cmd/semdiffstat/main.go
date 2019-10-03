package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/whatsnu/semdiffstat"
)

func main() {
	flag.Usage = func() {
		fmt.Println(os.Args)
		fmt.Fprintln(os.Stderr, "semdiffstat a.go b.go")
		os.Exit(2)
	}

	// Accept two args (manual mode) or seven args (git external diff tool mode).
	flag.Parse()
	if flag.NArg() != 2 && flag.NArg() != 7 {
		flag.Usage()
	}
	var aName, bName string
	if flag.NArg() == 2 {
		aName, bName = flag.Arg(0), flag.Arg(1)
	} else {
		// We are a git external diff tool.
		// Print the filename first.
		color.NoColor = false
		hiYellow.Println(flag.Arg(0))
		aName, bName = flag.Arg(1), flag.Arg(4)
	}
	asrc, err := ioutil.ReadFile(aName)
	if err != nil {
		log.Fatal("could not read file %v: %v", aName, err)
	}
	bsrc, err := ioutil.ReadFile(bName)
	if err != nil {
		log.Fatal("could not read file %v: %v", bName, err)
	}

	changes, err := semdiffstat.Go(asrc, bsrc)
	if err != nil {
		// TODO: emit a regular diffstat instead?
		log.Fatalf("could not parse Go files %v and %v: %v\n", aName, bName, err)
	}

	for _, c := range changes {
		// TODO: align | characters
		// TODO: general UI improvements
		if c.Inserted {
			fmt.Printf("%v (inserted) %s\n", c.Name, pluses(c.InsLines))
			continue
		}
		if c.Deleted {
			fmt.Printf("%v (deleted) %s\n", c.Name, minuses(c.DelLines))
			continue
		}
		// Modified/other
		fmt.Printf("%v | %d %s%s\n", c.Name, c.InsLines+c.DelLines, pluses(c.InsLines), minuses(c.DelLines))
	}

	fmt.Println()
}

var (
	hiYellow  = color.New(color.FgHiYellow)
	boldGreen = color.New(color.FgGreen, color.Bold)
	boldRed   = color.New(color.FgRed, color.Bold)
)

func pluses(ins int) string {
	return boldGreen.Sprint(strings.Repeat("+", ins))
}
func minuses(del int) string {
	return boldRed.Sprint(strings.Repeat("-", del))
}
