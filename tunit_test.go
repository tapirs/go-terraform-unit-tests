package main

import (
  "testing"
  "fmt"
)


func TestOptions(t *testing.T) {
  options([]string{"tunit", "test"})
}

func TestReadJson(t *testing.T) {
  result, _ := readJson("test/test.json")
  fmt.Println(result["aws_vpc.main"])

  var result2 map[string]interface{}

  result2, ok := result["aws_vpc.main"].(map[string]interface{})

  fmt.Println(result2["cidr_block"], ok)

  _, err := readJson("test/notest.json")
  if err == nil {
    t.Errorf("Should of been file not found, got %s", err)
  }
}
