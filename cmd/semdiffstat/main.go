package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

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

	max := maxCountLength(changes)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', 0)
	for _, c := range changes {
		// TODO: general UI improvements
		var annotation string
		switch {
		case c.Inserted:
			annotation = " (inserted)"
		case c.Deleted:
			annotation = " (deleted)"
		}
		fmt.Fprintf(w, "%v\t | %*d\t %s%s%s\t\n", c.Name, max, c.InsLines+c.DelLines,
			pluses(c.InsLines), minuses(c.DelLines), annotation)
	}
	fmt.Fprintln(w)
	w.Flush()
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
func maxCountLength(changes []*semdiffstat.Change) (max int) {
	for _, c := range changes {
		n := len(strconv.Itoa(c.InsLines + c.DelLines))
		if n > max {
			max = n
		}
	}
	return
}
