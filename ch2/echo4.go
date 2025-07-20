package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var n = flag.Bool("n", false, "omit end of line")
	var sep = flag.String("s", " ", "separator")
	//fmt.Println(os.Args)
	//fmt.Println(*n)
	//fmt.Println(*sep)
	flag.Parse()
	//fmt.Println(*n)
	//fmt.Println(*sep)
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
