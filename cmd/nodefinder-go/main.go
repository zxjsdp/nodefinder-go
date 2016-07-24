package main

import (
	"fmt"

	"github.com/zxjsdp/nodefinder-go/nodefindergo"
	"github.com/zxjsdp/nodefinder-go/utils"
	"flag"
	"log"
)

var (
	print = fmt.Println
)

func main() {
	argInputPtr := flag.String("input", "", "Input Newick tree file name")
	argConfigPtr := flag.String("config", "", "NodeFinder config file name")
	argOutputPtr := flag.String("output", "", "Output Newick tree file name")

	flag.Parse()

	if len(utils.RemoveBlankChars(*argInputPtr)) == 0 {
		log.Fatal("Invalid input filename: " + *argInputPtr +
			nodefindergo.USAGE)
	}
	if len(utils.RemoveBlankChars(*argConfigPtr)) == 0 {
		log.Fatal("Invalid config filename: " + *argConfigPtr +
			nodefindergo.USAGE)
	}
	if len(utils.RemoveBlankChars(*argOutputPtr)) == 0 {
		log.Fatal("Invalid output filename! " + *argOutputPtr +
			nodefindergo.USAGE)
	}

	if len(flag.Args()) != 0 {
		log.Fatal("Invalid command line options! " +
			nodefindergo.USAGE)
	}

	rawTreeStr := utils.ReadContent(*argInputPtr)
	calibrations := nodefindergo.ParseConfig(*argConfigPtr)

	nodefindergo.MultipleCalibration(rawTreeStr, calibrations)

	//nodefindergo.Test()
}
