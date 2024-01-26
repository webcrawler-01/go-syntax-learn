func sortColors(nums []int) {
	var one int = 0
	var two int = 0
	var three int = 0

	for k := 0; k < len(nums); k++ {
		if nums[k] == 0 {
			one++
		} else if nums[k] == 1 {
			two++
		} else {
			three++
		}
	}

	for n := 0; n < len(nums); n++ {
		if one > 0 {
			nums[n] = 0
			one--
		} else if two > 0 {
			nums[n] = 1
			two--
		} else {
			nums[n] = 2
			three--
		}
	}

	// Now you can proceed with the rest of your sorting logic.
}
