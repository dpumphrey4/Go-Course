# Protobufs
Protobufs are a custom serializing data structure invented by Google.

Protobufs are intended to replace JSON. JSON is human readable but profobufs are machine readable, this means that protobufs can store data more efficiently and have faster parsing times than JSON.

Protobufs come with their own compiler which convert protobuf designs stored in ```.proto``` files into usable code for a variety of languages like C/C++, Go, C#, Java, Python and others.

## Installing the Protobuf Compiler

Ubuntu/Debian
```bash
sudo apt install protobuf-compiler
```

Windows (Powershell)
```powershell
winget install Google.Protobuf
```

Once installed run this to confirm installation.
```bash
protoc --version
```

(Using go install enables you to download, install and build global terminal tools.)

Once you download the Protobuf compiler, then you'll have to install the protoc go generator.
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

Run this to confirm installation, if it cannot be found and you did install it, then there is likely an issue with the PATH.
```
protoc-gen-go --version
```

If it did not work, then there is likely a folder in your user directory named ```go``` and that will contain the binaries. Just add that to the PATH.

Use these to help find it, if GOBIN doesn't return anything then the binary will be inside GOPATH.
```
go env GOPATH
go env GOBIN
```

Add to your path by placing this line (or a variation of it) inside ```~/.bashrc``` or ```~/.zshrc```.
```
export PATH="$PATH:$(go env GOPATH)/bin"
```

Then run ```source ~/.bashrc``` or ```source ~/.zshrc```.

Now you can finally build a proto file into code usuable in Go.
```bash
protoc --go_out=. person.proto
```

When you place the file inside a module, the Go extension might give you errors stating that the packages you are importing are not available. This can be fixed by running the following code:
```
go get google.golang.org/protobuf
```

```go get``` adds online packages to your project. The code will be downloaded into a cache on your device. Go ensures that you are using the proper version of a package via it's version number and a corresponding hash. The external packages are kept track of in the ```go.mod``` and ```go.sum``` files.

Enough setup now moving onto the task at hand.

I want you to understand the power of protobufs, so I want you to create your own protobuf or protobufs and compare it to its JSON equivalent. Specifically the size difference between the two, you can compare the speed as well if you want.

Just add the following packages in your imports.
```
"google.golang.org/protobuf/encoding/protojson"
"google.golang.org/protobuf/proto"
```
