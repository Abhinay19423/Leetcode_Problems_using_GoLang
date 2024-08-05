package main

import "fmt"

func twoSum(nums [3]int, target int) []int {

	a := make(map[int]int)
	for i:= 0; i<len(nums); i++{
		temp := target - nums[i]
		val, ok := a[temp]
		if ok{
			return []int{val, i}
		}
		a[nums[i]] = i
		// fmt.Println(a)
	}
	return []int{}
}

func main(){
	var a [3]int
	// fmt.Println("emp ", a)
	for i:= 0; i<len(a); i+= 1{
		fmt.Scanf("%v", &a[i])
	}

	// fmt.Println("enter the target ")
	var tar int
	fmt.Scanf("%d", &tar)
	fmt.Println(twoSum(a, tar))
}


/*

// brute force 
temp := target
	for i:= 0; i<len(nums) -1 ; i++{
		temp = nums[i]
		for j := i+1; j<len(nums); j++{
			if  temp+nums[j] == target{
				return []int{i, j}
			}
		}
	}
	return []int{}
*/
