package day1

import (
	f "advent-24/functions"
	"fmt"
)

func P2() {
	// Split lines numbers
	lineList := f.FileAsLines("day1/P2.txt")
	splitLists := f.SeperateListsN(lineList, "   ", 2)
	inted := f.ListDStoDI(splitLists)
	numMap := make(map[int]int) // our map for counting

	// create mappable function for setting
	swappedSet := f.SwapArgsVoid[map[int]int, int](f.SetMapI[int, int])
	reducedSet := f.ReduceArgsVoid[int, map[int]int](swappedSet, numMap)
	returningSet := f.MakeReturn(reducedSet)

	// set all keys to 0
	f.Map(inted[0], returningSet)

	// create mappable function for increasing
	swappedInc := f.SwapArgsVoid[map[int]int, int](f.IncMapValByKey)
	reducedInc := f.ReduceArgsVoid(swappedInc, numMap)
	returningInc := f.MakeReturn(reducedInc)

	// set key values to the number multiplied by number of appearances
	f.Map(inted[1], returningInc)

	// change Map values to list and sum
	values := f.GetMapItems[int, int, int](numMap, f.VALUE)
	reduced := f.Reduce(values, f.Add)
	fmt.Println("", reduced)
}
