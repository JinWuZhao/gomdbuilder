# Gomdbuilder
A Go based DSL for building markdown

---


```go
package main

import (
	"io/ioutil"

	. "github.com/JinWuZhao/gomdbuilder"
)

func main() {
	doc := Doc(H1("Gomdbuilder"),
		P("A Go based DSL for building markdown"), 
		H1("Example"), 
		P("Example: ", Ln("example.go", "example/example.go")),
		P("Result:", Ln("example.md", "example/example.md")),
		)
	ioutil.WriteFile("example.md", []byte(doc), 0666)
}
```

# Example

Source: [example.go](example/example.go)  
Result: [example.md](example/example.md)

# License

BSD 3-Clause License

Copyright (c) 2018, jinwuzhao
All rights reserved.
