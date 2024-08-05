func twoSum(nums []int, target int) []int {
	a := make(map[int]int)
    for i:= 0; i<len(nums); i++{
        temp := target - nums[i]                
        if val, isTrue := a[temp]; isTrue{
            return []int{val, i}
        }
        a[nums[i]] = i
    }
    return []int{}

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