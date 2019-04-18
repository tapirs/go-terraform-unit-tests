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
	tunit(os.Args[1], os.Args[2])
}

func options(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: tunit terraformplan.json")
		os.Exit(1)
	}
}

func tunit(jsonfile string, testfile string) (error) {
	planjson,err := readJson(jsonfile)
	if err != nil {
		return err
	}
	fmt.Println(planjson)

	tests, err := readTests(testfile)

	if err != nil {
		return err
	}
	fmt.Println(tests)

	return nil
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

func readTests(testfile string) (test, error) {
	return *new(test),  nil
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
