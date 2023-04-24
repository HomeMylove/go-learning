# Golang

## 开始

### 第一个go程序

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

> 通过 <b>go build</b> 编译成二进制文件，可以指定名字

```shell
go build -o hello.exe main.go
```

### 注意事项

1) Go源文件以"go”为扩展名。

2) Go应用程序的执行入口是main()方法。

3) Go语言**严格区分大小写**。

4) Go方法由一条条语句构成，每个语句后不需要分号(Go语言会在每行后自动加分号)，这
   也体现出Golang的简洁性。

5) G0编译器是一行行进行编译的，因此我们一行就写一条语句，不能把多条语句写在同一
   个，否则报错

6) Go语言**定义的变量**或者**import的包**如果没有使用到，代码不能编译通过。

7) 大括号都是成对出现的，缺一不可。



## 变量

定义变量

```go
package main

import "fmt"

func main(){
    var i int
    i = 10
    fmt.Println("i=",i)
}
```

> 三种使用方式

1. 定义后不赋值，使用默认值

2. 根据值自动判断类型

   ```go
   var num = 3.14
   ```

3. 省略 var ，使用 **:=**     (注：左边的变量不应该是被定义过的)

   ```go
   name := "Yae"
   ```

   



















