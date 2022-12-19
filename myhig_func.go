package myhig


func From0r1[T1 any](fn func() T1) *Myhig[Tuple1[T1]] {
	return &Myhig[Tuple1[T1]]{}
}

func From0r2[T1 any, T2 any](fn func() (T1, T2)) *Myhig[Tuple2[T1, T2]] {
	return &Myhig[Tuple2[T1, T2]]{}
}
