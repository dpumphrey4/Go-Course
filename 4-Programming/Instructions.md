# Introduction to Go Programming
In this section I want you to focus on playing with Go. Specifically the syntax, and data structures built into the standard library. Once you build something, use the debugger and try to figure out how it works on a lower level.

## Slices
1. Attempt to build slices, slices are just arrays in Go and there are multiple ways of building them.
2. Add data to a slice, and find out what actually happens when data is added. Use the debugger to help you.
3. Make subslices of other slices, they are all still slices but subslices have interesting properties. Attempt to modify these subslices and see what happens in the debugger.

## Maps
1. Learn how to make your own maps, these are generic maps just like Python's dictionaries.
2. Figure out how Go prevents crashes when grabbing a key from a map that doesn't exist. (There are 2 answers to this)

## Files
For this I don't want you to look at `Go.md` instead I want you to look at the `os` package using Intellisense. Import `os` and type `os.`, see what functions show up and play with them.

Once you are using some functions from `os`, control + left click or hover over the function and press F12 to enter the Go lang source code. Try to get an understanding of how file operations are happening.

You can also use the official Go lang documentation.

1. Learn how to create files, write to them, read from them and delete them.
2. Understand how to handle errors and use ```defer```.

## Packages
1. Put any code you'd like inside it's own directory and try to import it.
2. Experiment with public and private elements.

## Structures and Pointers
1. Make a struct, give it a bunch of members and then pass it to a function. Have the function change it and check the values from the callee function.
2. Figure out how to append to a slice or change a struct by passing it to another function.

## Utilize Interfaces
1. Find an interface or make your own. Once you choose your interface make a struct that then implements that interface. (Hint: You can use the error interface)

## Goroutines
1. Figure out how to make a Goroutine, play with scheduling of Goroutines by attempting to start some at the same time.
2. Pass data between Goroutines using channels, keep in mind you can share a channel across an infinite number of Goroutines. Experiment and have fun.

## 


