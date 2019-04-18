package main

type testsuite struct {
  name string
  description string
  tests[] test
}

type test struct {
  name string
  description string

}

type testType string

const (
    equals testType = iota
)
