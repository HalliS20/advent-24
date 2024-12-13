package main

import (
	f "advent-24/functions"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func seperateListsN(list []string, seperator string, n int) [][]string {
	lists := [][]string{}
	for i := 0; i < n; i++ {
		newlist := []string{}
		lists = append(lists, newlist)
	}

	for i := 0; i < len(list); i++ {
		splitted := strings.Split(list[i], seperator)
		if len(splitted) != n {
			fmt.Printf("ERROR: Seperation Count Dont match N\n")
			panic("list location: " + strconv.Itoa(i))
		}
		for j := 0; j < n; j++ {
			lists[j] = append(lists[j], splitted[j])
		}
	}
	return lists
}

func listDStoDI(lists [][]string) [][]int {
	converted := [][]int{}
	for _, innerList := range lists {
		newList := f.Map(innerList, f.PanicErrors(strconv.Atoi))
		converted = append(converted, newList)
	}
	return converted
}

func day1() {
	fileString := f.FileAsString("day1-1.txt")
	lineSplit := strings.Split(fileString, "\n")
	lineSplit = lineSplit[:len(lineSplit)-1]
	seperated := seperateListsN(lineSplit, "   ", 2)
	intConverted := listDStoDI(seperated)
	sorted := f.Map(intConverted, f.MakeReturn[[]int](slices.Sort))
	zipped := f.Zip(sorted[0], sorted[1], f.GetDifference)
	reduced := f.Reduce(zipped, f.Add)
	fmt.Println("added differences", reduced)
}
