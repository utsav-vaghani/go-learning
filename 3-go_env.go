package main

/**
	GO Enviornment Variables

	- Once GO is installed, there are some variable which we might have to understand.

	- GOROOT: it defines where your GO SDK is located in other words where your GO is installed.
			DO not change this untill you want to install different GO versions)

	- GOPATH: it defines your workspace, in other words all your code and all dependencies will be located here.
			it contains three main directories.
				src/ : contains GO source code
				pkg/ : contains compiled package code.( dependencies )
				bin/ : compiled executable programs. (contains all the executables which are installed with `go install`)

	- GOMOD: contains name of the module. concept `module` will be taken in separate chapter.

	- GOOS: it specifies the OS of the GO executable(not your machine's OS)

	- GOARCH: it specifies the OS of the GO executable(not your machine's architecture)
			for instance if go mac/amd64 can be installed on macOS M1, GOARCH will be amd64

	There are more enviornment variables but its out of scope, will be taken up in upcoming chapters.

	See you in next chapter: First program in go hello world.
**/
