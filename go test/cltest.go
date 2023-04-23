package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	args := os.Args[1:]
	var iArgs = []int{}
	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}

	max := 0
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}

	fmt.Println("max value: ", max)

}
