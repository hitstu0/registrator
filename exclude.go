package main

import "fmt"

type Value interface {
    String() string
    Set(string) error
}

type arrayFlags []string

// Value ...
func (i *arrayFlags) String() string {
    return fmt.Sprint(*i)
}

func (i *arrayFlags) Set(value string) error {
    *i = append(*i, value)
    return nil
}
