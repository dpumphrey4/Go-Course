# HelloWorld in Go
This "HelloWorld" is going to be diffrent than the ones you usually do. For this, I want you to learn at a base level how to use the VSCode Debug environment. So I want you to create a Go program that has a global variable and a local variable and then in the main function I want you to print these variables out.

Start by creating your module by running ```go mod init HelloWorld```, this should create the go.mod file in your local directory and then you may begin programming.

Create a file with any name you want, but it must have the ```.go``` extension. It can be ```hello.go```.

Your code should be similar to this, if you are copying this code, try changing the data types for your example.

```go 
package main

var y = 54

func main() {
	x := 5
	println("This is the value of x: ", x)
	println("This is the value of y: ", y)
}
```

If a popup appears in the corner to install the debugger, attempt that, but if it fails run this command to install the debugger `go install github.com/go-delve/delve/cmd/dlv@latest`.

Use the debugger inside of VS Code. I want you to be able to stop execution and examine variables, maybe change the value of variables, step into and out of functions. Once you are confortable using the debugger, then you are done. Please make the code as complex or as simple as you wish.

To use the debugger the file with the `main` function must be open.

Note: If you are writing code and you save the code, saving it will run the Go automatic cleanup of your code. This may delete code, so try not to save if you have errors.

I may demo this if you haven't used the VSCode debugger before.