package main

func sum(nums ...float64) float64 {
	var sum float64
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func max(nums ...float64) float64 {
	var max float64
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func avg(nums ...float64) float64 {
	return sum(nums...) / float64(len(nums))
}
