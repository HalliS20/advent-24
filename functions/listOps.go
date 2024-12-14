package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func SplitList(list []string, seperator string) [][]string {
	newList := [][]string{}
	for i := 0; i < len(list); i++ {
		splitted := strings.Split(list[i], seperator)
		newList = append(newList, splitted)
	}
	return newList
}

func longerList[T any](list1 []T, list2 []T) []T {
	if len(list1) > len(list2) {
		return list1
	}
	return list2
}

func LongerListFirst[T any](list1 *[]T, list2 *[]T) {
	temp := *list1
	*list1 = *list2
	*list2 = temp
}

func SeperateListsN(list []string, seperator string, n int) [][]string {
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

func ListDStoDI(lists [][]string) [][]int {
	converted := [][]int{}
	for _, innerList := range lists {
		newList := Map(innerList, PanicErrors(strconv.Atoi))
		converted = append(converted, newList)
	}
	return converted
}
