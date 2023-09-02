http://localhost:8080/debug/pprof/heap?debug=0

after receiving heap profile, go to dir with heap profile file and
```shell
# 'heap' - is name of heap profile file
go tool pprof -http localhost:3435 heap
```