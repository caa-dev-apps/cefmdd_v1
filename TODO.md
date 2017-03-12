
= awaiting feedback 


# Insert Build tag

```
    http://www.golangbootcamp.com/book/tricks_and_tips

    $ go build -ldflags "-X main.Build a1064bc" example.go
    $ ./example
    Using build: a1064bc


    package main

    import "fmt"

    // compile passing -ldflags "-X main.Build <build sha1>"
    var Build string

    func main() {
        diag.Printf("Using build: %s\n", Build)
    }
```




        
# Misc Notes

```
    go test -run name-of-test-pattern  
    GOOS=linux GOARCH=386 go build
    set GOOS=linux&&set GOARCH=amd64&& go build -v .
    set GOOS=linux&&set GOARCH=amd64&& go build -v a1.go
    set GOOS=linux&&set GOARCH=386&& go build -v .

```

