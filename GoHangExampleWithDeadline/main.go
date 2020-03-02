package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		source, err := os.Open("srcFile")
		if err != nil {
			log.Println("os.Open(srcFile) fail")
			return
		}
		defer source.Close()

		destination, err := os.Create("/mnt/fuse/home/destFile")
		if err != nil {
			log.Println("os.Create(/mnt/fuse/home/destFile) fail")
			return
		}
		defer destination.Close()
		log.Println("io.Copy start")
		io.Copy(destination, source)
		log.Println("io.Copy done")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			log.Printf("loop %d", i)
			log.Println("systemctl is-active smb")
			//cmd := exec.Command("systemctl", "is-active", "smb")
			ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(5)*time.Second))
			defer cancel()
			log.Println("exec.CommandContext(ctx, systemctl, is-active, smb).Run()")

			if err := exec.CommandContext(ctx, "systemctl", "is-active", "smb").Run(); err != nil {
				fmt.Println("cmd: ", err)
				fmt.Println("ctx: ", ctx.Err())
			}
			log.Println("exec.CommandContext(ctx, systemctl, is-active, smb).Run() done")
			log.Println("systemctl is-active smb done")
		}
	}()

	wg.Wait()
}