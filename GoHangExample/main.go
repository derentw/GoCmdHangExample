package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
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
			log.Printf("exec.Command loop %d", i)
			log.Println("systemctl is-active smb")
			cmd := exec.Command("systemctl", "is-active", "smb")
			result, err := cmd.Output()
			log.Println("systemctl is-active smb done")
			if err != nil {
				log.Println("systemctl is-active smb: got a error")
			}
			log.Println("systemctl is-active smb: result")
			log.Println(string(result))
		}
	}()

	wg.Wait()
}