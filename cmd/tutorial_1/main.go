package main

import (
	"fmt"

	"github.com/RamaByte/Traininam/cmd/tutorial_1/mathx"

	"unicode/utf8"
)

func main() {
	a, b := 2, 2
	fmt.Println(mathx.Add(a, b))

	fmt.Println("Labas")

	var intNum int = 32767
	intNum++
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println("Dalyba:", intNum1/intNum2)
	fmt.Println("Liekana:", intNum1%intNum2)

	var myString1 string = "Hello World"
	var myString2 string = "Hello \nWorld"
	var myString3 string = `Hello
							World`
	var myString4 string = "Hello" + " " + "World"
	fmt.Println(myString1)
	fmt.Println(myString2)
	fmt.Println(myString3)
	fmt.Println(myString4)

	fmt.Println(len("test"))
	fmt.Println(len("A"))
	fmt.Println(len("ù"))
	fmt.Println(utf8.RuneCountInString("ù"))

	var myRune rune = 'a'
	fmt.Println(myRune)
	fmt.Println()

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 int
	fmt.Println(intNum3)

	var myVar = "text"
	fmt.Println(myVar)
	myVar2 := "text"
	fmt.Println(myVar2)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "const value"
	fmt.Println(myConst)

	const pi float32 = 3.1415
	fmt.Println(pi)
}
