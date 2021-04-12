package main

import "os/exec"
import "fmt"

func redShow(err string) {
	fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, err, 0x1B)
}

func main() {
	cmd := exec.Command("ssh-keygen", "-lf", "/tmp/497845385/id_rsa.pub")
	// cmd := exec.Command("ls", "-l", "/tmp/afasdfa")
	if err := cmd.Start(); err != nil {
		redShow(err.Error())
	} else {
		err = cmd.Wait()
		if err != nil {
			fmt.Println("--ok", err.Error())
		}
	}
}
