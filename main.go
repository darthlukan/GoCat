package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	if len(os.Args) > 1 {

		data, err := ioutil.ReadFile(os.Args[1])

		if err != nil {
			panic(err)
		} else {
			fmt.Printf("%v", string(data))
		}
	}
}
