# Tests
Tests in Go are written in files that end with `_test.go`. The functions inside of these test files must also begin with `Test`.

It is typical to see a file like `calculator.go`, and its corresponding test file `calulator_test.go` in the same directory/package.

However, this is not necessary, you can place tests in their own directory/packages.

Example of a test:
```go
import ("testing")

func TestSomething(t *testing.T) {
    // Your testing code goes here
    // ...

    
    if err != nil {
        t.Errorf("This was the thing that went wrong: %s", msg)
    }
    
}
```

Tests can be run with `go test` in the terminal, this will run all tests in your current directory, you can change the directory with `go test ./test_directory` or test all sub-directories/packages in the project with `go test ./...` in the base directory.

