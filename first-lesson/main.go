package main

import "fmt"

func main(){
	// fmt.Println("Hello World")

	// var isAuth bool
	// var name string
	// for{
	// var answ string
	// fmt.Println("Введите команду : /Регистрация  /Вход  /Выход")
	// fmt.Scan(&answ)




	// switch answ{
	// case "/Регистрация":
	// if isAuth{
	// 	fmt.Println("Вы уже зарегестрировались")
	// }else{
	// 	fmt.Println("Введите Имя")
	// 	fmt.Scan(&name)
	// 	isAuth = true 
	// }


	// case "/Вход":
	// if !isAuth {
	// 	fmt.Println("Вы не зарегистрированы")
	// } else {
	// 	fmt.Printf("Добро пожаловать %s\n", name)
	// }

   
  	// case "/Выход":
   	// 	fmt.Println("До связи")
   	// 	return
  	// default:
   	// 	fmt.Println("Че то не то")
  	// }
  
 	// }

	// var name string
	// var age int
	// var grade float32
	// fmt.Println("введите имя")
	// fmt.Scan(&name)

	// fmt.Println("Введите ваш возраст")
	// fmt.Scan(&age)

	// fmt.Println("Введите ваш средний бал")
	// fmt.Scan(&grade)

	// fmt.Printf("Ваше имя: %s/n Ваш возраст %d/n Ваш средний бал: %f",name,age,grade)

	// var a int 
	// var b int

	// fmt.Println("Введите число a")
	// fmt.Scan(&a)

	// fmt.Println("Введите число b")
	// fmt.Scan(&b)

	// var summ int = a + b
	// var diff int = a - b
	// var mul int = a * b
	// var divide int = a / b
	
	// fmt.Printf("Сумма чисел: %d/n, Разность чисел %d,Умножение чисел %d,Деление чисел %d",summ,diff,mul,divide)


	// var n int
	// fmt.Println("Введите число")
	// fmt.Scan(&n)
	// for i := 1; i < 11; i++ {
	// 	fmt.Println(n * i)
	// }

	// var n int
	// sum := 0
	// fmt.Println("Введите число")
	// fmt.Scan(&n)
	
	// for i := 1; i <= n; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	var A int
	var B int
	var choise string
	fmt.Println("Введите A")
	fmt.Scan(&A)

	fmt.Println("Введите B")
	fmt.Scan(&B)



	for{


		fmt.Println("Введите Оператор Или напишите /exit")
		fmt.Scan(&choise)

		switch choise{
		case "+":
			fmt.Println(A + B)
		case "-":
			fmt.Println(A - B)
		case "*":
			fmt.Println(A * B)
		case "/":
			fmt.Println(A / B)
		case "/exit":
			fmt.Println("До связи")
			return

	default:
		fmt.Println("Что то не так")

	}
}
}