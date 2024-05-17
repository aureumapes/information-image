package main

import (
	"information-image/internal"
	"os"
)

func main() {
	switch os.Args[1] {
	case "e":
		internal.ToImage(os.Args[2], os.Args[3])
		break
	case "d":
		internal.FromImage(os.Args[2])
		break
	default:
		println("Unknown \"" + os.Args[1] + "\"")
	}
}
