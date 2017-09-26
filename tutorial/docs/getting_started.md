# Where to start
There are several ways to start with the Go programming language

- Download Go tools from [golang.org/dl](https://golang.org/dl/) [download and install the Go toolchain.  Configure your local environment and use your own editor to craft sweet Go code]
- Try it out at the Playground ([play.golang.org](https://play.golang.org)) [ for instant gratification, you can try this browser-based IDE that let you run your own code to test and share ideas]
- Take the Go Tour at [tour.golang.org](https://tour.golang.org) [ a step-by-step exploration of the features of the language with browser-based editor]

## Supported OSs / Architectures
The Go compiler can target 8 architectures: 

- arm
- arm64
- mips
- mips64
- ppc64
- s390
- x86
- x86_64

The Go compiler can target 9 operating systems:

- DragonFly BSD
- FreeBSD
- Linux
- NetBSD
- OpenBSD
- OS X
- Plan 9
- Solaris
- Windows

This means you can create your program from the comfort your Mac running OS X on a x86_64 and compile that code to run on Linux running on an ARM6.

## Installing Go
Go can be installed on your local machine in either of the following ways:
- From archive files - for Linux, Mac OS, FreeBSD, and Windows
- Installer - available for Mac OS and Windows
- Build from source 
Use URL [https://golang.org/doc/install](https://golang.org/doc/install) to follow instruction on your preferred method of installation for your target environment.

## Setup your workspace
Before you get introduced to your first Go program, let us take a moment to discuss your environment for local development.  After you install your Go toolchain, it is crucial that you properly setup your workspace.  It is an arbitrary directory which stores your Go source files, organized as packages, and built artifacts (such as object files and executable binaries) as shown in the following sample workspace:
```sh
$HOME/go
 +- src/
 |  +- github.com/vladimirvivien/go-tutorial
 |  |  +- hello/hello_world.go 
 |  |  +- greetings2/greet.go
 |  +- github.com/vladimirvivien/automi
 |  |  +- operators/batch/funcs.go
 |  |  +- stream/stream.go
 |  + ... (may contain many more source packages) ...
 |
 +- pkg/
 |  +- darwin_amd64
 |     +- github.com/vladimirvivien/go-tutorial/greetings2.a
 |     +- github.com/vladimirvivien/automi/operators/batch.a
 |     +- github.com/vladimirvivien/automi/stream.a
 |     +- ... (each package compiled as an object file) ...
 |
 +- bin/
    +- hello
    +- ... (compiled programs stored here) ...
```
## Setting GOPATH
By convention (and for the remainder of this discussion), the Go toolchain assumes the workspace is located at `$HOME/go`. If different, you will need to set env variable `$GOPATH` pointing to the workspace location.  You can use the `go env` command to verify the location of your `GOPATH` as shown below.
```sh
go env GOPATH
/Users/vvivien/DEV/go
```

## Editors and IDE
Because of its minimal syntax, working with Go can be done using your favorite text editor.  Some editors include features such as syntax higlights while others have deep integrations with the Go toolchain to provide an IDE-like experience.  Some of the more well-known editors with Go support include:
- Atom
- Emacs
- Sublime Text
- TextWrangler
- Vim/Neovim
- ... and more

If you feel more comfortable with working with a full IDE, you will be well-served there as well.  There are several IDE's providing full support for Go including the followings:
- IntelliJ Gogland - Go IDE
- Visual Studio Code - Support with vscode-go plugin
- GoClipse - Eclipse based Go IDE

See this [lengthy list of editors and IDEs with Go support](https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins) 

### Next
Read on to the next section to find out how to [create your first Go program](./first_program.md).
