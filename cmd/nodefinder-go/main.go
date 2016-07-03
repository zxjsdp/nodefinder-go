package main

import (
	"fmt"

	"github.com/zxjsdp/nodefinder-go/nodefindergo"
)

var (
	print = fmt.Println
)

func main() {
	print(nodefindergo.GetIndexOfTMRCA("((a,((b,c),(ddd,e))),(f,g));", "b", "ddd"))
}
