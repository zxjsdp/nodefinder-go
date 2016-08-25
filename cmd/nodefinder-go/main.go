package main

import (
	"fmt"

	"github.com/zxjsdp/nodefinder-go/nodefindergo"
	"github.com/zxjsdp/nodefinder-go/utils"
	"flag"
	"log"
	"strings"
)

var (
	print = fmt.Println
)

func main() {
	argInputPtr := flag.String("input", "", "Input Newick tree file name")
	argConfigPtr := flag.String("config", "", "NodeFinder config file name")
	argOutputPrt := flag.String("output", "", "Output Newick tree file name")

	flag.Parse()

	utils.CheckFileExists(*argInputPtr, "-input", nodefindergo.USAGE)
	utils.CheckFileExists(*argConfigPtr, "-config", nodefindergo.USAGE)
	if len(strings.TrimSpace(*argOutputPrt)) == 0 {
		log.Fatal(fmt.Sprintf("ERROR! Blank argument for [ -output ].%s", nodefindergo.USAGE))
	}

	if len(flag.Args()) != 0 {
		log.Fatal("Invalid command line option number! " +
			nodefindergo.USAGE)
	}

	rawTreeStr := utils.ReadContent(*argInputPtr)
	calibrations := nodefindergo.ParseConfig(*argConfigPtr)

	output := nodefindergo.MultipleCalibration(rawTreeStr, calibrations)

	utils.WriteContent(*argOutputPrt, output)
}
