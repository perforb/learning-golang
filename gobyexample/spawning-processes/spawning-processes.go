package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	check(err)
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	check(grepCmd.Start())
	_, err = grepIn.Write([]byte("hello grep\ngoodbye grep"))
	check(err)
	check(grepIn.Close())
	grepBytes, err := io.ReadAll(grepOut)
	check(err)
	err = grepCmd.Wait()
	check(err)
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	check(err)
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
