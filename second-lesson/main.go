package main

import ("fmt")



func inRange(x int, min int, max int) bool{
	if x < max && x > min {
		return true
	}
	return false
}
func ageCategory(age int) string{
	switch {
	case age <= 12:
		return "child"
	case age > 12 && age <= 17:
		return "Teen"
	case age >=18 && age <=64:
		return "Adult"
	case age > 65:
		return "Senior"
	default:
		return "Some is wrong"
	}
}


func compareThree(a int, b int, c int) string {
	switch {
	case a == b && b == c:
		return "All equal"
	case a == b || a == c || b == c:
		return "Two equal"
	default:
		return "All different"
	}
}

func countDivisors(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}
func firsteven() string{
	var i int
	
	for {
		fmt.Println("Введите число")
		fmt.Scan(&i)

		if i % 2 == 0 {
			return "Четное !"
		}else{
			 fmt.Println("Продолжаем")
		}
	}
}


func averageUntilZero() float64 {
	var num int
	sum := 0
	count := 0

	for {
		fmt.Println("Введите число")
		fmt.Scan(&num)

		if num == 0 {
			break
		}

		sum += num
		count++
	}

	if count == 0 {
		return 0
	}

	return float64(sum) / float64(count)
}

func readPositive() int {
	var n int

	for {
		fmt.Println("Введите положительное число")
		fmt.Scan(&n)

		if n > 0 {
			return n
		}

		fmt.Println("Число должно быть больше 0, попробуйте ещё раз")
	}
}


func safeCompare(a int, b int) (string, bool) {
	if a == b {
		return "equal", true
	} else if a > b {
		return "a greater", true
	} else {
		return "b greater", true
	}
}


func main(){
	// var num int
	// var min int
	// var max int
	
	// fmt.Println("Enter Num")
	// fmt.Println("Enter Min")
	// fmt.Println("Enter Max")
	// fmt.Scan(&num)
	// fmt.Scan(&min)
	// fmt.Scan(&max)

	// fmt.Println(inRange(num,min,max))
	fmt.Println(firsteven())
}