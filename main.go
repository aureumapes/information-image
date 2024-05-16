package main

import (
	"bufio"
	"information-image/internal"
	"os"
	"strings"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows"{
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	switch strings.Split(input, "\r\n")[0] {
	case "e":
		internal.ToImage("image.png", "text.txt")
		break
	case "d":
		internal.FromImage("image.png", "text2.txt")
	}
	}else {
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	switch strings.Split(input, "\n")[0] {
	case "e":
		internal.ToImage("image.png", "text.txt")
		break
	case "d":
		internal.FromImage("image.png", "text2.txt")
	}
	}
}
