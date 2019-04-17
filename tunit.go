package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	//"github.com/tapirs/tfjson"
)

func main() {
  options(os.Args)
	tunit(os.Args[1])
}

func options(args []string) {
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: tunit terraformplan.json")
		os.Exit(1)
	}
}

func tunit(jsonfile string) {
	readJson(jsonfile)
}

func readJson(jsonfile string) (map[string]interface{}, error) {
	jf, err := os.Open(jsonfile)
	if err != nil {
		return nil, err
	}
	defer jf.Close()

	byteValue, _ := ioutil.ReadAll(jf)

	var result map[string]interface{}
  json.Unmarshal([]byte(byteValue), &result)

	return result, nil
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
