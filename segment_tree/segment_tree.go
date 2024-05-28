package segment_tree

import (
	"container/list"
	"fmt"
	"log"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//root := &TreeNode{Val: 1}
	//root.Left = &TreeNode{Val: 2}
	//root.Right = &TreeNode{Val: 3}
	//root.Left.Left = &TreeNode{Val: 4}
	//root.Left.Right = &TreeNode{Val: 5}
	//l := list.New()
	//inOrder(root, l)
	//
	//hello2(hello)

	//log.Println(removeDuplicates([]int{1, 1, 1,1, 2, 2, 3}))
	// log.Println([]int{-1, 2, 1, -4}, 1)
	// log.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))

	// log.Println(maxSumContiguousArray([]int{4, 2, 1, 7, 8, 1, 2, 8, 1, 0}, 3))
	// log.Println(minWindow("ADOBECODEBANC", "ABC"))
	// log.Println(checkInclusion("hello", "ooolleoooleh"))
	// log.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	// log.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
	// log.Println(characterReplacement("ABAB", 2))

	numsArr := Constructor([]int{7, 2, 7, 2, 0})
	log.Println(numsArr.sgt)
	numsArr.Update(4, 6)
	numsArr.Update(0, 2)
	numsArr.Update(0, 9)
	numsArr.Update(3, 8)
	log.Println(numsArr.SumRange(0, 4))
}

type NumArray struct {
	nums []int
	sgt  []int
}

func Constructor(nums []int) NumArray {
	x := int(math.Log2(float64(len(nums)))) + 1
	maxLength := int(2*math.Pow(2, float64(x)) - 1)
	sgt := make([]int, maxLength)
	buildSegmentTree(nums, sgt, 0, len(nums)-1, 0)
	return NumArray{
		nums: nums,
		sgt:  sgt,
	}
}

func buildSegmentTree(nums, sgt []int, ss, se, si int) int {
	if ss == se {
		sgt[si] = nums[ss]
		return nums[ss]
	}

	if ss < se {
		mid := (ss + se) / 2
		sgt[si] = buildSegmentTree(nums, sgt, ss, mid, 2*si+1) + buildSegmentTree(nums, sgt, mid+1, se, 2*si+2)
		return sgt[si]
	}

	return 0
}

func (this *NumArray) Update(i int, val int) {
	diff := val - this.nums[i]
	this.nums[i] = val
	updateUtil(this.sgt, 0, len(this.nums)-1, 0, i, diff)
}

func updateUtil(sgt []int, ss, se, si, i, diff int) int {
	if i < ss || i > se {
		return sgt[si]
	}
	if ss == se && ss == i {
		sgt[si] += diff
	}
	if ss == se {
		return sgt[si]
	}
	if i >= ss && i <= se {
		mid := (ss + se) / 2
		sgt[si] = updateUtil(sgt, ss, mid, si*2+1, i, diff) + updateUtil(sgt, mid+1, se, si*2+2, i, diff)
		return sgt[si]
	}
	return 0
}

func (this *NumArray) SumRange(i int, j int) int {
	return getSumUtil(this.sgt, 0, len(this.nums)-1, i, j, 0)
}

func getSumUtil(sgt []int, ss, se, qs, qe, si int) int {
	if qs <= ss && qe >= se {
		return sgt[si]
	}
	if qs > se || qe < ss {
		return 0
	}
	mid := (ss + se) / 2
	return getSumUtil(sgt, ss, mid, qs, qe, 2*si+1) + getSumUtil(sgt, mid+1, se, qs, qe, 2*(si+1))
}

func findLength(A []int, B []int) int {
	return len(A)
}

// func characterReplacement(s string, k int) int {

// }

func minSubArrayLen(s int, nums []int) int {
	l, r, minLen := 0, 0, math.MaxInt32
	sum := 0
	for r < len(nums) {
		sum += nums[r]
		if sum < s {
			r++
			continue
		}
		for sum >= s {
			d := r - l + 1
			if d < minLen {
				minLen = d
			}
			sum -= nums[l]
			l++
		}
		r++
	}
	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

func checkInclusion(s1 string, s2 string) bool {
	m := make(map[rune]int)
	for _, c := range s1 {
		m[c]++
	}
	counter, l, r := len(s1), 0, 0
	for r < len(s2) {
		rr := rune(s2[r])
		rl := rune(s2[l])
		if m[rr] > 0 {
			counter--
		}
		m[rr]--
		if counter == 0 {
			return true
		}
		if r >= len(s1)-1 {
			m[rl]++
			if m[rl] > 0 {
				counter++
			}
			l++
		}
		r++
	}
	return false
}

func minWindow(s string, t string) string {
	l, r, bestL, bestR, minWindowSize := 0, 0, 0, 0, math.MaxInt32
	targetMap, tempMap, tempMap2 := make(map[rune]int), make(map[rune]int), make(map[rune]int)
	for _, c := range t {
		targetMap[c]++
		tempMap[c]++
		tempMap2[c]++
	}
	flag := true
	for r < len(s) && minWindowSize > len(t) {
		rr, rl := rune(s[r]), rune(s[l])
		if targetMap[rr] > 0 && flag {
			tempMap[rr]--
			tempMap2[rr]--
			if tempMap[rr] <= 0 {
				delete(tempMap, rr)
			}
		}
		if len(tempMap) == 0 {
			windowSize := r - l + 1
			if windowSize < minWindowSize {
				minWindowSize, bestL, bestR = windowSize, l, r
			}
			if targetMap[rl] > 0 {
				if tempMap2[rl] == 0 {
					tempMap[rl]++
				}
				tempMap2[rl]++
			}
			l++
			flag = false
		} else {
			r++
			flag = true
		}
	}
	return s[bestL : bestR+1]
}

func maxSumContiguousArray(nums []int, size int) int {
	curSum := 0
	for i := 0; i < size; i++ {
		curSum += nums[i]
	}
	maxSum := curSum
	for i := size; i < len(nums); i++ {
		curSum -= nums[i-size]
		curSum += nums[i]
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}

func searchRange(nums []int, target int) []int {

	return []int{}
}

func removeDuplicates(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 0
		}
		if count > 1 {
			nums = append(nums[:i], nums[i+1:]...)
			count--
			i--
		}
	}
	return len(nums)
}

func hello2(a func() string) {
	a()
}

func inOrder(root *TreeNode, stack *list.List) {
	for stack.Len() != 0 || root != nil {
		if root != nil {
			stack.PushFront(root)
			root = root.Left
		} else {
			e := stack.Front()
			stack.Remove(e)
			fmt.Println(e.Value.(*TreeNode).Val)
			if e.Value.(*TreeNode).Right != nil {
				stack.PushFront(e.Value.(*TreeNode).Right)
			}
		}

	}
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	i := 0
	last := len(postorder) - 1
	for _, val := range inorder {
		if val == postorder[last] {
			break
		}
		i++
	}
	if len(inorder) == 0 {
		return nil
	}
	return &TreeNode{
		Val:   postorder[last],
		Left:  buildTree(inorder[:i], postorder[:i]),
		Right: buildTree(inorder[i+1:], postorder[i:last]),
	}
}
