package functions

import (
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
