package main

import (
	f "advent-24/functions"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func fileAsString(fileName string) string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %s", fileName)
		panic(err)
	}
	return string(data)
}

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
			fmt.Printf("ERROR: N: %d\n", n)
			fmt.Printf("ERROR: Seperations : %d\n", len(splitted))
			panic("list location: " + strconv.Itoa(i))
		}
		for j := 0; j < n; j++ {
			lists[j] = append(lists[j], splitted[j])
		}
	}
	return lists
}

// WE WILL FIX THIS WITH MAP DONT YOU WORRY CHILE
func listDStoDI(lists [][]string) [][]int {
	converted := [][]int{}
	for _, innerList := range lists {
		newList := []int{}
		for i := range innerList {
			intValue, err := strconv.Atoi(innerList[i])
			if err != nil {
				fmt.Println("I to S Failed")
				panic(err)
			}
			newList = append(newList, intValue)
		}
		converted = append(converted, newList)
	}
	return converted
}

func sortDList[Type cmp.Ordered](list [][]Type) [][]Type {
	for i, innerList := range list {
		slices.Sort(innerList)
		list[i] = innerList
	}
	return list
}

// TODO: get difference properly
// by subtracting the smaller number from the 2 lists from the larger
func day1() {
	fileString := fileAsString("day1.txt")
	lineSplit := strings.Split(fileString, "\n")
	lineSplit = lineSplit[:len(lineSplit)-1]
	seperated := seperateListsN(lineSplit, "   ", 2)
	fmt.Println("seperated length:", len(seperated))
	intConverted := listDStoDI(seperated)

	sorted := sortDList(intConverted)
	sorted = sorted
	// This approach doesn't work
	negative := f.Reduce(sorted[0], f.Sub)
	positive := f.Reduce(sorted[1], f.Add)
	fmt.Println("negative:", negative)
	fmt.Println("positive:", positive)
	fmt.Println("result:", negative+positive)
}
