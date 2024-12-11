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

// this is correct
// in go recursion is noticably less efficient than iteration
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

// MAKE A FUCKING MAP FUNCTION YOU INBRED CUNT
