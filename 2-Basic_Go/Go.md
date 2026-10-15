# Go
Go is a high-level, general-purpose programming language that is statically typed and compiled. It is known for the simplicity of its syntax and the efficiency of development that it enables through the inclusion of a large standard library supplying many needs for common projects.

## Go Documentation
This website contains everything about Go https://go.dev/doc/. To view the standard library, Docs -> Standard library.

The official Git repo is https://github.com/golang/go.

For the Wikipedia fans https://en.wikipedia.org/wiki/Go_(programming_language).

### Types
    bool
    int8
    uint8
    int16
    uint16
    int32
    uint32
    int64
    uint64
    int
    uint
    uintptr
    float32
    float64
    complex64
    complex128
    string

### Keywords
    break
    case
    chan
    const
    continue
    default
    defer
    else
    fallthrough
    for
    func
    go
    goto
    if
    import
    interface
    map
    package
    range
    return
    select
    struct
    switch
    type
    var


## Starting a Go Project/Module
Run ```go mod init ProjectName``` to start development. This will create a go.mod file. Then create a main file that begins with ```package main``` and begin to code.

Execution begins with a typical main function, there are no parameters.
```go
package main

func main() {

}
```

## Building a Go Binary
Run one of the following commands in the base directory of the project.

```go
// Just a basic go binary will be built
go build

// The binary will not have debug information
go build -ldflags="-s -w" -trimpath

// Specify the executable name, otherwise will use module name
go build -o myexec
```

There is no functional difference between Go binaries, as in the running code is not changed at all between debug and release builds.


## Basic Variable Declaration
Variables in Go can be declared with
```go
var varName varType = varValue
```
or
```go
varName := varValue
```
Global variables cannot be declared with ```:=```, only locals.

## Advanced Variable Declaration
Variables declared without a value will be set to a default value like 0 or nil.
```go
var (
    id int
 	gopher string = "abc"
 	alive bool = true
)
```

## Variable Assignment
Variables are assigned with ```=```.

## Slices
Slices are a variant of an array or list from Python, C, etc.

Slices can be declared with a predefined data set,
as shown below. Providing ```*sliceSize``` is optional, as the compiler will automatically know the size.

```go
x := [*sliceSize]varType {var1, var2, var3}
```
You can declare a slice with make, and a predetermined size.

```go
x := make([]varType, size)
or
x := make([]varType)
```

You can even create subslices, very similar to Python lists.
```go
x := []int{1, 2, 3, 4, 5}
y := x[1:4]
```

Add variables to a slice like so.
```go
x := []int{1, 2, 3, 4, 5}
x = append(x, 20)
```

`WARNING: Slices can cause a panic if you attempt to access a member out of bounds.`

## Maps
Maps are just like dictionaries in Python with (K,V) pairs.

```go
m := map[keyType]valueType {
    key1: value1,
    key2: value2,
    ...
}
```

Making a map with make.
```go
m := make(map[keyType]valueType) 
```

You can also declare maps with a predetermined size.

```go
m := make(map[keyType]valueType, size)
```

Place a value into a map just like Python.
```go
customMap[key] = value
```

Grabbing a value from a map is simple.
```go
// "val" will contain the value corresponding to the key or a default value
// "exists" is a boolean indiciating if the key was present in the map
val, exists := customMap[key]
```

Deleting from a map is also simple.
```go
delete(customMap, key)
```

`WARNING: Maps can cause a panic if you attempt to read and write to a map at the same time.`

## Conditionals
Go comes with the standard `if` and `else` statements.

```go
// Basic if
if x > y {
    // Do something
}

// Basic if-else
if x > y {

} else if y > z {

}
```

## Loops
Go comes with a `for` loop, it doesn't come with any other flavor.

```go

// For loop
for index, value := range myslice {
    fmt.Println(value)
}

// For loop as a while loop
for x > y {

}

// For loop as a do while loop
for {

    if condition {
        break
    }
}
```


## Basic Function Declaration
Function parameters in Go are pass by value, just like C/C++.
```go
func myFunc(varName1 VarType1, varName2 VarType2) (returnType1, returnType2) {
    ...
    return var1, var2
}
```

If a function returns multiple values and you don't need one of the values just ignore it with ```_```.
```go
// I care about both return values
result1, result2 := myFunc(input1, input2)

// I only care about one return value
_, result2 := myFunc(input1, input2)
```

## Advanced Function Declaration
The return values are automatically made and are set to a default value like 0 or nil.
```go
func myFunc(varName1 VarType1, varName2 VarType2) (returnName1 returnType1,returnName2 returnType2) {
    returnName1 = returnValue1
    returnName2 = returnValue2
}
```

## Declaring Structs
Structs are declared with type and then the name of the struct, followed by fields inside of the struct.
```go
type structName struct {
    fieldName1 fieldType1
    fieldName2 fieldType2
    ...
}
```

Structs can be declared by assigning a value to each field by name, or by the order in which they are defined.
```go
x := structName{fieldName1: fieldValue1, fieldName2: fieldValue2, ...}

y := structName{fieldValue1, fieldValue2, ...}
```

## Methods
Methods are functions that are tied to a struct.
```go
type User struct {
    Name string
}

// Method
func (u User) Greet() string {
    return "Hello, " + u.Name
}

// Code being ran
mystruct := User{Name: "Go"}
fmt.Println(mystruct.Greet())
```

## Pointers
Pointers work very similarly to C/C++, you do not have to free any memory as the Go runtime will take care of that for you. Passing pointers can be more efficient than passing values. Set a pointer to nil if you don't want it pointing to anything.

```go
x := value

y := &x
y = nil
```

## Interfaces
Interfaces are types that contain a set of method signatures. Implementing methods with the same signatures as those listed in an interface, implicitly allows you to pass the struct as a interface.
```go
// Interface
type Speaker interface {
	Speak() string
}

// Struct
type Dog struct {
    Message string
}

// Method
func (dog Dog) Speak() string {
    fmt.Println(dog.Message)
}

// Function takes in the interface
func MakeItSpeak(s Speaker) {
    fmt.Println(s.Speak())
}

// Main code, passing dog as a Speaker
dog := Dog{Message: "Simplicity is key"}
MakeItSpeak(dog)
```

## Generics
These can work with any type of data.
```go
// T is a type parameter restricted by the 'any' constraint
func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}
```

You can make your own generic types.
```go
type Box[T any] struct {
    Content T
}

intBox := Box[int]{Content: 10}
strBox := Box[string]{Content: "Hello"}
```

## Type Constraints
You can limit the generics to specific types, allowing finer control. 
```go
type Number interface {
    ~int32 | ~float64 // The ~ allows types derived from int32 or float64
}

// This function will accept parameters whose types are derived from int32 or float64
func Sum[T Number](a, b T) T {
    return a + b
}
```


## Imports
Functionality can be provided via the standard library.

```go
import (
    "fmt"
    "math/rand"
    crand "crypto/rand" // alias, use with crand.X
    "io"
)
```

View all of the standard libary here: https://pkg.go.dev/std.


## File Operations
File operations are extraordinarily similar to Python.

You can create and open a new file.
```go
// This will truncate the file if it already exists.
f, err := os.Create(filename)
if err != nil {
    // error handling
}

// Continue processing........

// Success, done and closing the file. (You can actually defer this and put it right after the error check)
f.Close()
```

You can open an existing file.
```go
f, err := os.Open(filename)
if err != nil {
    // error handling
}

// Continue processing........

// Success, done and closing the file. (You can actually defer this and put it right after the error check)
f.Close()
```

For simple matters creating, writing and closing a file can be done in one line.
Or opening, reading, and closing a file can be done in one line.
```go
// Error handle as appropriate
err := os.WriteFile(filename, data, 0644)

// Error handle as appropriate
data, err := os.ReadFile(filename)
```

## Network Connections
Network connections are almost as simple as they are in Python.

Creating a listener and accepting a client:
```go
// Open the listener
listener, err := net.Listen("tcp", ":8080")

// Accept a client connection
conn, err := listener.Accept()

// Read from client
buf := make([]byte, 1024)
bytes_read, err := conn.Read(buf)

// Write to client
conn.Write([]byte("Hello from server"))
```

A client contacting a listener:
```go
// Attempt to connect
conn, err:= net.Dial("tcp", "localhost:8080")

// Write to server
conn.Write([]byte("Hello from client"))

// Read from server
buf := make([]byte, 1024)
bytes_read, err := conn.Read(buf)
```

Remember to error handle as appropriate.

## Packages and Modules
A package in Go, is simply all the code in a single directory. All code in the same directory/package are accessible to each other.

All `.go` files must begin with the line `package <directory_name>`, unless they are in the same directory as `go.mod`, which begin with `package main`.

A module is a collection of packages used to build a program or library. It will contain the ```go.mod``` file which specifies what version of Go is being used as well as other dependencies.

## Imports
This is a list of basic imports, you should be able to deduce what they might be used for based on their names. There are plenty more in the standard library and you can check the Go documentation/internet to find the dependencies you need.

```go
import (
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "log/slog"
    "net/http"
    "os"
    "path/filepath"
    "slices"
    "strings"
    "sync"
    "time"
)
```

## External Dependencies
You can import external modules and utilize their functionality.

Download external modules using `go get`.
```bash
go get github.com/google/uuid@latest
```

External modules will show up in `go.mod` and `go.sum` files. 

`go.mod` Will contain a list of external imports.

`go.sum` Will contain two lines with hashes for every external module, one will be the hash of the entire module and the other will be the hash of it's `go.mod`, this helps it quickly find dependencies that have changed.

Import them as so:

```go
import (
    "fmt"
    "github.com/google/uuid"
)
```

Here are a list of commands to manage external dependencies.

| Command  | Purpose |
|-------|----:|
| `go get <package>@<version>`   |   Add or upgrade a dependency.
| `go get <package>@none` | Remove a dependency.  |
| `go mod tidy` | Add missing modules and remove unused module requirements. |
| `go mod download` | Download dependencies to local module cache. |
| `go mod verify` | Validate cached modules against `go.sum`. |


## Identifiers (Public and Private)
Variables, contants, functions, structs, struct fields, interfaces and methods can be public or private, depending on the capitialization of the first character of the name. If the first character is uppercase, then it is public, if the character is lowercase then it is private.

## Nil
Instead of ```NULL``` or ```nullptr``` or ```None``` Go uses ```nil```.

## Defer
You can prepend any function/method call with the ```defer``` keyword. Once a line of code with ```defer``` is ran, the function will be called whenever the function returns. This allows simpler resource cleanup.

```go
// Something that may error
resource, err := DangerousFunc()
if err != nil {
    fmt.Println("Error occurred")
    return
}

// No error occurred, 
defer resource.close()

// Continue with dangerous functions and whenever an error is encountered, a return can happen and the resource will be cleaned up
```

# The Magic of Go

## Goroutines, Channels and Gophers


### Goroutines
Go supports asynchronous and multithreaded execution via an extraordinary simple mechanism, ```go```.

Simply by calling a function prepended with ```go``` will begin running it asynchronously and potentially on another thread.

```go
message := "Running in a go routine"

// Starting a go routine with an anonymous function (function without a name)
go func(msg string) {
    time.Sleep(2 * time.Second)
    fmt.Println(msg)
}(message)

// Continue execution

// Maybe start another gorountine
go myGopherFunction()

// Continue execution
```

To put it plainly, this is wizardry. Go will automatically handle the scheduling of this goroutine, it could be placed on a different thread or it might run on the current one that called it. Load balancing is automatically handled by Go and the OS scheduler.

Goroutines can spawn more goroutines, and thus you can have nested goroutines. But they aren't really nested, just spawned in a nested fashion.

The main thread of a Go program is not actually a main thread, it's the main goroutine.

```WARNING: Resources can be locked betwen goroutines and create a deadlock situation. If this occurs the runtime will recognize this and kill the process.```

```WARNING: If the main goroutine closes, all goroutines will close and the program will end.```


### Channels
Channels are how data can be sent between functions running in goroutines.

```go
// A basic channel, will block the sender until the receiver is ready to receive
chan1 := make(chan type)

// A buffered channel, will only block the sender if the capacity is hit
chan2 := make(chan type, capacity)

// Storing data into a channel
chan1 <- data1
chan2 <- data2

// Receiving data from a channel
chan1data := <- chan1
chan2data := <- chan2
```

If there is no data in a channel you are attempting to receive from then it will block.

You can attempt to receive data from many channels using a select block.
```go
select {
    case msg1 := <-chan1:
        fmt.Println("received 1", msg1)
    case msg2 := <-chan2:
        fmt.Println("received 2", msg2)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout")
}
```

You may use for loops on channels to grab all information inside of them.
```go
for i := range chan1 {
	fmt.Println(i)
}
```

Channels can be closed as well.
```go
// Closed channel
close(chan1)

// "ok" will be false and v will be a default value, unless there is still content that has not been processed
// All data will be taken from the channel before "ok" is false
v, ok := <-chan1
```
It is important to know that you cannot send data to a closed channel, you may only receive from one. The sender should always close the channel and you do not have to always close channels, if a channel is no longer used in the program, the garbage collector will rid of it.

Generally speaking, a channel should only be closed when it is signaling to receivers that it is no longer sending anymore data.

`WARNING: Go will panic if you attempt to send data to a closed channel.`