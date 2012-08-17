// wlparser
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"kpa/wlparser"
	"os"
	"strings"
)

const (
	DEFAUT_PROJECT = ""
)

var (
	projName string
	ptype    string
	pw       bool
)

// Declare flags and parse argument
func init() {

	flag.StringVar(&projName, "p", DEFAUT_PROJECT, "Project name")
	flag.StringVar(&ptype, "t", "export", "input type [export, import]")
	flag.BoolVar(&pw, "w", false, "Group weeks")
	flag.Parse()
	if flag.NArg() == 0 {
		usage()
		os.Exit(0)
	}
}

//
// Main function
//
func main() {
	var (
		err    error
		parser wlparser.LineParser
	)
	if parser, err = wlparser.NewParser(ptype, projName); err != nil {
		printError(err)
	}

	lines := readLines()
	parseRes := wlparser.Parse(parser, lines)
	if pw {
		parseRes.PrintWeeks()
	} else {
		parseRes.Print()
	}

}

//
// Print flags usage 
//
func usage() {
	fmt.Println("Avaliable parameters: [OPT] filename")
	fmt.Println("OPT:")
	flag.PrintDefaults()
}

//
// Print error and exit program
//
func printError(err error) {
	fmt.Println("Error!:")
	fmt.Println(err)
	os.Exit(1)
}

//
// Read lines from file passed as argument
//
func readLines() []string {
	var (
		filecontent []byte
		err         error
	)
	filename := flag.Arg(0)
	if filecontent, err = ioutil.ReadFile(filename); err != nil {
		printError(err)
	}
	content := string(filecontent)
	return strings.Split(content, "\n")
}
