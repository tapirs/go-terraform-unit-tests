package main

import (
	"fmt"
	"os"
	//"github.com/tapirs/tfjson"
)

func main() {
  options(os.Args)
	tunit()
}

func options(args []string) {
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: tunit terraformplan.json")
		os.Exit(1)
	}
}

func tunit() {
	
}
// func tunit() {
// 	if len(os.Args) != 2 {
// 		fmt.Fprintln(os.Stderr, "usage: tfjson terraform.tfplan")
// 		os.Exit(1)
// 	}
//
// 	tfjson := tfjson.main{}
//
// 	j, err := &main.tfjson(os.Args[1])
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(1)
// 	}
//
// 	fmt.Println(j)
// }
