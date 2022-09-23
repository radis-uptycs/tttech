package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("start")
	args := []string{"in_data/gobook.pdf"}
	cmd := exec.Command("exiftool", args...)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		fmt.Println(&errb)
	}
	fmt.Println(&outb)
	fmt.Println("end")
}
