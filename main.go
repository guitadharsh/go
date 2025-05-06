package main // entry point of the package (main which compiler looks for to start (special one))
import "fmt"
import "utf"

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
}
