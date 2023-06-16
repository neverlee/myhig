Myhig is a library for making you happy in golang

[简体中文文档](https://github.com/neverlee/myhig/blob/main/README_zh.md)

# Dependencies

go version >= 1.18


# How to get
```
go get github.com/neverlee/myhig
```

# What myhig can do
Let golang gracefully handle checking returns and handling errors. Provides a rust-like error handling mechanism

* Provides an elegant way to handle `if err` checks
* Provides an elegant way to handle `if ok` checks
* Provides an elegant way to handle `return`

# Example

Raw Code
```
func ReadConfig(path string) (*Config, error) {
    f, err := os.Open(path)
    defer f.Close()
    if err != nil {
        return nil, err
    }

    buf, err := ioutil.ReadAll(f)
    if err != nil {
        return nil, err
    }

    var conf Config
    if err := json.Unmarshal(buf, &conf); err != nil {
        return nil, err
    }

    if formal, ok := conf.Keys["formal"]; !ok {
        conf.GWKey = formal
    } else {
        conf.GWKey = "a test key"
    }

    return &conf, nil
}
```

with myhig
```
func ReadConfig(path string) (*Config, error) {
    rp := GetProcOfFunc1r2(ReadConfig) // or: rp := NewRetProc2[*Config, error]()
    return rp.Dow(func() {
        f := myhig.Must2(os.Open(path)).OrReturnTo(rp)
        defer f.Close()

        buf := myhig.Must2(ioutil.ReadAll(f)).OrReturnTo(rp)

        var conf Config
        rp.ErrorReturn(json.Unmarshal(buf, &conf))

        conf.GWKey = myhig.Must2(conf.Keys["formal"]).Or("a test key")

        rp.Return(&conf, nil)
    })
}
```
Does it looks much better?


# Reference

## Process Controller `type RetProc$`

RetProc$ is a series of process controllers. 
The rp in the above example `rp := GetProcOfFunc1r2(ReadConfig)` is a process controller object.
You should replace the placeholder `$` by a number. Such as `RetProc1, RetProc2, RetProc3 ...`. it is the number of parameters returned of the process controller.

Example:
```
// RetProc2 is used for the functions that has two parameters returned. Such as:
func A1() (int, bool)
func A2(int) (int, bool)
func A3(string, int) (int, bool)
func A4(string, int, decimal.Decimal) (int, error)

// RetProc3 is used for the functions that has three parameters returned. Such as:
func B1() (int, int, error)
func B2(int) (int, int, error)
func B3(int, string) (int, int, error)
func B4(int, string, decimal.Decimal) (int, int, error)
```


## Create RetProc$

### The first is based on generic parameters
Constructor function `func NewRetProc$`

示例:
```
// create NewRetProc
rp := NewRetProc1[int]()
rp := NewRetProc2[int, bool]()
rp := NewRetProc3[string, int, error]()
rp := NewRetProc5[string, int, error, bool, decimal.Decimal]()
```
`$` should be replaced with the number of generic parameters.


### The second is based on the outer function
* for normal function `func GetProcOfFunc$1r$2(f func)`
* for variadic function `func GetProcOfFunc$1vr$2(f func)`

Example:
```
// for normal function
func NormalFunc1(a, b, c, d int) int {
    rp := GetProcOfFunc4r1(NormalFunc1)
    // ...
}

func NormalFunc2(a, b, c, d, e int) (int, string) {
    rp := GetProcOfFunc5r2(NormalFunc2)
    // ...
}

func NormalFunc3(a, b, c, d, e int) (int, string, bool) {
    rp := GetProcOfFunc5r3(NormalFunc3)
    // ...
}

func NormalFunc4(a, b, c int) (int, string, bool, error) {
    rp := GetProcOfFunc3r5(NormalFunc3)
    // ...
}

// for variadic function
func NormalFunc1(a, b, c, d... int) int {
    rp := GetProcOfFunc4vr1(NormalFunc1)
    // ...
}

func NormalFunc2(a, b, c, d, e... int) (int, string) {
    rp := GetProcOfFunc5vr2(NormalFunc2)
    // ...
}

func NormalFunc3(a, b, c, d, e... int) (int, string, bool) {
    rp := GetProcOfFunc5vr3(NormalFunc3)
    // ...
}

func NormalFunc4(a, b, c... int) (int, string, bool, error) {
    rp := GetProcOfFunc3vr5(NormalFunc3)
    // ...
}
```

* The placeholder `$1` should be replaced with the arity of the function f
* The placeholder `$2` should be replaced with the number of parameters returned of the function f

## Methods of RetProc$

### `func (rt *RetProc$) ErrorReturn(err error)`
if err is not nil, set the `..., err` or `..., false` as the return values and end the process.

Example:
```
// the last parameter returned is error type
func DoSomething(arg1 string, arg2 int) (int, string, error) {
    rp := GetProcOfFunc2r3(DoSomething) // or: rp := NewRetProc3[int, string, error]()
    return rp.Dow(func() {
        err := fmt.Errorf("an error")
        rp.ErrorReturn(err) // It works like the following code
        // if anerr != nil {
        //     return 0, "", err
        // }
        
        // ...
    })
}

// the last parameter returned is bool type
func FuncWithALastBoolReturn(arg1 string, arg2 int) (int, string, bool) {
    rp := GetProcOfFunc2r3(DoSomething) // or: rp := NewRetProc3[int, string, bool]()
    return rp.Dow(func() {
        err := CheckSomeThing()

        rp.ErrorReturn(err) // It works like the following code
        // if err != nil {
        //     return 0, "", false
        // }
        
        // ...
    })
}
```

### `func (rt *RetProc$) FalseReturn(b bool)`
if b is not false, set the `..., err` or `..., false` as the return values and end the process.

Example:
```
// the last parameter returned is error type
func DoSomething(arg1 string, arg2 int) (int, string, error) {
    rp := GetProcOfFunc2r3(DoSomething) // or: rp := NewRetProc3[int, string, error]()
    return rp.Dow(func() {
        var ok bool
        rp.FalseReturn(ok) // It works like the following code
        // if !ok {
        //     return 0, "", myhig.ErrIsFalse
        // }
        
        // ...
    })
}

// the last parameter returned is bool type
func FuncWithALastBoolReturn(arg1 string, arg2 int) (int, string, bool) {
    rp := GetProcOfFunc2r3(DoSomething) // or: rp := NewRetProc3[int, string, bool]()
    return rp.Dow(func() {
        var ok bool
        rp.FalseReturn(ok) // It works like the following code
        // if !ok {
        //     return 0, "", false
        // }
        
        // ...
    })
}
```

### `func (rt *RetProc$) Return(...)`
set the arguments as the return values and end the process.

Exmaple:
```
func DoSomething(arg1 string, arg2 int) (int, string) {
    rp := GetProcOfFunc2r2(DoSomething) // or: rp := NewRetProc2[int, string]()
    return rp.Dow(func() {
        rp.Return(9, "hello") // It works like the following code
        // return 9, "hello"
    })
}
```

### `func (rt *RetProc$) IfReturn(b bool, ...)`
if b is true, set the arguments as the return values and end the process.

Example:
```
func DoSomething(m map[int]string) (int, string) {
    rp := GetProcOfFunc1r2(DoSomething) // or: rp := NewRetProc2[int, string]()
    return rp.Dow(func() {
        var isEnd bool
        rp.IfReturn(isEnd, 10, "hello") // It works like the following code
        // if isEnd {
        //     return 10, "hello"
        // }
        
        // ...
    })
}
```


### `func (rt *RetProc$) Dov(fn func()) *TupleX[...]`
start the process and run the fn and return a tuple

Example:
```
func DoSomething(arg1 string, arg2 int) (int, string) {
    rp := GetProcOfFunc2r2(DoSomething) // or: rp := NewRetProc2[int, string]()

    ret := rp.Dov(func() {
        rp.Return(9, "hello")
    })
    return ret.Unwrap()

    // It works like the following code
    // return rp.Dov(func() {
    //     rp.Return(9, "hello")
    // }).Unwrap()
}
```

### `func (rt *RetProc$) Dow(fn func()) (...)`
start the process and run the fn and return all values

示例:
```
func DoSomething(arg1 string, arg2 int) (int, string) {
    rp := GetProcOfFunc2r2(DoSomething) // or: rp := NewRetProc2[int, string]()

    return rp.Dow(func() {
        rp.Return(9, "hello")
    })

    // It works like the following code
    // return rp.Dov(func() {
    //     rp.Return(9, "hello")
    // }).Unwrap()
}
```


## Must$ functions
* The last parameter of Must$ must be bool type or error type, otherwise it will panic
* Must$ is usually used in conjunction with RetProc$
* When Must$ works with RetProc$, the last parameter returns of RetProc$ must be bool type or error type, otherwise it will panic
* The placeholder `$` of Must$ should be replced with the arity of Must$

### `Must$(..., last).OrReturnTo(rp *RetProc$)`
if 'last' is false or a non empty error, call the rp.ErrorReturn(last) or rp.FalseReturn(last)

Example：
```
func DoSomething(path string) (int, string, error) {
    rp := GetProcOfFunc1r3(ReadConfig) // or: rp := NewRetProc3[int, string, error]()
    return rp.Dow(func() {
        f := myhig.Must2(os.Open(path)).OrReturnTo(rp)

        // It works like the following code
        // f, err := os.Open(path)
        // rp.ErrorReturn(err)

        // It works like the following code too也
        // f, err := os.Open(path)
        // if err != nil {
        //     return 0, "", err
        // }

        // ...
    })
}
```

### `Must$(..., last).OrDoReturnTo(fn func(error),  rp *RetProc$)`
if 'last' is false or a non empty error, call the fn(err) then call rp.ErrorReturn(last) or rp.FalseReturn(last)

Example：
```
func DoSomething(path string) (int, string, error) {
    rp := GetProcOfFunc1r3(ReadConfig) // or: rp := NewRetProc3[int, string, error]()
    return rp.Dow(func() {
        errCheck := func (err error) {
            if errors.Is(err, io.EOF) {
                fmt.Println("error is io.EOF)
            }
        }
        f := myhig.Must2(os.Open(path)).OrDoReturnTo(errCheck, rp)

        // It works like the following code
        // f, err := os.Open(path)
        // if err != nil {
        //     errCheck(err)
        //     return 0, "", err
        // }

        // ...
    })
}
```

### `Must$(..., last).Or(...)`
if 'last' is false or a non empty error, return the parameters of 'Or'

Example：
```
func DoSomething(m map[int]string) (int, string, error) {
    rp := GetProcOfFunc1r3(ReadConfig) // or: rp := NewRetProc3[int, string, error]()
    return rp.Dow(func() {
        v := myhig.Must2(m[0]).Or("it is zero")

        // It works like the following code
        // v, ok := m[0]
        // if !ok {
        //     v = "it is zero"
        // }

        // ...
    })
}
```

### `Must$(..., last).OrFunc(fn func()(...))`
if 'last' is false or a non empty error, return the result of `fn()`

Example：
```
func DoSomething(m map[int]string) (int, string, error) {
    rp := GetProcOfFunc1r3(ReadConfig) // or: rp := NewRetProc3[int, string, error]()
    return rp.Dow(func() {
        v := myhig.Must2(m[0]).OrFunc(func() string {
            return "str1" + "str2"
        })

        // It works like the following code
        // v, ok := m[0]
        // if !ok {
        //     v = func() string {
        //         return "str1" + "str2"
        //     }()
        // }

        // ...
    })
}
```

# Warning
* Because the `*Return`s of RetProc$ are slower than return：
    * It is not recommended to use it in small functions that need to be called frequently
    * It is recommended to be used in complex business functions (such as HTTP, RPC interface implementation functions)
* When using RetProc$, it does not affect the normal call of panic. Can be used with confidence
