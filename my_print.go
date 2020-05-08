package golang_test

import "fmt"

// Hi returns a friendly greeting
func M_Print(name string) string {
   return fmt.Sprintf("MSG from golang_test::  %s", name)
}
