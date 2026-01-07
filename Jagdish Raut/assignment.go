package main

import(
	"fmt"
)

func reverse(str string) string{

	var  reverse_str string
	
	for i:=len(str)-1 ;i >= 0; i--{
		
		reverse_str += string(str[i])
	}

return reverse_str
	

}

func fibonacci (num int) int{
	if num == 1 || num == 0{
		return num
	}

	return fibonacci(num-1) + fibonacci(num-2)
}

func factorial (num int) int{
	if num <= 1{
		return 1
	}

	return num * factorial(num-1)
}



func main(){
	

	var str string
	var num int
	fmt.Print("Enter a string :")

	fmt.Scan(&str)

	fmt.Println("reverse of string is :", reverse(str))

	fmt.Println("Enter number: ")
	fmt.Scan(&num)
	fmt.Println("the",num ,"th fibonacci number is", fibonacci(num))

	if num < 15 {		
	fmt.Println("the factorial of ",num ," is", factorial(num))

	}else
	{
		fmt.Println("factorial is too high")
	}

	

}