package main // entry point of the package (main which compiler looks for to start (special one))
import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello world!")

	// int
	var num int = 100000
	fmt.Println("Integer", num)

	// float
	var floatNum float32 = 1.2222
	fmt.Println("Floating number", floatNum)

	// vairable declaring without type (just another syntax)
	variableWithoutSpecifyingType := "Special Varaible"
	fmt.Println("Without explicit type", variableWithoutSpecifyingType)

	// string
	var myName string = "Adharsh \n D"
	fmt.Println("String", myName)

	// string length in bytes
	var stringLength int = len(myName)
	fmt.Println("String Length in bytes", stringLength)

	// actual string length
	var actualStringLength = utf8.RuneCountInString(myName)
	fmt.Println("Actual String length", actualStringLength)

	// boolean
	var booleanVaraible bool = false
	fmt.Println("Boolean varaible..", booleanVaraible)

	// default values
	// int - 0 string - '' boolean - false

	// short hand varaible declaration
	shortHandVariable := "Rahul"
	fmt.Println("shortHandVariable...", shortHandVariable)

	// multiple varaible
	var1, var2 := 3, 1
	fmt.Println("1-st variable", var1, "2-nd varaible", var2)

	// constant vairable
	const constantVaraible float32 = 3.14
	fmt.Println("constand variable pie", constantVaraible) 
}
