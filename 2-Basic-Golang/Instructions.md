# Basic Golang
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

## Basic Variable Declaration
Variables in Go can be declared with
```go
var varName varType = varValue
```
or
```go
varName := varValue
```

## Basic Function Declaration
```go
func myFunc(varName1 VarType1, varName2 VarType2) (returnType1, returnType2) {
    ...
    return var1, var2
}
```

## Advanced Function Declaration
```go
func myFunc(varName1 VarType1, varName2 VarType2) (returnName1 returnType1,returnName2 returnType2) {
    returnName1 = returnValue1
    returnName2 = returnValue2
}
```


## Syntax
Go's syntax includes changes from C aimed at keeping code concise and readable. A combined declaration/initialization operator was introduced that allows the programmer to write ```i := 3``` or ```s := "Hello, world!"```, without specifying the types of variables used. This contrasts with C's ```int i = 3;``` and ```string s = "Hello, world!";``` (though since C23 type inference has been supported using auto, like C++). Go also removes the requirement to use parentheses in if statement conditions.

Semicolons still terminate statements; but are implicit when the end of a line occurs.

Methods may return multiple values, and returning a result, err pair is the conventional way a method indicates an error to its caller in Go. Go adds literal syntaxes for initializing struct parameters by name and for initializing maps and slices. As an alternative to C's three-statement for loop, Go's range expressions allow concise iteration over arrays, slices, strings, maps, and channels.


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





You may define your own types using structs.

#### Strings
Strings are immutable; built-in operators and keywords (rather than functions) provide concatenation, comparison, and UTF-8 encoding/decoding

#### Slices
For each type ```T``` and each non-negative integer constant ```n```, there is an array type denoted ```[n]T;``` arrays of differing lengths are thus of different types. Dynamic arrays are available as "slices", denoted ```[]T``` for some type ```T``` (compared to other languages like C/C++ and Java, where instead the arrays are denoted ```T[]```) These have a length and a capacity specifying when new memory needs to be allocated to expand the array. Several slices may share their underlying memory. If several slices share the same underlying memory and one slice suddenly wants to exceed the capacity, only that slice will get a new copy of the original slice with a higher capacity.

#### Pointers
Pointers are available for all types, and the pointer-to-```T``` type is denoted ```*T``` (similar to Rust; compare to other languages like C/C++ and C#, where pointers are denoted ```T*```). Address-taking and indirection use the ```&``` and ```*``` operators, as in C, or happen implicitly through the method call or attribute access syntax. There is no pointer arithmetic, except via the special ```unsafe.Pointer``` type in the standard library. Although possible, you should not be using pointer arithmetic in Go.

#### Maps
For a pair of types ```K```, ```V```, the type ```map[K]V``` is the type mapping type-```K``` keys to type-```V``` values, which can be thought of as equivalent to ```Map<K, V>``` in other languages. The Go Programming Language specification does not give any performance guarantees or implementation requirements for map types, though it is usually implemented as a hash table (equivalent to ```HashMap<K, V>``` in other languages). Hash tables are built into the language, with special syntax and built-in functions. 

#### Channels
```chan T``` is a channel that allows sending values of type ```T``` between concurrent Go processes.




BEN FOR THE FIRST THING, LETS GO OVER HOW SLICES CONTAIN THE SAME DATA UNTIL ONE GETS BIGGER THAN THE REST. Then go over using maps and function, pretty much do everything with local built in data types as the first assignment.

