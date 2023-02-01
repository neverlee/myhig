package myhig


func Ifvv[T any](cond bool, a, b T) T {
    if cond {
        return a
    }
    return b
}

func Ifvn[T any](cond bool, a T, bf func() T) T {
    if cond {
        return a
    }
    return bf()
}

func Ifnn[T any](cond bool, af, bf func() T) T {
    if cond {
        return af()
    }
    return bf()
}

