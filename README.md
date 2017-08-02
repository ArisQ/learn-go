# Notes for learning go language

1. 所有的包名都应该使用小写字母

1. 可见性规则

    当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个**大写字母开头**，如：Group1，那么使用这种形式的标识符的对象就**可以被外部包的代码所使用**（客户端程序需要先导入这个包），这被称为**导出**（像面向对象语言中的 public）；标识符如果以**小写字母开头，则对包外是不可见的**，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 private ）。

    *（大写字母可以使用任何 Unicode 编码的字符，比如希腊文，不仅仅是 ASCII 码中的大写字母）。*

1. 包可以作为命名空间使用，帮助避免命名冲突（名称冲突）

    包别名：
    ```go
    import fm "fmt"
    ```

1. 所有的内存在 Go 中都是经过初始化的

1. 并行赋值。

    ```go
    a,b = b,a
    ```
1. 练习

    推断以下程序的输出，并解释你的答案，然后编译并执行它们。

    ```go
    //local_scope.go:
    package main

    var a = "G"

    func main() {
       n()
       m()
       n()
    }

    func n() { print(a) }

    func m() {
       a := "O"
       print(a)
    }
    ```

    答案: GOG

    ```go
    //global_scope.go:
    package main

    var a = "G"

    func main() {
       n()
       m()
       n()
    }

    func n() {
       print(a)
    }

    func m() {
       a = "O"
       print(a)
    }
    ```

    答案: GOO

    ```go
    //function_calls_function.go
    package main

    var a string

    func main() {
       a = "G"
       print(a)
       f1()
    }

    func f1() {
       a := "O"
       print(a)
       f2()
    }

    func f2() {
       print(a)
    }
    ```

    答案: GOG

1. **源文件名以 *_test* 结尾不会被编译，如只包含 *string_test.go* 文件的目录go build会失败**

    输出:

    ``go build _/E_/GitHub/learn-go/8_strings_tests: no buildable Go source files in E:\GitHub\learn-go\8_strings_tests``

    文件夹名以 *_tests* 结尾，不包含以 *_test.go* 结尾的文件会报错:
    ``?   	_/E_/GitHub/learn-go/8_strings_tests	[no test files]``

1. strings包，strconv包

    字符串处理与字符串转换

1. time包

    ```go
    time.Time
    time.Now()
    t.Day() t.Minute() t.Month() t.Year()
    Duration Location
    t.Format()
    time.After() time.Ticker()
    time.Sleep(Duration d)
    ```

1. 指针

    指针的格式化标识符为``%p``

    ```go
    var intPointer *int //=nil
    var aInt int
    intPointer = &aInt
    *intPointer = 0
    ```

    **在书写表达式类似 var p \*type 时，切记在 \* 号和指针名称间留有一个空格，因为 - var p\*type 是语法正确的，但是在更复杂的表达式中，它容易被误认为是一个乘法表达式！**

1. 控制结构

    * if-else

        ```go
        if condition1 {
            // do something
        } else if condition2 {
            // do something else
        }else {
            // catch-all or default
        }

        if initialization; condition {
    	// do something
        }
        ```

    * switch

        ```go
        switch var1 {
        	case val1:
        		//...
        	case val2:
        		//...
        	default:
        		//...
        }

        switch {
        	case condition1:
        		//...
        	case condition2:
        		//...
        	default:
        		//...
        }

        switch initialization {
        	case val1:
        		//...
        	case val2:
        		//...
        	default:
        		//...
        }
        ```

        * **可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：``case val1, val2, val3``**

        * **一旦成功地匹配到某个分支，在执行完相应代码后就会退出整个 switch 代码块，也就是说您不需要特别使用 break 语句来表示结束。**

        * ``fallthrough``

        *  **type-switch**

    * select

    * for (range)

        ```go
        for initialization; condition;  afterthought {}

        for condition {}

        // The following forms of for loop is equivalent
        for {}
        for ;; {}
        for true { }

        for index, value := range collection { }
        ```

        * **不需要括号 () 将它们括起来**

        * **不要在循环体内修改计数器**

    * break continue

        * break只能退出最内层循环

        * for、switch或select语句都可以配合标签（label）形式的标识符使用（标签的名称是大小写敏感的，为了提升可读性，一般建议使用全部大写字母）

    * return goto

1. 多返回值函数的错误

    ```go
    var err error   //error类型
    ```

    习惯用法:
    ```go
    value, err := pack1.Function1(param1)
    if err != nil {
        fmt.Printf("An error occured in pack1.Function1 with parameter %v", param1)
        return err
    }
    // 未发生错误，继续执行

    if err != nil {
    	fmt.Printf("Program stopping with error %v", err)
    	os.Exit(1)
    }

    if err := file.Chmod(0664); err != nil {
    	fmt.Println(err)
    	return err
    }

    if value, ok := readData(); ok {
        //…
    }
    ```

1. 函数

    * return语句可以带有零个或多个返回值

    * return语句也可以用来结束 for 死循环，或者结束一个协程（goroutine）

    * Go里面有三种类型的函数：

        * 普通的带有名字的函数

        * 匿名函数或者lambda函数

        * 方法

    * 函数参数、返回值以及它们的类型被统称为函数签名

    * **Go不允许函数重载**

    * 函数声明

        ```go
        func flushICache(begin, end uintptr) // implemented externally
        type binOp func(int, int) int
        //可以赋值给变量
        add := binOp
        ```

    * 函数不能在其它函数里面声明（不能嵌套）,除匿名函数

    * Go 没有泛型（generic）的概念, 不过在大部分情况下可以通过接口（interface），特别是空接口与类型选择（type switch）与/或者通过使用反射（reflection）来实现相似的功能

    * 没有参数的函数通常被称为 niladic 函数（niladic function）

    * 按值传递（call by value） 按引用传递（call by reference）

    * 在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）

    * 尽量使用**命名返回值**,会使代码更清晰、更简短，同时更加容易读懂

    * 空白符（blank identifier）``_``

    * 通过指针改变外部变量（outside variable）

    * **变参函数**

    * defer

    * 内置函数

        |名称|说明|
        |--|--|
        |close|用于管道通信|
        |len、cap|len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；cap 是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）|
        |new、make|new 和 make 均是用于分配内存：new 用于值类型和用户定义的类型，如自定义结构，make 用于内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针（详见第 10.1 节）。它也可以被用于基本类型：v := new(int)。make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作（详见第 7.2.3/4 节、第 8.1.1 节和第 14.2.1 节）new() 是一个函数，不要忘记它的括号|
        |copy、append|用于复制和连接切片|
        |panic、recover|两者均用于错误处理机制|
        |print、println|底层打印函数（详见第 4.2 节），在部署环境中建议使用 fmt 包|
        |complex、real|imag 用于创建和操作复数（详见第 4.5.2.2 节）|

    * 递归

        * 懒惰求值

    * 回调

    * 闭包 匿名函数

        ```go
        func Adder() func(int) int {
            var x int   //Adder退出以后，x还能被匿名函数访问
            return func(delta int) int {
                x += delta
                return x
            }
        }
        ```

        * 在闭包中使用到的变量可以是在闭包函数体内声明的，也可以是在外部函数声明的

        * 工厂函数

        * 高阶函数

            可以返回其它函数的函数和接受其它函数作为参数的函数

        * 闭包调试

            ``runtime``,``log``包

    * ``time``包计算函数执行时间

    * 内存缓存

        * 纯函数(相同输入必定获得相同输出的函数)

1. 数组array

    * 容器

        数组，切片，map


    * **数组长度最大为 2Gb**

    * 声明格式

        ```go
        var identifier [len]type
        ```

    * 数组是**可变的**

    * 当使用等于或者大于 len(array) 的索引时：如果编译器可以检测到，会给出索引超限的提示信息；如果检测不到的话编译会通过而运行时会panic

    * **Go 语言中的数组是一种<u>值类型</u>（不像 C/C++ 中是指向首元素的指针）**

    * 常量初始化

        ```go
        var arrAge = [5]int{18, 20, 15, 22, 16}
        var arrLazy = [...]int{5, 6, 7, 8, 22}  //... 可忽略，从技术上说它们其实变化成了切片
        var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}   //key: value syntax
        ```

    * 多维数组

        内部数组总是长度相同的。Go 语言的多维数组是矩形式的（唯一的例外是切片的数组）

    * 把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：

        * 传递数组的指针

        * 使用数组的切片

1. 切片slice

    * 切片（slice）是对数组一个连续片段的引用（该数组我们称之为**相关数组**，通常是匿名的）

    * 切片是一个长度可变的数组

    * 声明
        ```go
        var identifier []type //（不需要说明长度）
        ```

    * 未初始化之前默认为 nil，长度为 0

    * 初始化格式
        ```go
        var slice1 []type = arr1[start:end] //start:end 被称为 slice 表达式
        var slice1 []type = arr1[:]
        slice1 = &arr1

        //使用make创建切片
        var slice1 []type = make([]type, len)
        slice1 := make([]type, len)

        s = s[:cap(s)]  //扩展到上限
        ```

    * 切片只能向后移动

    * **绝对不要用指针指向 slice。切片本身已经是一个引用类型，所以它本身就是一个指针**

    * new make 区别
        * new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这种方法 **返回一个指向类型为 T，值为 0 的地址的指针**，它适用于值类型如数组和结构体；它相当于 &T{}。
        * make(T) **返回一个类型为 T 的初始值**，它**只适用于3种内建的引用类型**：切片、map 和 channel。

    * 函数

        ```go
        len()
        cap()
        //0 <= len(s) <= cap(s)
        ```

    * bytes 包

    * for-range

    * 切片重组（reslice）

    * 切片的复制与追加

        * copy

        * append

    * 字符号，数组和切片

        ```go
        c := []byte(s)  //字符串转切片
        substr := str[start:end]//子串
        ```

        * sort

        * **想要在数组或切片中搜索一个元素，该数组或切片必须先被排序（因为标准库的搜索算法使用的是二分法）**

1. map

    * 或被称为字典（Python）、hash 和 HashTable

    * 声明
        map 是**引用类型**，可以使用如下声明：

        ```go
        var map1 map[keytype]valuetype

        var map1 map[string]int
        ```

    * 未初始化的 map 的值是 nil

    * key 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float。数组、切片和结构体不能作为 key，但是指针和接口类型可以。如果要用结构体作为 key 可以提供 Key() 和 Hash() 方法。 value 可以是任意类型的；通过使用空接口类型,可以存储任意值，但是使用这种类型作为值时需要先做一次类型断言。

    * **map 也可以用函数作为自己的值**，这样就可以用来做**分支结构**：key 用来选择要执行的函数。

    * len(map1) 方法可以获得 map 中的 pair 数目

    * 初始化
        **不要使用 new，永远用 make 来构造 map**

        ```go
        var map1[keytype]valuetype = make(map[keytype]valuetype)
        ```

    * 用切片作为 map 的值使得一个key可以对应多个value

    * 测试键值是否存在,删除元素

        ```go
        val1, isPresent = map1[key1]
        delete(map1, key1)
        ```

    * for-range

        **map 不是按照 key 的顺序排列的，也不是按照 value 的序排列的**

    * map排序
    map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序。
    如果想为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序（使用 sort），然后可以使用切片的 for-range 方法打印出所有的 key 和 value。

1. 包package

    * 标准库

    * 正则表达式 regexp包

        ```go
        ok, _ := regexp.Match(pat, []byte(searchIn))
        ok, _ := regexp.MatchString(pat, searchIn)
        regexp.Compile(pat)
        ```

    * 锁和sync包

        * 资源竞争

        * map类型是非线程安全

        * ``sync.Mutex RWMutex/RLock() Once/once.Do(call)``

    * 精密计算和big包

        ```go
        big.Intbig.Rat //类型
        big.NewInt(n)   big.NewRat(N,D)
        Add()   Mul()
        ```

    * 自定义包

        * 自定义包要使用短小的不含有_(下划线)的小写单词来为文件命名

        * 编译，导入，初始化

        * **实际操作注意事项**

            * 自定义包要在``$GOPATH/src/``的文件夹下

            * 写完后使用``go build``,``go install``编译安装包

            * 调用方``import "pack"``导入包，如果包源文件为``$GOPATH/src/packname/packfile.go``，则为``import "packname/pack"``

        * 包的结构

            ``src / bin / pkg``

    * godoc

    * go install

    * go test

    * 平台相关代码

        ``prog_$GOOS.go prog_$GOARCH.go``

    * 通过git打包与安装

    * go外部包与项目

        [Go Walker](https://gowalker.org)

        * **go get需要VPN,或从GitHub克隆**

1. 结构体struct和方法method

    * 结构体是**值类型**

    * 结构体定义

        ```go
        type identifier struct {
            field1 type1
            field2 type2
            ...
        }
        type T struct {a, b int}
        var t *T=new(T)
        t:=new(T)
        ```

        * **如果字段在代码中从来也不会被用到，那么可以命名它为 _**

        * 结构体的字段可以是任何类型，甚至是结构体本身

    * 选择器selector 选择器符selector-notation ``.``

    * 初始化结构体

        ```go
        m:=&struct1{1,0.1,"foo"}//*struct1类型
        //结构体字面量struct-literal
        var m struct1
        m=struct1{1,0.1,"foo"}

        n:=struct2{fieldi:vi,fj:vj}
        ```

        * 混合字面量语法composite literal syntax ``&struct1{a, b, c}``是一种简写，底层仍然会调用 ``new()``

    * 类型 strcut1 在定义它的包 pack1 中必须是唯一的，它的完全类型名是：pack1.struct1

    * 链表Node 二叉树Tree

    * 工厂方法创建结构体 结构体工厂

        ```go
        new(File)
        &File{}
        size := unsafe.Sizeof(T{})
        ```

        * 强制使用工厂方法

            将类型设为包私有即可

    * 结构体标签tag

        ```go
        type TagType struct { // tags
        field1 bool   "An important answer"
        field2 string "The name of the thing"
        field3 int    "How much there are"
        }
        ```

    * 匿名字段

        ```go
        type innerS struct {
            in1 int
            in2 int
        }
        type outerS struct {
            b    int
            c    float32
            int  // anonymous field
            innerS //anonymous field
        }
        ```

        **在一个结构体中对于每一种数据类型只能有一个匿名字段**

    * 内嵌结构体

    * 命名冲突

        1. 外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式；

        1. 如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。没有办法来解决这种问题引起的二义性，必须由程序员自己修正

    * 方法

        * Go 方法是作用在**接收者(receiver)** 上的一个函数，接收者是某种类型的变量。因此方法是一种特殊类型的函数。

        * 接收者类型可以是（几乎）任何类型，可以是int,bool,string或结构体类型,数组的别名类型,函数类型。接收者不能是一个接口类型，因为接口是一个抽象定义，但是方法却是具体实现；**接收者不能是一个指针类型，但是它可以是任何其他允许类型的指针(何解？)**；如果这样做会引发一个编译错误：invalid receiver type…。

        * 类型和方法必须是同一个包

        * 类型 T（或 \*T）上的所有方法的集合叫做类型 T（或 \*T）的方法集。

        * 方法定义

            ```go
            func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
            func (_ receiver_type) methodName(parameter_list) (return_value_list) { ... }
            ```

        * 自动解引用

            **指针方法和值方法都可以在指针或非指针上被调用**,比如可以使用 p3.Abs() 来替代 (\*p3).Abs()。

        * 未导出字段

            * 结构体中小写开始的字段为未导出字段，不能被外部访问

            * 可以使用get,set方法

            * 并发访问

                使用``sync``包

        * 当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌

        * 结构体内嵌和自己在同一个包中的结构体时，可以彼此访问对方所有的字段和方法

        * 主要有两种方法来实现在类型中嵌入功能：

            1. 聚合（或组合）：包含一个所需功能类型的具名字段。

            1. 内嵌：内嵌（匿名地）所需功能类型。

        * 多重继承

            通过在类型中嵌入所有必要的父类型，可以简单的实现多重继承

        * 通用方法和方法命名

        * 组件编程（Component Programming）

    * String()方法

        * ``%T`` ``%#v``

        * **定义String()方法会被用在``fmt.Printf()``中生成默认输出（实测没有生效,需要显式调用）**

    * 垃圾回收

        * ``runtime``包  GC进程``runtime.GC()``

        * 对象被移除前的操作

            ```go
            runtime.SetFinalizer(obj, func(obj *typeObj))
            ```

1. 接口Interface与反射reflection

    * 接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。

    * 接口格式定义
        ```go
        type Namer interface {  //Namer是接口类型
            Method1(param_list) return_type
            Method2(param_list) return_type
            ...
        }
        ```

    * 按照约定，只包含一个方法的接口名字由方法名加[e]r后缀组成， 如 Printer、Reader、Writer、Logger、Converter 等等。还有一些不常用的方式，以able结尾或以I开头，如 Recoverable。

    * Go 语言中的接口都很**简短**，通常它们会包含**0 个、最多 3 个方法**。

    * 接口值

        ```go
        var ai Namer    //nil
        ```

        * 指向接口值的指针是非法的

    * Tips

        * 类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。

        * 实现某个接口的类型（除了实现接口方法外）可以有其他的方法。

        * 一个类型可以实现多个接口。

        * 接口类型可以包含一个实例的引用， 该实例的类型实现了此接口（接口是动态类型）。

    * 接口 类型数组

        ```go
        r := Rectangle{5, 3}
    	q := &Square{5}
    	shapes := []Shaper{r, q}
        ```

    * 接口嵌套

    * 类型断言

        类型断言用来测试在某个时刻 varI 是否包含类型 T 的值

        ```go
        //varI 必须是一个接口变量
        v := varI.(T)       // unchecked type assertion

        if v, ok := varI.(T); ok {  // checked type assertion
            Process(v)
            return
        }
        // varI is not of type T  

        if _, ok := varI.(T); ok {
            // ...
        }      
        ```

        * 如果转换合法，v 是 varI 转换到类型 T 的值，ok 会是 true；

        * 否则， v 是类型 T 的零值，ok 是 false，没有运行时错误发生。

    * 类型判断type-switch

        ```go
        switch t := areaIntf.(type) {
        case *Square:
        	fmt.Printf("Type Square %T with value %v\n", t, t)
        case *Circle:
        	fmt.Printf("Type Circle %T with value %v\n", t, t)
        case nil:
        	fmt.Printf("nil value: nothing to check?\n")
        default:
        	fmt.Printf("Unexpected type %T\n", t)
        }
        ```

    * 测试值是否实现某个接口

    * 方法集和接口

        * 指针方法可以通过指针调用
        * 值方法可以通过值调用
        * 接收者是值的方法可以通过指针调用，因为指针会首先被解引用
        * 接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
    即
        * 类型\*T的可调用方法集包含接受者为\*T或T的所有方法集
        * 类型T的可调用方法集包含接受者为T的所有方法
        * 类型T的可调用方法集不包含接受者为\*T的方法

    * 空接口

        ```go
        type Any interface{}
        var val interface{}
        ```

        * 不同类型变量的数组

            ```go
            type Vector struct {
                a []interface{}
            }
            ```

        * 复制数据切片刀空接口切片（需要一个一个复制）

        * 通用类型的节点数据结构（树）

        * 接口到接口转换（运行时检查）

    * **反射包**

        * 反射是**元编程**的一种形式

        * ``Type Value reflect.TypeOf() reflect.ValueOf() Kind()``

        * 通过反射修改值``Elem() CanSet SetFloat() ...``

        * 反射结构 ``NumField() Field(i) Method(n).Call(nil)``

        * Printf和反射

    * 接口与动态类型

        * 实现了某个接口的类型可以被传给任何以此接口为参数的函数

        * **依赖注入模式**

        * 接口提取``Rank()``

        * 显式地指明类型实现了某个接口,向接口的方法集中添加一个具有描述性名字的方法,如``ImplementsFooer()``方法

        * 空接口与函数重载

    * Go面向对象

        * 封装:以包分割，标识符大小写决定可见性

        * 继承:组合，内嵌类型

        * 多态:接口

    * 结构体，集合，高阶函数

1. 读写数据

    * 读取输入

        * ``os.Stdin fmt.Scanf() fmt.Sscanf() io.EOF`` ``os.Stdout os.Stderr`` ``bufio NewReader() ReadString() ``

        * **当``scanln(&str)``写成``scanln(str)``时不会报错，会直接跳过读取**

    * 文件读写

        * 读文件

            * ``os.File os.Open()``

            * ReadString 和 ReadBytes 方法的时候，我们不需要关心操作系统的类型，**直接使用\n** 就可以了

            * 读取整个文件到字符串

                ``io/ioutil ioutil.ReadFile() WriteFile``

            * ``bufio.Reader.Read()``

            * 按列读取``Fscan``

            * ``path.filePath Base()``             

            * ``compress``包 读取压缩文件，支持bzip2、flate、gzip、lzw 和 zlib

        * 写文件

            * ``bufio.Writer`` ``os.O_RDONLY os.O_WRONLY os.O_CREATE os.O_TRUNC``

    * 文件拷贝

        ``io.Copy()``

    * 读取命令行参数

        * ``os.Args``

        * ``flag``包

            ```go
            type Flag struct {
            	Name     string // name as it appears on command line
            	Usage    string // help message
            	Value    Value  // value as set
            	DefValue string // default value (as text); for usage message
            }
            ```
            **修改Arg,重新parse**

    * 切片读写文件

        ``ReadBytes()  Read([]byte)``

    * 用defer关闭文件

    * ``io.Writer bufio.Writer`` (12.8节)

    * JSON

        * ``json.Marshal() json.NewEncoder().Encode()``

        * **实测要进行JSON转换的结构体字段首字母要大写**

        * ``Unmarshal()``


        * ``Decoder Encoder``

    * XML

        * ``encoding/xml``包

    * Gob (Go binary)

    * 密码学

        * hash 包：实现了 adler32、crc32、crc64 和 fnv 校验；

        * crypto 包：实现了其它的 hash 算法，比如 md4、md5、sha1 等。以及完整地实现了 aes、blowfish、rc4、rsa、xtea 等加密算法。  

1. 错误处理

    * error接口

        ```go
        type error interface{
            Error() string
        }
        ```

       ``errorString``结构体

    * 定义新错误

        ```go
        err := errors.New("math - square root of negative number")
        ```

    * 包可以用额外方法定义错误

        ```go
        type Error interface {
        	Timeout() bool   // Is the error a timeout?
        	Temporary() bool // Is the error temporary?
        }
        ```

        * ``syscall.Errno os.EINAL``

        * ``fmt.Errorf()``

    * 运行时异常 panic

        * ``runtime.Error RuntimeError()``

            ```go
            panic("panic string")
            ```

        * **panic以后，defer语句会被执行**

            * 在多层嵌套的函数调用中调用 panic，可以马上中止当前函数的执行，所有的 defer 语句都会保证执行并把控制权交还给接收到 panic 的函数调用者。这样向上冒泡直到最顶层，并执行（每层的） defer，在栈顶处程序崩溃，并在命令行中用传给 panic 的值报告错误情况：这个终止过程就是**panicking**

            * 不能随意地用 panic 中止程序，必须尽力补救错误让程序能继续执行。

        * 从panic中恢复 recover

            * recover内建函数被用于从 panic 或 错误场景中恢复：让程序可以从 panicking 重新获得控制权，停止终止过程进而恢复正常执行。

            * recover 只能在 defer 修饰的函数，即panic 会导致栈被展开直到 defer 修饰的 recover() 被调用或者程序中止

        * 这是所有自定义包实现者应该遵守的最佳实践：

            * 在包内部，总是应该从 panic 中 recover：不允许显式的超出包范围的 panic()

            * 向包的调用者返回错误值（而不是 panic）。

        * 用闭包处理错误

    * 启动外部命令和程序

        * ``os`` 包有一个 ``StartProcess`` 函数可以调用或启动外部系统命令和二进制可执行文件；它的第一个参数是要运行的进程，第二个参数用来传递选项或参数，第三个参数是含有系统环境基本信息的结构体。这个函数返回被启动进程的 id（pid），或者启动失败返回错误。

        * ``exec`` 包中也有同样功能的更简单的结构体和函数；主要是 ``exec.Command(name string, arg ...string)`` 和 ``Run()``。首先需要用系统命令或可执行文件的名字创建一个 ``Command`` 对象，然后用这个对象作为接收者调用 ``Run()``。

1. 单元测试

    * 所有的包都应该有一定的必要文档

    * 对包的测试同样重要

    * gotest ``testing``包

    * 测试程序必须属于被测试的包，并且文件名满足这种形式 ``*_test.go``，所以测试代码和包中的业务代码是分开的。

    * **``_test``程序不会被普通的Go编译器编译**，所以当放应用部署到生产环境时它们不会被部署；只有 gotest 会编译所有的程序：普通程序和测试程序。

    * 测试文件中**必须导入``testing``包**，并写一些名字以TestZzz开头的全局函数，这里的 Zzz 是被测试函数的字母描述，如 TestFmtInterface，TestPayEmployees 等。

        ```go
        func TestAbcde(t *testing.T)
        ```

    * 通知测试失败

        ```go
        func (t *T) Fail()//标记测试函数为失败，然后继续执行（剩下的测试）。
        func (t *T) FailNow()//标记测试函数为失败并中止执行；文件中别的测试也被略过，继续执行下一个文件。
        func (t *T) Log(args ...interface{})//args被用默认的格式格式化并打印到错误日志中。
        func (t *T) Fatal(args ...interface{})//结合 先执行 3），然后执行 2）的效果。
        ```

    * ``--chatty``

        实测已改为``-test.v``

    * 基准测试

        ```go
        func BenchmarkReverse(b *testing.B) {
        	//...
        }
        ```

    * 测试用例

        * 正常的用例

        * 反面的用例（错误的输入，如用负数或字母代替数字，没有输入等）

        * 边界检查用例（如果参数的取值范围是 0 到 1000，检查 0 和 1000 的情况）

    * 测试数据表

    * 性能调试

        * gotest

            ``-cpuprofile -memprofile``

        * pprof

            ``topN web list``

1. 并行、并发和协程

    * 一个应用程序是运行在机器上的一个进程；进程是一个运行在自己内存地址空间里的独立执行体。

    * 一个进程由一个或多个操作系统线程组成，这些线程其实是共享同一个内存地址空间的一起工作的执行体。（并发程序）

    * 并行是一种通过使用多处理器以提高速度的能力。

    * 竞态 **不要使用全局变量或者共享内存，它们会给你的代码在并发运算的时候带来危险。**

    * 数据加锁会带来更高的复杂度，更容易使代码出错以及更低的性能

    * *thread-per-connection* 模型不够有效。

    * *Communicating Sequential Processes*（顺序通信处理）（CSP）

    * *message passing-model*（消息传递）

    * Go应用程序并发处理的部分被称作 *goroutines*（协程）

    * Go使用 *channels* 来同步协程

    * 协程是通过使用关键字 go 调用（执行）一个函数或者方法来实现的（也可以是匿名或者 lambda 函数）。

    * 协程可以在程序初始化的过程中运行（在 init() 函数中）。

    * ``runtime.GOMATPROCS(int)``核心数量

    * Go 协程（goroutines）和协程（coroutines）

    * 携程间通信 channel

        * 通道服务于通信的两个目的：**值的交换**，**同步**。

        * channel定义

            ```go
            var identifier chan datatype
            ```

        * 通道是引用类型

        * 通信操作符<-

            * 流向通道（发送）

                ``ch <- int1``

            * 从通道流出（接收）的三种方式：

                ```go
                int2 = <- ch
                int2 := <- ch
                if <- ch != 1000{
                    //...
                }
                ```

            * getData() 使用了无限循环：它随着 sendData() 的发送完成和 ch 变空也结束了。(?)

            * **默认情况下，通信是同步且无缓冲的**

                * 对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的：如果ch中的数据无人接收，就无法再给通道传入其他数据：新的输入无法在通道非空的情况下传入。所以发送操作会等待 ch 再次变为可用状态：就是通道值被接收时（可以传入变量）。

                * 对于同一个通道，接收操作是阻塞的（协程或函数中的），直到发送者可用：如果通道中没有数据，接收者就阻塞了。

            * 带缓冲的通道

                ```go
                ch1:=make(chan string,100)
                ch :=make(chan type, value)
                //value == 0 -> synchronous, unbuffered (阻塞）
                //value > 0 -> asynchronous, buffered（非阻塞）取决于value元素
                ```

                                //更具有伸缩性（scalable）

            * `cap`获取缓冲区容量

            * 在其他协程运行时让 main 程序无限阻塞的通常做法是在 main 函数的最后放置一个``{}``。

    * select 切换协程

        协程开关

        ```go
        select {
        case u:= <- ch1:
            //...
        case v:= <- ch2:
            //...
        default: // no value ready to be received//可选
            //...
        }
        ```

        * 在 select 中使用发送操作并且有 default可以确保发送不被阻塞！如果没有 case，select 就会一直阻塞。

        * select 语句实现了一种监听模式，通常用在（无限）循环中；在某种情况下，通过 break 语句使循环退出。

    * Ticker

        * ``time.After()``

    * 使用recover处理协程错误

    * 协程-通道类似于Simulink仿真的节点和连线

    * 惰性生成器 惰性求值

    * Futures模式

    * Multiplexing多元化

    * client-server pattern

    * teardown

1. 网络

    * TCP服务器

        ```go
        net.Listen()
        listener.Accept()
        conn.Read()
        conn.Write()        
        net.Dial()        
        ```

    * Web服务器

        ```go
        http.ListenAndServe()
        http.URL    //内含path
        http.Request()  //内含URL

        v,found:=request.Form["v"]//var v map[string][[]string]
        req.FormValue("v")
        //or
        request.ParseForm()
        v,found:=request.Fomr["v"]

        http.HandleFunc()//HandlerFunc(w ReponseWriter, req)
        ```

    * 访问并读取页面

        ```go
        http.Head()
        res,err:=http.Get()
        //ioutil.ReadAll(res.Body)
        http.Redirect()
        http.NotFound()
        http.Error()
        ```

    * ``text/template``

        ```go
        t:=template.New()
        t.Parse()
        t.Execute()
        template.must()
        ```

        * if-else

            ``{{if `anything`}} Print IF part. {{else}} Print ELSE part.{{end}}``

        * ``{{.}}``

        * ``with-end``

        * ``$``

        * ``{{range pipeline}} T1 {{else}} T0 {{end}}``

        * Predefined template functions

    * ``net/rpc-package`` remote procedure calls

    * ``netchan`` channels over a network (``old/netchan``)

    * ``code.google.com/p/go/websocket`` websocket

    * ``net/smtp SendMail()``

1. 常见错误

    * 由于编译优化和依赖于使用缓存操作的字符串大小，当循环次数大于15时，效率才会更佳。

    * make new

        * 切片、映射和通道，使用make

        * 数组、结构体和所有的值类型，使用new

    * 永远不要使用一个指针指向一个接口类型，因为它已经是一个指针。

    * **如果你传递一个指针，而不是一个值类型，go编译器大多数情况下会认为需要创建一个对象，并将对象移动到堆上，所以会导致额外的内存分配**

    * **当且仅当代码中并发执行非常重要，才使用协程和通道。**

1. 模式

    * comma,ok模式

    * defer模式

    * visibility模式

    * operator模式 interface

1. 使用代码片段

    * 字符串、数组、切片、映射、结构体、接口、函数、文件、协程、通道、模板、

1. 构建完整的应用程序

    * map goroutine gob json rpc
