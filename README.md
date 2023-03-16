# logfind
This oss will find things like log.Println(), which you wrote for debugging but often forget to erase. 
## Install
```go
import "github.com/seipan/logfind/cmd/logfind"
```

## Use
```go
go vet -vettool=`which logfind` pkgname
```
