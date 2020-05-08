package golang_test

import "fmt"

// Hi returns a friendly greeting
func M_Print(name string) string {
	fmt.Println("Hello !!!")
   return fmt.Sprintf("MSG from golang_test::  %s", name)
}
