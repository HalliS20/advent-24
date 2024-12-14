package functions

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
