package main

import "os/exec"
import "fmt"
import "time"
import "os"
import "io/ioutil"

func redShow(err string) {
	fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, err, 0x1B)
}

func tssh() {
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
func tread() {
	f, err := os.Open("/tmp/497845385/id_rsa.pub")
	if err != nil {
		fmt.Println("open:", err.Error())
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read:", err.Error())
	} else {
		fmt.Println(string(b)[0:10])
	}
}
func main() {
	for idx := 0; idx != 10000; idx++ {
		time.Sleep(time.Second * 1)
		tssh()
		tread()
	}
}
