package main

import "fmt"
func edit(s []int){
	s[0] = 999
}
func elem(s []int){
	s = append(s, 3)
	fmt.Println("Внутри :",s,len(s),cap(s))
}

func main() {
	// 1)
	// nums := make([]int, 0, 5)
	// nums = append(nums, 1, 2, 3, 4, 5)
	// fmt.Println(nums,len(nums),cap(nums))
	// 2)
	// nums := []int{1,2,3,4}
	// nums[0] = 10
	// nums[3] = 40
	// fmt.Println(nums)
	//3)
	// nums := []int{1,2,3,4}
	// nums2 := nums

	// nums[0] = 777;
	// fmt.Println(nums,"\n",nums2)
	//Он показывает в обоих одинаковое все потому что это срез,есди бы я делал массив то они были бы разные
	//4)
	// var s []int
	// fmt.Println(s,len(s),cap(s), s== nil)
	//5)
	// nums := []int{1,2,3}
	// nums = append(nums, 4)
	// fmt.Println(nums,len(nums),cap(nums))
	//6)
	// nums := make([]int,3,5)
	// nums[0] = 1
	// nums[1] = 2
	// nums[2] = 3

	// nums = append(nums, 4,5)
	// fmt.Println(nums,len(nums),cap(nums))
	//7)
	// nums := []int{1,2,3,4,5,6}
	// slice := nums[3:]
	// slice[0] = 111
	// fmt.Println(slice)
	// fmt.Println(nums)
	//8)
	// nums := []int{1,2,3,4,5,6}
	// slice := nums[3:4]
	// slice2 := slice
	// slice2[0] = 222
	
	// fmt.Println(slice)
	// fmt.Println(slice2)
	// fmt.Println(nums)
	//9)
	// nums := make([]int,3,5)
	// edit(nums)
	// fmt.Println(nums)
	//10
	// nums := make([]int, 2,2)
	// nums[0] = 1
	// nums[1] = 2
	// elem(nums)
	// fmt.Println("Снаружи : ",nums, len(nums),cap(nums))

}	

	
	
	
	
	