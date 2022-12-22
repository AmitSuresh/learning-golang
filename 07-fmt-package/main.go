package main

import (
	"fmt"
)

var a string
var b int
var c bool

func main() {
	fmt.Printf("1. %T\n",a)
	fmt.Println("2value of a=",a)
	fmt.Printf("3. \n%T\n",b)
	fmt.Println("4. \nvalue of b=",b)
	fmt.Printf("5. \n%T\n",c)
	fmt.Println("6. \nvalue of c=",c)
	fmt.Printf("7. \n%t\n",c)
	a = "Johnny Silverhand"
	fmt.Printf("8. \n%T\n",a)
	fmt.Println("9. \nvalue of a=",a)
	fmt.Print("9.1. \nvalue o\af a=",a)
	fmt.Printf("10. \nvalue of a=%#v\n",a)
	b = 7
	fmt.Printf("11. \n%T\n",b)
	fmt.Println("12. \nvalue of b=",b)
	fmt.Printf("13. \nvalue of b=%#v",b)
	fmt.Printf("\n14. below\n\n%#v\n%b\n%c\n%d\n%o\n%O\n%q\n%x\n%X\n%U\n%T\n\nabove",b,b,b,b,b,b,b,b,b,b,b)

	//string print
	s := fmt.Sprintf("\n14.1 below\n\n%#v\n%b\n%c\n%d\n%o\n%O\n%q\n%x\n%X\n%U\n%T\n\nabove",b,b,b,b,b,b,b,b,b,b,b)
	fmt.Printf("asdf%v",s)
	c = true
	fmt.Printf("15. \n%T\n",c)
	fmt.Printf("16. \n%t\n",c)
	fmt.Println("17. \nvalue of c=",c)
	fmt.Printf("18. \nvalue of c=%t",c)
	fmt.Printf("19. \nvalue of c=%#v",c)
}