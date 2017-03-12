
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
