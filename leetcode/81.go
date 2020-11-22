package leetcode

func search(nums []int, target int) bool {
	ss, se := 0, len(nums)-1
	for ss <= se && se < len(nums) && ss >= 0 {
		mid := (ss + se) / 2
		if nums[mid] == target {
			return true
		}
		pss, pse := ss, se
		i := mid + 1
		for ; i < len(nums); i++ {
			if nums[i] != nums[mid] {
				break
			}
		}
		if i == len(nums) {
			se = mid - 1
			continue
		}
		if mid-1 >= 0 && nums[ss] <= nums[mid-1] {
			if nums[ss] <= target && nums[mid-1] >= target {
				se = mid - 1
				continue
			} else {
				ss = mid + 1
				continue
			}
		}

		if mid+1 < len(nums) && nums[mid+1] <= nums[se] {
			if nums[mid+1] <= target && nums[se] >= target {
				ss = mid + 1
				continue
			} else {
				se = mid - 1
				continue
			}
		}
		if pss == ss && pse == se {
			break
		}
	}
	return false
}
