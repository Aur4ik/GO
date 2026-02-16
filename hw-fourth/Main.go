package main

import "fmt"



func main() {
	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 1)
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }
	//2)
	// sum := 0;
	// for i := 0; i < len(arr); i++ {
	// 	sum += arr[i]
	// }
	// fmt.Println(sum)
	//3)
	// max := arr[0]
	// min := arr[0]

	// for i := 0; i < len(arr); i++ {
	// 	if arr[i] > max {
	// 		max = arr[i]
	// 	}
	// 	if arr[i] < min  {
	// 		min = arr[i]
	// 	}
	// }
	// fmt.Println("Максимальное: ", max)
	// fmt.Println("Минимальное: ", min)
	//4)
	// count := 0

	// for i := 0; i < len(arr); i++ {
	// 	if arr[i] % 2 == 0 {
	// 		count++
	// 	}
	// }
	// fmt.Println(count)
	//5)
	// for i := 0; i < len(arr); i++ {
	// 	arr[i] *= 2

	// }
	// fmt.Println(arr)

	//6)
	slice := []int{1,20,30,4,50,6,70,8,90,10}

	// slice = append(slice, 1,2,3)

	// fmt.Println(slice)
	//7)

	// fmt.Println(slice)

	// slice = append(slice[:1], slice[1 + 1] )
	// fmt.Println(slice)

	//8)
	// sum := 0
	// for _, v := range slice{
	// 	sum += v
	// }
	// avg := float32(sum)/float32(len(slice))
	
	// fmt.Println(sum)
	// fmt.Println(avg)
	//9)
	// names := []string{"Alice", "Bob", "Christina", "Dan"}
	// min := names[0]

	// for _, name := range names{
	// 	if len(name) < len(min) {
	// 		min = name
	// 	}
	// }
	// fmt.Println("Самое короткое имя : ", min)
	//10)
	// var newSlice []int
	// for _, n := range slice{
	// 	if n > 10{
	// 		newSlice = append(newSlice, n)
	// 	}
	// }
	// fmt.Println(newSlice)

}