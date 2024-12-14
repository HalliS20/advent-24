package functions

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

func ReduceArgs[T, U, V any](f func(T, U) V, variable U) func(T) V {
	newFunc := func(t T) V {
		return f(t, variable)
	}
	return newFunc
}

func ReduceArgsVoid[T, U any](f func(T, U), variable U) func(T) {
	newFunc := func(t T) {
		f(t, variable)
	}
	return newFunc
}

func SwapArgsVoid[T, U any](f func(T, U)) func(U, T) {
	newFunc := func(u U, t T) {
		f(t, u)
	}
	return newFunc
}

func SwapArgs[T, U, V any](f func(T, U) V) func(U, T) V {
	newFunc := func(u U, t T) V {
		return f(t, u)
	}
	return newFunc
}

func PanicErrors[T, U any](f func(T) (U, error)) func(T) U {
	function := func(t T) U {
		result, err := f(t)
		if err != nil {
			panic(err)
		}
		return result
	}
	return function
}

func MakeReturn[T any](f func(T)) func(T) T {
	function := func(t T) T {
		f(t)
		return t
	}
	return function
}
