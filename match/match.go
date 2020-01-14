package match

func Fib(n int) int64 {
     if n <= 1 {
     	retrun 1
     }
     return Fib(n-1) + Fib(n-2)
}

func Fact(n int) int64 {
     if n <= 1 {
     	retrun 1
     }
     return n * Fact(n-1)
}