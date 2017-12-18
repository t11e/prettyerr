# prettyerr [![Build Status](https://travis-ci.org/t11e/prettyerr.svg)](https://travis-ci.org/t11e/prettyerr) [![GoDoc](https://godoc.org/github.com/t11e/prettyerr?status.svg)](http://godoc.org/github.com/t11e/prettyerr) [![Report card](https://goreportcard.com/badge/github.com/t11e/prettyerr)](https://goreportcard.com/report/github.com/t11e/prettyerr)

Adds pretty error printing to the errors created by 
[github.com/pkg/errors](https://github.com/pkg/errors).

# Usage

Minimal:

```go
fmt.Printf("something went wrong:\n%s\n",
    prettyerr.Format{err})
```

With options:

```go
fmt.Printf("something went wrong:\n%s\n",
    prettyerr.Format{
        Err: err,
        Flags: prettyerr.FlagTesting,
        Prefix: "TEST: ",
    })
```

Example output:

```
problem polishing widget: ran out of polish
Caused by: ran out of polish
    at github.com/t11e/prettyerr_test.ExampleFormat ($GOPATH/src/github.com/t11e/prettyerr/prettyerr_test.go)
    at testing.runExample ($GOROOT/src/testing/example.go)
    at testing.runExamples ($GOROOT/src/testing/example.go)
    at testing.(*M).Run ($GOROOT/src/testing/testing.go)
    at main.main (github.com/t11e/prettyerr/_test/_testmain.go)
```

[Read the package documentation for more information](https://godoc.org/github.com/t11e/prettyerr).

# Contributions

Clone this repository into your GOPATH and use [dep](https://github.com/golang/dep) to install its dependencies.

```shell
brew install dep
go get github.com/t11e/prettyerr
cd "$GOPATH"/src/github.com/t11e/prettyerr
dep ensure
```

You can then run the tests:

```shell
go test $(go list ./... | grep -v /vendor/)
```

# License

MIT. See [LICENSE](LICENSE) file.

