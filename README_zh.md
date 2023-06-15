Myhig 是一个可以让你开心写go的库

# 依赖

go 版本 >= 1.18


# 获取
```
go get github.com/neverlee/myhig
```

# 功能
让 go 优雅地处理检查返回，处理错误。提供类似 rust 的错误处理机制。

* 提供优雅的方式处理 `if err` 检查
* 提供优雅的方法处理 `if ok` 检查
* 提供优雅的方法返回

# 示例

原始代码，这里一个个的检查是不是让人很不舒服
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

有了myhig，就可以将上面的代码改成这样:
```
func ReadConfig(path string) (*Config, error) {
    rp := GetProcOfFunc1r2(ReadConfig) // 或者 rp := NewRetProc2[*Config, error]()
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
这样是不是看起来好很多了。


# 文档索引

## 流程控制器 `type RetProc$`

RetProc$ 是一系列的流程控制器。上面例子`rp := GetProcOfFunc1r2(ReadConfig)`中的rp就是一个流程控制器对象。实际使用时，占位符`$`需要替换成数字。例如`RetProc1, RetProc2, RetProc3 ...`。该数字为该流程控制器返回参数的个数。

举例:
```
// RetProc2 是一个用于那些返回参数是两个的函数的流程控制器，比如下面这些函数:
func A1() (int, bool)
func A2(int) (int, bool)
func A3(string, int) (int, bool)
func A4(string, int, decimal.Decimal) (int, error)

// RetProc2 是一个用于那些返回参数是三个的函数的流程控制器，比如下面这些函数:
func B1() (int, int, error)
func B2(int) (int, int, error)
func B3(int, string) (int, int, error)
func B4(int, string, decimal.Decimal) (int, int, error)
```


## 构造 RetProc$

### 第一种方式，基于泛型构造
构造函数 `func NewRetProc$`

示例:
```
// 构造 NewRetProc
rp := NewRetProc1[int]()
rp := NewRetProc2[int, bool]()
rp := NewRetProc3[string, int, error]()
rp := NewRetProc5[string, int, error, bool, decimal.Decimal]()
```
`$` 为返回参数个数，也是泛型参数个数


### 第二种方法，传入外层参数名进行构造
* 对于常规函数使用 `func GetProcOfFunc$1r$2(f func)`
* 对于带可变参数的函数使用 `func GetProcOfFunc$1vr$2(f func)`
示例:
```
// 常规函数示例
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

// 可变参函数示例
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

* 第一个占位符`$1`需要替换为函数f的入参个数
* 第二个占位符`$2`需要替换为函数f的返回参数个数

## RetProc$ 的方法

### `func (rt *RetProc$) ErrorReturn(err error)`
如果 err 不为 nil，就返回`..., err` 或 `..., false`，否则什么也不做

示例:
```
// 函数返回参数最后一个是error类型
func DoSomething(arg1 string, arg2 int) (int, string, error) {
    rp := GetProcOfFunc2r3(DoSomething) // or: rp := NewRetProc3[int, string, error]()
    return rp.Dow(func() {
        err := fmt.Errorf("an error")
        rp.ErrorReturn(err) // 相当于下面三行的效果
        // if anerr != nil {
        //     return 0, "", err
        // }
        
        // ...
    })
}

// 函数返回参数最后一个是bool类型
func FuncWithALastBoolReturn(arg1 string, arg2 int) (int, string, bool) {
    rp := GetProcOfFunc2r3(DoSomething) // or: rp := NewRetProc3[int, string, bool]()
    return rp.Dow(func() {
        err := CheckSomeThing()

        rp.ErrorReturn(err) // just like these
        // if err != nil {
        //     return 0, "", false
        // }
        
        // ...
    })
}
```

### `func (rt *RetProc$) FalseReturn(b bool)`
如果 b 不为 nil，就返回`..., err` 或 `..., false`

示例:
```
// 函数返回参数最后一个是error类型
func DoSomething(arg1 string, arg2 int) (int, string, error) {
    rp := GetProcOfFunc2r3(DoSomething) // or: rp := NewRetProc3[int, string, error]()
    return rp.Dow(func() {
        var ok bool
        rp.FalseReturn(ok) // 相当于下面三行的效果
        // if !ok {
        //     return 0, "", myhig.ErrIsFalse
        // }
        
        // ...
    })
}

// 函数返回参数最后一个是bool类型
func FuncWithALastBoolReturn(arg1 string, arg2 int) (int, string, bool) {
    rp := GetProcOfFunc2r3(DoSomething) // or: rp := NewRetProc3[int, string, bool]()
    return rp.Dow(func() {
        var ok bool
        rp.FalseReturn(ok) // 相当于下面三行的效果
        // if !ok {
        //     return 0, "", false
        // }
        
        // ...
    })
}
```

### `func (rt *RetProc$) Return(...)`
返回函数，结束该Proc并返回数据

示例:
```
// 函数返回参数最后一个是error类型
func DoSomething(arg1 string, arg2 int) (int, string) {
    rp := GetProcOfFunc2r2(DoSomething) // or: rp := NewRetProc2[int, string]()
    return rp.Dow(func() {
        rp.Return(9, "hello") // 相当于下面一行的效果
        // return 9, "hello"
    })
}
```

### `func (rt *RetProc$) IfReturn(b bool, ...)`
条件返回函数，如果b为true结束该Proc并返回数据

示例:
```
// 函数返回参数最后一个是error类型
func DoSomething(m map[int]string) (int, string) {
    rp := GetProcOfFunc1r2(DoSomething) // or: rp := NewRetProc2[int, string]()
    return rp.Dow(func() {
        var isEnd bool
        rp.IfReturn(isEnd, 10, "hello") // 相当于下面三行的效果
        // if isEnd {
        //     return 10, "hello"
        // }
        
        // ...
    })
}
```


### `func (rt *RetProc$) Dov(fn func()) *TupleX[...]`
启动流程并执行fn，然后返回一个tuple

示例:
```
// 函数返回参数最后一个是error类型
func DoSomething(arg1 string, arg2 int) (int, string) {
    rp := GetProcOfFunc2r2(DoSomething) // or: rp := NewRetProc2[int, string]()

    ret := rp.Dov(func() {
        rp.Return(9, "hello")
    })
    return ret.Unwrap()

    // 也可以写成
    // return rp.Dov(func() {
    //     rp.Return(9, "hello")
    // }).Unwrap()
}
```

### `func (rt *RetProc$) Dow(fn func()) (...)`
启动流程并执行fn，然后返回全部参数

示例:
```
// 函数返回参数最后一个是error类型
func DoSomething(arg1 string, arg2 int) (int, string) {
    rp := GetProcOfFunc2r2(DoSomething) // or: rp := NewRetProc2[int, string]()

    return rp.Dow(func() {
        rp.Return(9, "hello")
    })

    // 相当于
    // return rp.Dov(func() {
    //     rp.Return(9, "hello")
    // }).Unwrap()
}
```


## Must$ 系函数
* Must$最后一个入参的类型必须是bool或者error，否则会panic
* 通常情况下，Must$ 需要与 RetProc$ 配合使用
* 与RetProc$配合使用时, RetProc$的最后一个返回参数必须是bool或者error，否则会panic
* 其中 Must$ 中的占位符$需要替换为调用Must$时的入参个数

### `Must$(..., last).OrReturnTo(rp *RetProc$)`
如果 last 是 false 或者不为nil的 error，则让流程结束并返回error或者false，类似ErrorReturn和FalseReturn

示例：
```
func DoSomething(path string) (int, string, error) {
    rp := GetProcOfFunc1r3(ReadConfig) // 或者 rp := NewRetProc3[int, string, error]()
    return rp.Dow(func() {
        f := myhig.Must2(os.Open(path)).OrReturnTo(rp)

        // 相当于
        // f, err := os.Open(path)
        // rp.ErrorReturn(err)

        // 也相当于
        // f, err := os.Open(path)
        // if err != nil {
        //     return 0, "", err
        // }

        // ...
    })
}
```

### `Must$(..., last).OrDoReturnTo(fn func(error),  rp *RetProc$)`
如果 last 是 false 或者不为nil的 error，则调用fn，然后让流程结束并返回error或者false，类似ErrorReturn和FalseReturn

示例：
```
func DoSomething(path string) (int, string, error) {
    rp := GetProcOfFunc1r3(ReadConfig) // 或者 rp := NewRetProc3[int, string, error]()
    return rp.Dow(func() {
        errCheck := func (err error) {
            if errors.Is(err, io.EOF) {
                fmt.Println("error is io.EOF)
            }
        }
        f := myhig.Must2(os.Open(path)).OrDoReturnTo(errCheck, rp)

        // 相当于
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
如果 last 是 false 或者不为nil的 error，则返回 Or 中的内容

示例：
```
func DoSomething(m map[int]string) (int, string, error) {
    rp := GetProcOfFunc1r3(ReadConfig) // 或者 rp := NewRetProc3[int, string, error]()
    return rp.Dow(func() {
        v := myhig.Must2(m[0]).Or("it is zero")

        // 相当于
        // v, ok := m[0]
        // if !ok {
        //     v = "it is zero"
        // }

        // ...
    })
}
```

### `Must$(..., last).OrFunc(fn func()(...))`
如果 last 是 false 或者不为nil的 error，则调用fn，返回 fn 的返回结果
示例：
```
func DoSomething(m map[int]string) (int, string, error) {
    rp := GetProcOfFunc1r3(ReadConfig) // 或者 rp := NewRetProc3[int, string, error]()
    return rp.Dow(func() {
        v := myhig.Must2(m[0]).OrFunc(func() string {
            return "str1" + "str2"
        })

        // 相当于
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


