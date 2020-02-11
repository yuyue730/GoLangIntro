package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("Unsupported operator %q", op)
	}
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return q, r
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args %d %d\n", opName, a, b)
	return op(a, b)
}

func sum(nums ...int) int {
	result := 0
	for i := range nums {
		result += nums[i]
	}
	return result
}

func main() {
	fmt.Println("Go language Function grammar")

	fmt.Println("Deal with return error")
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}

	q, r := div(13, 3)
	fmt.Printf("13 / 3 = %d, 13 mod 3 = %d\n", q, r)

	fmt.Printf("3 ^ 4 = [%d]\n",
		apply(func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Printf("Sum 1 through 5 = [%d]\n", sum(1, 2, 3, 4, 5))
}
