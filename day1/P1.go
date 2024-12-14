package day1

import (
	f "advent-24/functions"
	"fmt"
	"slices"
)

func P1() {
	lineSplit := f.FileAsLines("day1/P1.txt")
	seperated := f.SeperateListsN(lineSplit, "   ", 2)
	intConverted := f.ListDStoDI(seperated)
	sorted := f.Map(intConverted, f.MakeReturn[[]int](slices.Sort))
	zipped := f.Zip(sorted[0], sorted[1], f.GetDifference)
	reduced := f.Reduce(zipped, f.Add)
	fmt.Println("", reduced)
}
