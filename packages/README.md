# Packages

`go mod init` creates `go.mod`

```
module somemodule

go 1.12
```

/go.mod
/somepackage/somefiles*
/otherpackage/otherfiles

```golang
import (
    "somemodule/somepackage"            // x
    somealias "somemodule/somepackage"  // y
    . "somemodule/somepackage"          // z
    _ "somemodule/somepackage"
)

func main() {
    x := somepackage.SomeStruct{}
    y := somealias.SomeStruct{}
    z := SomeStruct{}
}
```