package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}

	files, err := os.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		info, e := file.Info()
		if e != nil {
			fmt.Println(e)
			return
		}
		if info.Size() == 0 {
			name := file.Name()
			fmt.Println(name)
		}
	}
}
