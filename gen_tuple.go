package myhig

type Tuple1[T1 any]struct {
	V1 T1
}

func Tw1[T1 any](v1 T1) Tuple1[T1] {
	return Tuple1[T1]{v1}
}

func (t *Tuple1[T1]) Unwrap() (T1) {
	return t.V1
}

func (t *Tuple1[T1]) Set(v1 T1) {
	t.V1 = v1
}

func (t *Tuple1[T1]) setFail(err error) {
	setError(&t.V1, err)
}

type RetProc1[T1 any]struct {
	Tuple1[T1]
	Proc
}

func NewRetProc1[T1 any]() *RetProc1[T1] {
	return &RetProc1[T1]{}
}

func (rt *RetProc1[T1]) ErrorReturn(err error) {
	if err != nil {
		rt.Tuple1.setFail(err)
		rt.ReturnOnly()
	}
}

func (rt *RetProc1[T1]) FalseReturn(b bool) {
	if !b {
		rt.Tuple1.setFail(ErrIsFalse)
		rt.ReturnOnly()
	}
}

func (rt *RetProc1[T1]) Return(v1 T1) {
	rt.Tuple1.Set(v1)
	rt.ReturnOnly()
}

func (rt *RetProc1[T1]) IfReturn(b bool, v1 T1) {
	if b {
		rt.Tuple1.Set(v1)
		rt.ReturnOnly()
	}
}

func (rt *RetProc1[T1]) Dov(fn func()) *Tuple1[T1] {
	rt.Do(fn)
	return &rt.Tuple1
}

func (rt *RetProc1[T1]) Dow(fn func()) (T1) {
	rt.Do(fn)
	return rt.Tuple1.Unwrap()
}

type RetTuple1[T1 any]struct {
	Tuple1[T1]
}

func Must1[T1 any](v1 T1) *RetTuple1[T1] {
	return &RetTuple1[T1]{Tuple1[T1]{v1}}
}

func (t *RetTuple1[T1]) OrReturnTo(ra failReturnAble) () {
	tupleMust(ra, t.V1, nil)
	return 
}

func (t *RetTuple1[T1]) OrDoReturnTo(fn func(error), ra failReturnAble) () {
	tupleMust(ra, t.V1, fn)
	return 
}

func (t *RetTuple1[T1]) OrFunc(fn func() ) () {
	if isFail(t.V1) {
		return
	}
	return 
}

func (t *RetTuple1[T1]) Or() () {
	if isFail(t.V1) {
		return 
	}
	return 
}


type Tuple2[T1 any, T2 any]struct {
	V1 T1
	V2 T2
}

func Tw2[T1 any, T2 any](v1 T1, v2 T2) Tuple2[T1, T2] {
	return Tuple2[T1, T2]{v1, v2}
}

func (t *Tuple2[T1, T2]) Unwrap() (T1, T2) {
	return t.V1, t.V2
}

func (t *Tuple2[T1, T2]) Set(v1 T1, v2 T2) {
	t.V1 = v1
	t.V2 = v2
}

func (t *Tuple2[T1, T2]) setFail(err error) {
	setError(&t.V2, err)
}

type RetProc2[T1 any, T2 any]struct {
	Tuple2[T1, T2]
	Proc
}

func NewRetProc2[T1 any, T2 any]() *RetProc2[T1, T2] {
	return &RetProc2[T1, T2]{}
}

func (rt *RetProc2[T1, T2]) ErrorReturn(err error) {
	if err != nil {
		rt.Tuple2.setFail(err)
		rt.ReturnOnly()
	}
}

func (rt *RetProc2[T1, T2]) FalseReturn(b bool) {
	if !b {
		rt.Tuple2.setFail(ErrIsFalse)
		rt.ReturnOnly()
	}
}

func (rt *RetProc2[T1, T2]) Return(v1 T1, v2 T2) {
	rt.Tuple2.Set(v1, v2)
	rt.ReturnOnly()
}

func (rt *RetProc2[T1, T2]) IfReturn(b bool, v1 T1, v2 T2) {
	if b {
		rt.Tuple2.Set(v1, v2)
		rt.ReturnOnly()
	}
}

func (rt *RetProc2[T1, T2]) Dov(fn func()) *Tuple2[T1, T2] {
	rt.Do(fn)
	return &rt.Tuple2
}

func (rt *RetProc2[T1, T2]) Dow(fn func()) (T1, T2) {
	rt.Do(fn)
	return rt.Tuple2.Unwrap()
}

type RetTuple2[T1 any, T2 any]struct {
	Tuple2[T1, T2]
}

func Must2[T1 any, T2 any](v1 T1, v2 T2) *RetTuple2[T1, T2] {
	return &RetTuple2[T1, T2]{Tuple2[T1, T2]{v1, v2}}
}

func (t *RetTuple2[T1, T2]) OrReturnTo(ra failReturnAble) (T1) {
	tupleMust(ra, t.V2, nil)
	return t.V1
}

func (t *RetTuple2[T1, T2]) OrDoReturnTo(fn func(error), ra failReturnAble) (T1) {
	tupleMust(ra, t.V2, fn)
	return t.V1
}

func (t *RetTuple2[T1, T2]) OrFunc(fn func() (T1)) (T1) {
	if isFail(t.V2) {
		return fn()
	}
	return t.V1
}

func (t *RetTuple2[T1, T2]) Or(v1 T1) (T1) {
	if isFail(t.V2) {
		return v1
	}
	return t.V1
}


type Tuple3[T1 any, T2 any, T3 any]struct {
	V1 T1
	V2 T2
	V3 T3
}

func Tw3[T1 any, T2 any, T3 any](v1 T1, v2 T2, v3 T3) Tuple3[T1, T2, T3] {
	return Tuple3[T1, T2, T3]{v1, v2, v3}
}

func (t *Tuple3[T1, T2, T3]) Unwrap() (T1, T2, T3) {
	return t.V1, t.V2, t.V3
}

func (t *Tuple3[T1, T2, T3]) Set(v1 T1, v2 T2, v3 T3) {
	t.V1 = v1
	t.V2 = v2
	t.V3 = v3
}

func (t *Tuple3[T1, T2, T3]) setFail(err error) {
	setError(&t.V3, err)
}

type RetProc3[T1 any, T2 any, T3 any]struct {
	Tuple3[T1, T2, T3]
	Proc
}

func NewRetProc3[T1 any, T2 any, T3 any]() *RetProc3[T1, T2, T3] {
	return &RetProc3[T1, T2, T3]{}
}

func (rt *RetProc3[T1, T2, T3]) ErrorReturn(err error) {
	if err != nil {
		rt.Tuple3.setFail(err)
		rt.ReturnOnly()
	}
}

func (rt *RetProc3[T1, T2, T3]) FalseReturn(b bool) {
	if !b {
		rt.Tuple3.setFail(ErrIsFalse)
		rt.ReturnOnly()
	}
}

func (rt *RetProc3[T1, T2, T3]) Return(v1 T1, v2 T2, v3 T3) {
	rt.Tuple3.Set(v1, v2, v3)
	rt.ReturnOnly()
}

func (rt *RetProc3[T1, T2, T3]) IfReturn(b bool, v1 T1, v2 T2, v3 T3) {
	if b {
		rt.Tuple3.Set(v1, v2, v3)
		rt.ReturnOnly()
	}
}

func (rt *RetProc3[T1, T2, T3]) Dov(fn func()) *Tuple3[T1, T2, T3] {
	rt.Do(fn)
	return &rt.Tuple3
}

func (rt *RetProc3[T1, T2, T3]) Dow(fn func()) (T1, T2, T3) {
	rt.Do(fn)
	return rt.Tuple3.Unwrap()
}

type RetTuple3[T1 any, T2 any, T3 any]struct {
	Tuple3[T1, T2, T3]
}

func Must3[T1 any, T2 any, T3 any](v1 T1, v2 T2, v3 T3) *RetTuple3[T1, T2, T3] {
	return &RetTuple3[T1, T2, T3]{Tuple3[T1, T2, T3]{v1, v2, v3}}
}

func (t *RetTuple3[T1, T2, T3]) OrReturnTo(ra failReturnAble) (T1, T2) {
	tupleMust(ra, t.V3, nil)
	return t.V1, t.V2
}

func (t *RetTuple3[T1, T2, T3]) OrDoReturnTo(fn func(error), ra failReturnAble) (T1, T2) {
	tupleMust(ra, t.V3, fn)
	return t.V1, t.V2
}

func (t *RetTuple3[T1, T2, T3]) OrFunc(fn func() (T1, T2)) (T1, T2) {
	if isFail(t.V3) {
		return fn()
	}
	return t.V1, t.V2
}

func (t *RetTuple3[T1, T2, T3]) Or(v1 T1, v2 T2) (T1, T2) {
	if isFail(t.V3) {
		return v1, v2
	}
	return t.V1, t.V2
}


type Tuple4[T1 any, T2 any, T3 any, T4 any]struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
}

func Tw4[T1 any, T2 any, T3 any, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) Tuple4[T1, T2, T3, T4] {
	return Tuple4[T1, T2, T3, T4]{v1, v2, v3, v4}
}

func (t *Tuple4[T1, T2, T3, T4]) Unwrap() (T1, T2, T3, T4) {
	return t.V1, t.V2, t.V3, t.V4
}

func (t *Tuple4[T1, T2, T3, T4]) Set(v1 T1, v2 T2, v3 T3, v4 T4) {
	t.V1 = v1
	t.V2 = v2
	t.V3 = v3
	t.V4 = v4
}

func (t *Tuple4[T1, T2, T3, T4]) setFail(err error) {
	setError(&t.V4, err)
}

type RetProc4[T1 any, T2 any, T3 any, T4 any]struct {
	Tuple4[T1, T2, T3, T4]
	Proc
}

func NewRetProc4[T1 any, T2 any, T3 any, T4 any]() *RetProc4[T1, T2, T3, T4] {
	return &RetProc4[T1, T2, T3, T4]{}
}

func (rt *RetProc4[T1, T2, T3, T4]) ErrorReturn(err error) {
	if err != nil {
		rt.Tuple4.setFail(err)
		rt.ReturnOnly()
	}
}

func (rt *RetProc4[T1, T2, T3, T4]) FalseReturn(b bool) {
	if !b {
		rt.Tuple4.setFail(ErrIsFalse)
		rt.ReturnOnly()
	}
}

func (rt *RetProc4[T1, T2, T3, T4]) Return(v1 T1, v2 T2, v3 T3, v4 T4) {
	rt.Tuple4.Set(v1, v2, v3, v4)
	rt.ReturnOnly()
}

func (rt *RetProc4[T1, T2, T3, T4]) IfReturn(b bool, v1 T1, v2 T2, v3 T3, v4 T4) {
	if b {
		rt.Tuple4.Set(v1, v2, v3, v4)
		rt.ReturnOnly()
	}
}

func (rt *RetProc4[T1, T2, T3, T4]) Dov(fn func()) *Tuple4[T1, T2, T3, T4] {
	rt.Do(fn)
	return &rt.Tuple4
}

func (rt *RetProc4[T1, T2, T3, T4]) Dow(fn func()) (T1, T2, T3, T4) {
	rt.Do(fn)
	return rt.Tuple4.Unwrap()
}

type RetTuple4[T1 any, T2 any, T3 any, T4 any]struct {
	Tuple4[T1, T2, T3, T4]
}

func Must4[T1 any, T2 any, T3 any, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) *RetTuple4[T1, T2, T3, T4] {
	return &RetTuple4[T1, T2, T3, T4]{Tuple4[T1, T2, T3, T4]{v1, v2, v3, v4}}
}

func (t *RetTuple4[T1, T2, T3, T4]) OrReturnTo(ra failReturnAble) (T1, T2, T3) {
	tupleMust(ra, t.V4, nil)
	return t.V1, t.V2, t.V3
}

func (t *RetTuple4[T1, T2, T3, T4]) OrDoReturnTo(fn func(error), ra failReturnAble) (T1, T2, T3) {
	tupleMust(ra, t.V4, fn)
	return t.V1, t.V2, t.V3
}

func (t *RetTuple4[T1, T2, T3, T4]) OrFunc(fn func() (T1, T2, T3)) (T1, T2, T3) {
	if isFail(t.V4) {
		return fn()
	}
	return t.V1, t.V2, t.V3
}

func (t *RetTuple4[T1, T2, T3, T4]) Or(v1 T1, v2 T2, v3 T3) (T1, T2, T3) {
	if isFail(t.V4) {
		return v1, v2, v3
	}
	return t.V1, t.V2, t.V3
}


type Tuple5[T1 any, T2 any, T3 any, T4 any, T5 any]struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
}

func Tw5[T1 any, T2 any, T3 any, T4 any, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Tuple5[T1, T2, T3, T4, T5] {
	return Tuple5[T1, T2, T3, T4, T5]{v1, v2, v3, v4, v5}
}

func (t *Tuple5[T1, T2, T3, T4, T5]) Unwrap() (T1, T2, T3, T4, T5) {
	return t.V1, t.V2, t.V3, t.V4, t.V5
}

func (t *Tuple5[T1, T2, T3, T4, T5]) Set(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) {
	t.V1 = v1
	t.V2 = v2
	t.V3 = v3
	t.V4 = v4
	t.V5 = v5
}

func (t *Tuple5[T1, T2, T3, T4, T5]) setFail(err error) {
	setError(&t.V5, err)
}

type RetProc5[T1 any, T2 any, T3 any, T4 any, T5 any]struct {
	Tuple5[T1, T2, T3, T4, T5]
	Proc
}

func NewRetProc5[T1 any, T2 any, T3 any, T4 any, T5 any]() *RetProc5[T1, T2, T3, T4, T5] {
	return &RetProc5[T1, T2, T3, T4, T5]{}
}

func (rt *RetProc5[T1, T2, T3, T4, T5]) ErrorReturn(err error) {
	if err != nil {
		rt.Tuple5.setFail(err)
		rt.ReturnOnly()
	}
}

func (rt *RetProc5[T1, T2, T3, T4, T5]) FalseReturn(b bool) {
	if !b {
		rt.Tuple5.setFail(ErrIsFalse)
		rt.ReturnOnly()
	}
}

func (rt *RetProc5[T1, T2, T3, T4, T5]) Return(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) {
	rt.Tuple5.Set(v1, v2, v3, v4, v5)
	rt.ReturnOnly()
}

func (rt *RetProc5[T1, T2, T3, T4, T5]) IfReturn(b bool, v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) {
	if b {
		rt.Tuple5.Set(v1, v2, v3, v4, v5)
		rt.ReturnOnly()
	}
}

func (rt *RetProc5[T1, T2, T3, T4, T5]) Dov(fn func()) *Tuple5[T1, T2, T3, T4, T5] {
	rt.Do(fn)
	return &rt.Tuple5
}

func (rt *RetProc5[T1, T2, T3, T4, T5]) Dow(fn func()) (T1, T2, T3, T4, T5) {
	rt.Do(fn)
	return rt.Tuple5.Unwrap()
}

type RetTuple5[T1 any, T2 any, T3 any, T4 any, T5 any]struct {
	Tuple5[T1, T2, T3, T4, T5]
}

func Must5[T1 any, T2 any, T3 any, T4 any, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) *RetTuple5[T1, T2, T3, T4, T5] {
	return &RetTuple5[T1, T2, T3, T4, T5]{Tuple5[T1, T2, T3, T4, T5]{v1, v2, v3, v4, v5}}
}

func (t *RetTuple5[T1, T2, T3, T4, T5]) OrReturnTo(ra failReturnAble) (T1, T2, T3, T4) {
	tupleMust(ra, t.V5, nil)
	return t.V1, t.V2, t.V3, t.V4
}

func (t *RetTuple5[T1, T2, T3, T4, T5]) OrDoReturnTo(fn func(error), ra failReturnAble) (T1, T2, T3, T4) {
	tupleMust(ra, t.V5, fn)
	return t.V1, t.V2, t.V3, t.V4
}

func (t *RetTuple5[T1, T2, T3, T4, T5]) OrFunc(fn func() (T1, T2, T3, T4)) (T1, T2, T3, T4) {
	if isFail(t.V5) {
		return fn()
	}
	return t.V1, t.V2, t.V3, t.V4
}

func (t *RetTuple5[T1, T2, T3, T4, T5]) Or(v1 T1, v2 T2, v3 T3, v4 T4) (T1, T2, T3, T4) {
	if isFail(t.V5) {
		return v1, v2, v3, v4
	}
	return t.V1, t.V2, t.V3, t.V4
}


type Tuple6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any]struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
}

func Tw6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) Tuple6[T1, T2, T3, T4, T5, T6] {
	return Tuple6[T1, T2, T3, T4, T5, T6]{v1, v2, v3, v4, v5, v6}
}

func (t *Tuple6[T1, T2, T3, T4, T5, T6]) Unwrap() (T1, T2, T3, T4, T5, T6) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6
}

func (t *Tuple6[T1, T2, T3, T4, T5, T6]) Set(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) {
	t.V1 = v1
	t.V2 = v2
	t.V3 = v3
	t.V4 = v4
	t.V5 = v5
	t.V6 = v6
}

func (t *Tuple6[T1, T2, T3, T4, T5, T6]) setFail(err error) {
	setError(&t.V6, err)
}

type RetProc6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any]struct {
	Tuple6[T1, T2, T3, T4, T5, T6]
	Proc
}

func NewRetProc6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any]() *RetProc6[T1, T2, T3, T4, T5, T6] {
	return &RetProc6[T1, T2, T3, T4, T5, T6]{}
}

func (rt *RetProc6[T1, T2, T3, T4, T5, T6]) ErrorReturn(err error) {
	if err != nil {
		rt.Tuple6.setFail(err)
		rt.ReturnOnly()
	}
}

func (rt *RetProc6[T1, T2, T3, T4, T5, T6]) FalseReturn(b bool) {
	if !b {
		rt.Tuple6.setFail(ErrIsFalse)
		rt.ReturnOnly()
	}
}

func (rt *RetProc6[T1, T2, T3, T4, T5, T6]) Return(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) {
	rt.Tuple6.Set(v1, v2, v3, v4, v5, v6)
	rt.ReturnOnly()
}

func (rt *RetProc6[T1, T2, T3, T4, T5, T6]) IfReturn(b bool, v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) {
	if b {
		rt.Tuple6.Set(v1, v2, v3, v4, v5, v6)
		rt.ReturnOnly()
	}
}

func (rt *RetProc6[T1, T2, T3, T4, T5, T6]) Dov(fn func()) *Tuple6[T1, T2, T3, T4, T5, T6] {
	rt.Do(fn)
	return &rt.Tuple6
}

func (rt *RetProc6[T1, T2, T3, T4, T5, T6]) Dow(fn func()) (T1, T2, T3, T4, T5, T6) {
	rt.Do(fn)
	return rt.Tuple6.Unwrap()
}

type RetTuple6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any]struct {
	Tuple6[T1, T2, T3, T4, T5, T6]
}

func Must6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) *RetTuple6[T1, T2, T3, T4, T5, T6] {
	return &RetTuple6[T1, T2, T3, T4, T5, T6]{Tuple6[T1, T2, T3, T4, T5, T6]{v1, v2, v3, v4, v5, v6}}
}

func (t *RetTuple6[T1, T2, T3, T4, T5, T6]) OrReturnTo(ra failReturnAble) (T1, T2, T3, T4, T5) {
	tupleMust(ra, t.V6, nil)
	return t.V1, t.V2, t.V3, t.V4, t.V5
}

func (t *RetTuple6[T1, T2, T3, T4, T5, T6]) OrDoReturnTo(fn func(error), ra failReturnAble) (T1, T2, T3, T4, T5) {
	tupleMust(ra, t.V6, fn)
	return t.V1, t.V2, t.V3, t.V4, t.V5
}

func (t *RetTuple6[T1, T2, T3, T4, T5, T6]) OrFunc(fn func() (T1, T2, T3, T4, T5)) (T1, T2, T3, T4, T5) {
	if isFail(t.V6) {
		return fn()
	}
	return t.V1, t.V2, t.V3, t.V4, t.V5
}

func (t *RetTuple6[T1, T2, T3, T4, T5, T6]) Or(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) (T1, T2, T3, T4, T5) {
	if isFail(t.V6) {
		return v1, v2, v3, v4, v5
	}
	return t.V1, t.V2, t.V3, t.V4, t.V5
}


type Tuple7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any]struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
}

func Tw7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) Tuple7[T1, T2, T3, T4, T5, T6, T7] {
	return Tuple7[T1, T2, T3, T4, T5, T6, T7]{v1, v2, v3, v4, v5, v6, v7}
}

func (t *Tuple7[T1, T2, T3, T4, T5, T6, T7]) Unwrap() (T1, T2, T3, T4, T5, T6, T7) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7
}

func (t *Tuple7[T1, T2, T3, T4, T5, T6, T7]) Set(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) {
	t.V1 = v1
	t.V2 = v2
	t.V3 = v3
	t.V4 = v4
	t.V5 = v5
	t.V6 = v6
	t.V7 = v7
}

func (t *Tuple7[T1, T2, T3, T4, T5, T6, T7]) setFail(err error) {
	setError(&t.V7, err)
}

type RetProc7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any]struct {
	Tuple7[T1, T2, T3, T4, T5, T6, T7]
	Proc
}

func NewRetProc7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any]() *RetProc7[T1, T2, T3, T4, T5, T6, T7] {
	return &RetProc7[T1, T2, T3, T4, T5, T6, T7]{}
}

func (rt *RetProc7[T1, T2, T3, T4, T5, T6, T7]) ErrorReturn(err error) {
	if err != nil {
		rt.Tuple7.setFail(err)
		rt.ReturnOnly()
	}
}

func (rt *RetProc7[T1, T2, T3, T4, T5, T6, T7]) FalseReturn(b bool) {
	if !b {
		rt.Tuple7.setFail(ErrIsFalse)
		rt.ReturnOnly()
	}
}

func (rt *RetProc7[T1, T2, T3, T4, T5, T6, T7]) Return(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) {
	rt.Tuple7.Set(v1, v2, v3, v4, v5, v6, v7)
	rt.ReturnOnly()
}

func (rt *RetProc7[T1, T2, T3, T4, T5, T6, T7]) IfReturn(b bool, v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) {
	if b {
		rt.Tuple7.Set(v1, v2, v3, v4, v5, v6, v7)
		rt.ReturnOnly()
	}
}

func (rt *RetProc7[T1, T2, T3, T4, T5, T6, T7]) Dov(fn func()) *Tuple7[T1, T2, T3, T4, T5, T6, T7] {
	rt.Do(fn)
	return &rt.Tuple7
}

func (rt *RetProc7[T1, T2, T3, T4, T5, T6, T7]) Dow(fn func()) (T1, T2, T3, T4, T5, T6, T7) {
	rt.Do(fn)
	return rt.Tuple7.Unwrap()
}

type RetTuple7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any]struct {
	Tuple7[T1, T2, T3, T4, T5, T6, T7]
}

func Must7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) *RetTuple7[T1, T2, T3, T4, T5, T6, T7] {
	return &RetTuple7[T1, T2, T3, T4, T5, T6, T7]{Tuple7[T1, T2, T3, T4, T5, T6, T7]{v1, v2, v3, v4, v5, v6, v7}}
}

func (t *RetTuple7[T1, T2, T3, T4, T5, T6, T7]) OrReturnTo(ra failReturnAble) (T1, T2, T3, T4, T5, T6) {
	tupleMust(ra, t.V7, nil)
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6
}

func (t *RetTuple7[T1, T2, T3, T4, T5, T6, T7]) OrDoReturnTo(fn func(error), ra failReturnAble) (T1, T2, T3, T4, T5, T6) {
	tupleMust(ra, t.V7, fn)
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6
}

func (t *RetTuple7[T1, T2, T3, T4, T5, T6, T7]) OrFunc(fn func() (T1, T2, T3, T4, T5, T6)) (T1, T2, T3, T4, T5, T6) {
	if isFail(t.V7) {
		return fn()
	}
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6
}

func (t *RetTuple7[T1, T2, T3, T4, T5, T6, T7]) Or(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) (T1, T2, T3, T4, T5, T6) {
	if isFail(t.V7) {
		return v1, v2, v3, v4, v5, v6
	}
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6
}


type Tuple8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any]struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
}

func Tw8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{v1, v2, v3, v4, v5, v6, v7, v8}
}

func (t *Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Unwrap() (T1, T2, T3, T4, T5, T6, T7, T8) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8
}

func (t *Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Set(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) {
	t.V1 = v1
	t.V2 = v2
	t.V3 = v3
	t.V4 = v4
	t.V5 = v5
	t.V6 = v6
	t.V7 = v7
	t.V8 = v8
}

func (t *Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) setFail(err error) {
	setError(&t.V8, err)
}

type RetProc8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any]struct {
	Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]
	Proc
}

func NewRetProc8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any]() *RetProc8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return &RetProc8[T1, T2, T3, T4, T5, T6, T7, T8]{}
}

func (rt *RetProc8[T1, T2, T3, T4, T5, T6, T7, T8]) ErrorReturn(err error) {
	if err != nil {
		rt.Tuple8.setFail(err)
		rt.ReturnOnly()
	}
}

func (rt *RetProc8[T1, T2, T3, T4, T5, T6, T7, T8]) FalseReturn(b bool) {
	if !b {
		rt.Tuple8.setFail(ErrIsFalse)
		rt.ReturnOnly()
	}
}

func (rt *RetProc8[T1, T2, T3, T4, T5, T6, T7, T8]) Return(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) {
	rt.Tuple8.Set(v1, v2, v3, v4, v5, v6, v7, v8)
	rt.ReturnOnly()
}

func (rt *RetProc8[T1, T2, T3, T4, T5, T6, T7, T8]) IfReturn(b bool, v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) {
	if b {
		rt.Tuple8.Set(v1, v2, v3, v4, v5, v6, v7, v8)
		rt.ReturnOnly()
	}
}

func (rt *RetProc8[T1, T2, T3, T4, T5, T6, T7, T8]) Dov(fn func()) *Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
	rt.Do(fn)
	return &rt.Tuple8
}

func (rt *RetProc8[T1, T2, T3, T4, T5, T6, T7, T8]) Dow(fn func()) (T1, T2, T3, T4, T5, T6, T7, T8) {
	rt.Do(fn)
	return rt.Tuple8.Unwrap()
}

type RetTuple8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any]struct {
	Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]
}

func Must8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) *RetTuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return &RetTuple8[T1, T2, T3, T4, T5, T6, T7, T8]{Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{v1, v2, v3, v4, v5, v6, v7, v8}}
}

func (t *RetTuple8[T1, T2, T3, T4, T5, T6, T7, T8]) OrReturnTo(ra failReturnAble) (T1, T2, T3, T4, T5, T6, T7) {
	tupleMust(ra, t.V8, nil)
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7
}

func (t *RetTuple8[T1, T2, T3, T4, T5, T6, T7, T8]) OrDoReturnTo(fn func(error), ra failReturnAble) (T1, T2, T3, T4, T5, T6, T7) {
	tupleMust(ra, t.V8, fn)
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7
}

func (t *RetTuple8[T1, T2, T3, T4, T5, T6, T7, T8]) OrFunc(fn func() (T1, T2, T3, T4, T5, T6, T7)) (T1, T2, T3, T4, T5, T6, T7) {
	if isFail(t.V8) {
		return fn()
	}
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7
}

func (t *RetTuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Or(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) (T1, T2, T3, T4, T5, T6, T7) {
	if isFail(t.V8) {
		return v1, v2, v3, v4, v5, v6, v7
	}
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7
}


type Tuple9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any]struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
}

func Tw9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{v1, v2, v3, v4, v5, v6, v7, v8, v9}
}

func (t *Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Unwrap() (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9
}

func (t *Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Set(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) {
	t.V1 = v1
	t.V2 = v2
	t.V3 = v3
	t.V4 = v4
	t.V5 = v5
	t.V6 = v6
	t.V7 = v7
	t.V8 = v8
	t.V9 = v9
}

func (t *Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) setFail(err error) {
	setError(&t.V9, err)
}

type RetProc9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any]struct {
	Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]
	Proc
}

func NewRetProc9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any]() *RetProc9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return &RetProc9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{}
}

func (rt *RetProc9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) ErrorReturn(err error) {
	if err != nil {
		rt.Tuple9.setFail(err)
		rt.ReturnOnly()
	}
}

func (rt *RetProc9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) FalseReturn(b bool) {
	if !b {
		rt.Tuple9.setFail(ErrIsFalse)
		rt.ReturnOnly()
	}
}

func (rt *RetProc9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Return(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) {
	rt.Tuple9.Set(v1, v2, v3, v4, v5, v6, v7, v8, v9)
	rt.ReturnOnly()
}

func (rt *RetProc9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) IfReturn(b bool, v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) {
	if b {
		rt.Tuple9.Set(v1, v2, v3, v4, v5, v6, v7, v8, v9)
		rt.ReturnOnly()
	}
}

func (rt *RetProc9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Dov(fn func()) *Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	rt.Do(fn)
	return &rt.Tuple9
}

func (rt *RetProc9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Dow(fn func()) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	rt.Do(fn)
	return rt.Tuple9.Unwrap()
}

type RetTuple9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any]struct {
	Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]
}

func Must9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) *RetTuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return &RetTuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{v1, v2, v3, v4, v5, v6, v7, v8, v9}}
}

func (t *RetTuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) OrReturnTo(ra failReturnAble) (T1, T2, T3, T4, T5, T6, T7, T8) {
	tupleMust(ra, t.V9, nil)
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8
}

func (t *RetTuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) OrDoReturnTo(fn func(error), ra failReturnAble) (T1, T2, T3, T4, T5, T6, T7, T8) {
	tupleMust(ra, t.V9, fn)
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8
}

func (t *RetTuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) OrFunc(fn func() (T1, T2, T3, T4, T5, T6, T7, T8)) (T1, T2, T3, T4, T5, T6, T7, T8) {
	if isFail(t.V9) {
		return fn()
	}
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8
}

func (t *RetTuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Or(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) (T1, T2, T3, T4, T5, T6, T7, T8) {
	if isFail(t.V9) {
		return v1, v2, v3, v4, v5, v6, v7, v8
	}
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8
}


type Tuple10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any]struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
}

func Tw10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10) Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{v1, v2, v3, v4, v5, v6, v7, v8, v9, v10}
}

func (t *Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Unwrap() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9, t.V10
}

func (t *Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Set(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10) {
	t.V1 = v1
	t.V2 = v2
	t.V3 = v3
	t.V4 = v4
	t.V5 = v5
	t.V6 = v6
	t.V7 = v7
	t.V8 = v8
	t.V9 = v9
	t.V10 = v10
}

func (t *Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) setFail(err error) {
	setError(&t.V10, err)
}

type RetProc10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any]struct {
	Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]
	Proc
}

func NewRetProc10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any]() *RetProc10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return &RetProc10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{}
}

func (rt *RetProc10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) ErrorReturn(err error) {
	if err != nil {
		rt.Tuple10.setFail(err)
		rt.ReturnOnly()
	}
}

func (rt *RetProc10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) FalseReturn(b bool) {
	if !b {
		rt.Tuple10.setFail(ErrIsFalse)
		rt.ReturnOnly()
	}
}

func (rt *RetProc10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Return(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10) {
	rt.Tuple10.Set(v1, v2, v3, v4, v5, v6, v7, v8, v9, v10)
	rt.ReturnOnly()
}

func (rt *RetProc10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) IfReturn(b bool, v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10) {
	if b {
		rt.Tuple10.Set(v1, v2, v3, v4, v5, v6, v7, v8, v9, v10)
		rt.ReturnOnly()
	}
}

func (rt *RetProc10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Dov(fn func()) *Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	rt.Do(fn)
	return &rt.Tuple10
}

func (rt *RetProc10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Dow(fn func()) (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) {
	rt.Do(fn)
	return rt.Tuple10.Unwrap()
}

type RetTuple10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any]struct {
	Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]
}

func Must10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10) *RetTuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return &RetTuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{v1, v2, v3, v4, v5, v6, v7, v8, v9, v10}}
}

func (t *RetTuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) OrReturnTo(ra failReturnAble) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	tupleMust(ra, t.V10, nil)
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9
}

func (t *RetTuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) OrDoReturnTo(fn func(error), ra failReturnAble) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	tupleMust(ra, t.V10, fn)
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9
}

func (t *RetTuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) OrFunc(fn func() (T1, T2, T3, T4, T5, T6, T7, T8, T9)) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	if isFail(t.V10) {
		return fn()
	}
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9
}

func (t *RetTuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Or(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	if isFail(t.V10) {
		return v1, v2, v3, v4, v5, v6, v7, v8, v9
	}
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8, t.V9
}

