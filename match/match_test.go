package main

import (
    "testing"
    "fmt"
    "github.com/Albertwzp/go-note/math"
)

func Test_Fib(t *testing.T) {
    var fibTests = []struct {
        in int
        expected int
    }{
        {1, 1},
        {2, 1},
        {3, 2},
        {4, 3},
        {5, 5},
        {6, 8},
        {7, 13},
        {8, 21},
        {9, 34},
    }
    for _, tt = range fibTests {
        out := Fib(tt.in)
        if out != tt.expected {
            t.Errorf("Fib(%d) = %d; expected %d", tt.in, out, tt.expected)
        }
    }
}

//func BenchmarkFib(b *testing.B) {
//}
