package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func Reduce[T, U any](list []T, f func(T, U) U) U {
	switch len(list) {
	case 0:
		panic("Too few elements to reduce")
	case 1:
		var zero U
		return f(list[0], zero)
	default:
		return f(list[0], Reduce(list[1:], f))
	}
}

func Map[T, U any](list []T, f func(T) U) []U {
	switch len(list) {
	case 0:
		return []U{}
	case 1:
		return []U{f(list[0])}
	default:
		return append([]U{f(list[0])}, Map(list[1:], f)...)
	}
}

func Zip[T, U, V any](list1 []T, list2 []U, f func(T, U) V) []V {
	if len(list1) != len(list2) {
		panic("list lengths don't match")
	}
	switch len(list2) {
	case 1:
		return []V{f(list1[0], list2[0])}
	default:
		return append([]V{f(list1[0], list2[0])}, Zip(list1[1:], list2[1:], f)...)
	}
}

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
