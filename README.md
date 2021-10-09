# go_test_parallel

## Verify the concurrency granularity of go test
### Within Package
```bash
cd ./demo1

# chdemo Pass --> pkg 内线程级别隔离
go test -race -cover -count=1 ./chdemo

# limit parallel to one, then chdemo blocked
go test -race -cover -count=1 -parallel=1 ./chdemo

# limit parallel to two, then chdemo pass
go test -race -cover -count=1 -parallel=2 ./chdemo
```

### Between Package
```bash
cd ./demo1

# ch1 & ch2 blocked --> pkg 间进程级别的隔离，
go test -race -cover -count=1 ./ch1 ./ch2
```
At the same time, run 
```bash
ps aux | grep test | grep -v "[grep|race]"
## The result is similar to:
# charles  92353   0.0  0.0 409044688   2544 s016  S+    8:12PM   0:00.02 /var/folders/nr/y4gxj6zx7ds7tsqp6k13xr9h0000gn/T/go-build3415734639/b063/ch2.test -test.paniconexit0 -test.timeout=10m0s -test.count=1
# charles  92352   0.0  0.0 409044688   2608 s016  S+    8:12PM   0:00.01 /var/folders/nr/y4gxj6zx7ds7tsqp6k13xr9h0000gn/T/go-build3415734639/b001/ch1.test -test.paniconexit0 -test.timeout=10m0s -test.count=1
```
