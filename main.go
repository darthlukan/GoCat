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
			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("%v", string(data))
		}
	}
}
