http://localhost:8080/debug/pprof/

![profile_image.png](profile_image.png)

after receiving profile, go to dir with profile file and
```shell
# 'profile' - is name of profile file
go tool pprof -http localhost:3435 profile
```
