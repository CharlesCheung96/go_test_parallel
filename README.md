# go_test_parallel
Verify the concurrency granularity of go test
```
#! /bin/bash

# Default demo Fail， pkg* Succ
go test -race --cover -count=1 ./demo ./pkg1 ./pkg2

# 限制 pkg 间并发为1 demo Fail， pkg* Succ
go test -p 1 -race --cover -count=1 ./demo ./pkg1 ./pkg2

# 限制 pkg 内并发为1 all Succ
go test -race --cover -count=1 -parallel=1  ./demo ./pkg1 ./pkg2

# 限制 pkg 内并发为1，结果同默认情况 demo Fail， pkg* Succ
go test -race --cover -count=1 -parallel=2 ./demo ./pkg1 ./pkg2

## chdemo Pass, ch1 & ch2 blocked --> pkg 间进程级别的隔离，pkg 内线程级别隔离
go test -race --cover -count=1 ./chdemo ./ch1 ./ch2
```
