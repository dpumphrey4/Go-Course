# HelloWorld in Go
This "HelloWorld" is going to be diffrent than the ones you usually do. For this, I want you to learn at a base level how to use the VSCode Debug environment. So I want you to create a Go program that has a global variable and a local variable and then in the main function I want you to print these variables out.

Start by creating your module by running ```go mod init HelloWorld```, this should create the go.mod file in your local directory and then you may begin programming.

Your code should be similar to this, if you are copying this code, try changing the data types for your example.
``` 
package main

var y = 54

func main() {
	x := 5
	println("This is the value of x: ", x)
	println("This is the value of y: ", y)
}
```

