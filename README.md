# GoCmdHangExample

## Issue
Go hang on `cmd.Output()`

## fuseExample
This fuseExample is copy from https://github.com/libfuse/libfuse/blob/master/example/passthrough.c
and add `sleep(60)` on `xmp_write()`
### How to run

## GoHangExample
This example have two go routine.
1. copy srcFile to fuse file system
2. loop to run shell systemctl --is-active smb for 10 times.

## GoHangExampleWithDeadline
An other example in Go. Use `context.WithDeadline()` to run shell. It also hangon `Run()`

## Duplicate Issue
1. run fuse file system
```
cd fuseExample/Debug
make 
mkdir /mnt/fuse
./fuseExample /mnt/fuse
```
2. run GoExample
```
cd GoHangExample
go run main.go
```
3. you will see
```
[root@archive GoHangExample]# go run main.go
2020/03/02 09:41:58 exec.Command loop 0
2020/03/02 09:41:58 systemctl is-active smb
2020/03/02 09:41:58 io.Copy start
2020/03/02 09:41:58 systemctl is-active smb done
2020/03/02 09:41:58 systemctl is-active smb: result
2020/03/02 09:41:58 active

2020/03/02 09:41:58 exec.Command loop 1
2020/03/02 09:41:58 systemctl is-active smb
2020/03/02 09:42:58 systemctl is-active smb done
2020/03/02 09:42:58 systemctl is-active smb: result
2020/03/02 09:42:58 active

2020/03/02 09:42:58 exec.Command loop 2
2020/03/02 09:42:58 systemctl is-active smb
2020/03/02 09:42:58 systemctl is-active smb done
2020/03/02 09:42:58 systemctl is-active smb: result
2020/03/02 09:42:58 active

...

2020/03/02 09:42:58 exec.Command loop 9
2020/03/02 09:42:58 systemctl is-active smb
2020/03/02 09:42:58 systemctl is-active smb done
2020/03/02 09:42:58 systemctl is-active smb: result
2020/03/02 09:42:58 active

```

We can saw `cmd.Output()` cost 60 seconds. It should not hang on by `io.copy()`.
```
2020/03/02 09:41:58 systemctl is-active smb
2020/03/02 09:42:58 systemctl is-active smb done
```
## What I want to see
`cmd.Output()` run without delay.

## Develop Environments
```
[root@archive GoCmdHangExample]# uname -a
Linux archive 3.10.0-1062.9.1.el7.x86_64 #1 SMP Fri Dec 6 15:49:49 UTC 2019 x86_64 x86_64 x86_64 GNU/Linux
[root@archive GoCmdHangExample]# rpm -q centos-release
centos-release-7-6.1810.2.el7.centos.x86_64
[root@archive GoCmdHangExample]# go version
go version go1.13.3 linux/amd64
[root@archive GoCmdHangExample]# go env
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/root/.cache/go-build"
GOENV="/root/.config/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/root/go"
GOPRIVATE=""
GOPROXY="direct"
GOROOT="/usr/lib/golang"
GOSUMDB="off"
GOTMPDIR=""
GOTOOLDIR="/usr/lib/golang/pkg/tool/linux_amd64"
GCCGO="gccgo"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build609385503=/tmp/go-build -gno-record-gcc-switches"
```

## CHangExample
C language also have this problem.
This example have two thread

1. copy srcFile to fuse file system
2. loop to run shell systemctl --is-active smb for 10 times.


## How to stop example
```
umount -f /mnt/fuse
```

## Q&A
1. When hang on, systemctl is-active smb not hang on. on shell.
2. The last Linux Kernel also have this problem.
`Linux localhost.localdomain 5.5.7-1.el7.elrepo.x86_64 #1 SMP Fri Feb 28 12:21:58 EST 2020 x86_64 x86_64 x86_64 GNU/Linux`
