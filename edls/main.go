package main

import (
	"flag"
	"fmt"
)

func main() {

	//filter pattern flags

	flagPattern := flag.String("p", "", "Filter by pattern")
	flagAll := flag.Bool("a", false, "All files include hidden files")
	flagNumberRecords := flag.Int("n", 0, "Number of records")

	//order files flags

	hasOrderByTime := flag.Bool("t", false, "Sort by time, oldest to newest")
	hasOrderBySize := flag.Bool("s", false, "Sort by file size, smallest to biggest")
	hasOrderReverse := flag.Bool("r", false, "Reverse order of sorting")

	flag.Parse()

	fmt.Println("Pattern:", *flagPattern)
	fmt.Println("All:", *flagAll)
	fmt.Println("NumberRecords:", *flagNumberRecords)
	fmt.Println("OrderByTime:", *hasOrderByTime)
	fmt.Println("OrderBySize:", *hasOrderBySize)
	fmt.Println("Orderreverse", *hasOrderReverse)

}
