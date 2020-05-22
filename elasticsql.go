package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cch123/elasticsql"
)

func main() {
	if len(os.Args) < 2 {
		printHelpAndExit()
	}

	args := os.Args[1:]
	var sql string
	var isPretty bool
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			printHelpAndExit()
		}
		if arg == "-p" || arg == "--pretty" {
			isPretty = true
		}
		if !strings.HasPrefix(arg, "-") {
			sql = arg
		}
	}
	var dsl string
	var err error
	if isPretty {
		dsl, _, err = elasticsql.ConvertPretty(sql)
	} else {
		dsl, _, err = elasticsql.Convert(sql)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(dsl)
	os.Exit(0)
}

func printHelpAndExit() {
	fmt.Println(
		`Usage:	
  elasticsql "SELECT * FROM my_table m WHERE m.id > 3"
  elasticsql -p "SELECT * FROM my_table m WHERE m.id > 3"
Options:	
  -p / --pretty
    prettifies the dsl query
  -h / --help
    shows this message`)
	os.Exit(0)
}
