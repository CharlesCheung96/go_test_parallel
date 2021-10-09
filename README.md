# go_test_parallel

## Verify the concurrency granularity of go test
### Within Package
```bash
cd ./demo1

# chdemo Pass --> pkg 内线程级别隔离
go test -race -cover -count=1 ./chdemo
```
next, verify [testflag]():
```bash
# limit parallel to one, then chdemo blocked
go test -race -cover -count=1 -parallel=1 ./chdemo

# limit parallel to two, then chdemo pass
go test -race -cover -count=1 -parallel=2 ./chdemo

# limit parallel to two, cpu to one, then chdemo create 2 goroutine, but they are serially scheduled by single cpu 
go test -race -cover -count=1 -parallel=2 -cpu=1 ./chdemo
```
Here is an example about the `-cpu | -parallel` flag
```bash
# The following command runs on an 8-core machine
# only limit cpu
➜  demo1 git:(main) ✗ go test -p 1 -count=1 -cpu=1 ./cputest 
ok      testdemo/demo1/cputest  8.248s
➜  demo1 git:(main) ✗ go test -p 1 -count=1 -cpu=2 ./cputest
ok      testdemo/demo1/cputest  4.425s
➜  demo1 git:(main) ✗ go test -p 1 -count=1 -cpu=4 ./cputest
ok      testdemo/demo1/cputest  2.292s
➜  demo1 git:(main) ✗ go test -p 1 -count=1 -cpu=8 ./cputest
ok      testdemo/demo1/cputest  1.760s
➜  demo1 git:(main) ✗ go test -p 1 -count=1 -cpu=16 ./cputest
ok      testdemo/demo1/cputest  1.792s

# only limit parallel
➜  demo1 git:(main) ✗ go test -p 1 -count=1 -parallel=1 ./cputest
ok      testdemo/demo1/cputest  8.035s
➜  demo1 git:(main) ✗ go test -p 1 -count=1 -parallel=2 ./cputest
ok      testdemo/demo1/cputest  4.484s
➜  demo1 git:(main) ✗ go test -p 1 -count=1 -parallel=4 ./cputest
ok      testdemo/demo1/cputest  2.538s
➜  demo1 git:(main) ✗ go test -p 1 -count=1 -parallel=8 ./cputest
ok      testdemo/demo1/cputest  2.040s
➜  demo1 git:(main) ✗ go test -p 1 -count=1 -parallel=16 ./cputest
ok      testdemo/demo1/cputest  1.699s

# limit both
➜  demo1 git:(main) ✗ go test -p 1 -count=1 -parallel=2 -cpu=1 ./cputest
ok      testdemo/demo1/cputest  8.087s
➜  demo1 git:(main) ✗ go test -p 1 -count=1 -parallel=1 -cpu=2 ./cputest
ok      testdemo/demo1/cputest  8.212s
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
use -p parameter to limit the number of parallel programs
```bash
go test -p 1 -race -cover -count=1 ./ch1 ./ch2
ps aux | grep test | grep -v "[grep|race]"
## The result is similar to:
# charles          93591   0.0  0.0 409043920   3328 s016  S+    8:26PM   0:00.01 /var/folders/nr/y4gxj6zx7ds7tsqp6k13xr9h0000gn/T/go-build1711208026/b001/ch1.test -test.paniconexit0 -test.timeout=10m0s -test.count=1
```

