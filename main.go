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

	var names []byte

	for _, file := range files {
		info, e := file.Info()
		if e != nil {
			fmt.Println(e)
			return
		}
		if info.Size() == 0 {
			name := file.Name()
			names = append(names, name...)
			names = append(names, '\n')
		}
	}

	e := os.WriteFile("out.txt", names, 0644)
	if e != nil {
		fmt.Println(e)
		return
	}

	fmt.Printf("%s\n", names)
}
