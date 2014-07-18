viewerr
=======

golang error helper to view the error line and file.

## Installation

go get github.com/luosangnanka/viewerr

## Features
- When you debug just see the error info is not enough, but with this, you can see more info just like,
	
	[migo.go:74] ERROR error... 

or that in http applications,

	[migo.go:74] [127.0.0.1 /index/index] ERROR error...

## Demos
- You can just use that when:

	if err != nil {
		fmt.Println(viewerr.WrapError(err))
		os.Exit(2)
	}

When the package 'fmt' and 'os' is import indeed.