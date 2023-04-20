[![go](https://github.com/seipan/logfind/actions/workflows/go.yml/badge.svg)](https://github.com/seipan/logfind/actions/workflows/go.yml)
# logfind
This oss will find things like log.Println(), which you wrote for debugging but often forget to erase. 
## Install
```go
go install github.com/seipan/logfind/cmd/logfind
```

## Use
```go
go vet -vettool=`which logfind` pkgname
```
