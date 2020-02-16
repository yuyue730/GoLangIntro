package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	pNum    = 4
	pIsTrue = false
	pStr    = "package"
)

func variableZeroValues() {
	var zero int
	var emptyStr string

	fmt.Printf("1. variableZeroValues zero=%d emptyStr=%q\n", zero, emptyStr)
}

func variableInitialValues() {
	var num1, num2 int = 4, 7
	var str string = "Hello World"
	fmt.Printf("2. variableInitialValues num1=%d num2=%d emptyStr=%q\n",
		num1, num2, str)
}

func variableTypeDeduction() {
	var num1, num2, isTrue, str = 3, 5, true, "hello"
	fmt.Printf("3. variableDeduction num1=%d num2=%d isTrue=%t emptyStr=%q\n",
		num1, num2, isTrue, str)
}

func variableShorter() {
	num1, num2, isTrue, str := 3, 5, true, "hello"
	num1 = 2
	str = "world"
	fmt.Printf("4. variableShorter num1=%d num2=%d isTrue=%t emptyStr=%q\n",
		num1, num2, isTrue, str)
}

func euler() {
	result := cmplx.Exp(1i*math.Pi) + 1
	fmt.Printf("6. Test whether euler's formula is correct. result=%.3f\n", result)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Printf("7. Type conversion result c=%d", calculateTriangle(a, b))
}

func calculateTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const (
		a, b     = 3, 4
		filename = "abc.txt"
	)
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Printf("8. Const filename=%q c=%d\n", filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
	)

	fmt.Printf("9. Enums [cpp, javascript, python, golang] = [%d, %d, %d, %d]\n",
		cpp, javascript, python, golang)
	fmt.Printf("[b, kb, mb, gb] = [%d, %d, %d, %d]\n", b, kb, mb, gb)
}

func main() {
	fmt.Println("Go language Basic Variables")
	variableZeroValues()
	variableInitialValues()
	variableTypeDeduction()
	variableShorter()
	fmt.Printf("5. Print package scope variable pNum=%d, pIsTrue=%t, pStr=%q\n",
		pNum, pIsTrue, pStr)
	euler()
	triangle()
	consts()
	enums()
}
