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
`cmd.Output()` is no hang on.

## How to stop example
```
umount -f /mnt/fuse
```
