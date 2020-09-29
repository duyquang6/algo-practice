package main

import (
	"container/heap"
	"container/list"
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"strconv"
	"sync"
	"time"
)

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func main() {
	// obj := Constructor()
	// param_1 := obj.Book(10, 20)
	// log.Println(param_1)
	// param_2 := obj.Book(50, 60)
	// log.Println(param_2)
	// param_3 := obj.Book(10, 40)
	// log.Println(param_3)
	// param_4 := obj.Book(5, 15)
	// log.Println(param_4)
	// param_5 := obj.Book(5, 10)
	// log.Println(param_5)
	// param_6 := obj.Book(25, 55)
	// log.Println(param_6)

	// log.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	// nums := []int{2, 3, 2}
	// log.Println(rob(nums))
	// s, p := "abab", "ab"
	// log.Println(findAnagrams(s, p))

	// log.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))

	// ["ate","eat","tea"],
	// ["nat","tan"],
	// ["bat"]

	// V := 6
	// adj := make([][]int, V)
	// addEdge(adj, 0, 1)
	// addEdge(adj, 1, 2)
	// // addEdge(adj, 2, 0)
	// addEdge(adj, 2, 3)
	// // addEdge(adj, 3, 4)
	// addEdge(adj, 3, 5)
	// addEdge(adj, 5, 4)
	// addEdge(adj, 0, 4)
	// log.Println(adj)
	// log.Println(isCyclicUndirectedDFS(adj))
	// nums := []int{2, 0, 2, 1, 1, 0}
	// log.Println(nums)
	// sortColors(nums)
	// log.Println(nums)
	// log.Println(nthUglyNumber(5, 2, 11, 13))

	// root := &TreeNode{Val: 5}
	// root.Left = &TreeNode{Val: 4}
	// root.Right = &TreeNode{Val: 8}
	// root.Left.Left = &TreeNode{Val: 11}
	// root.Left.Left.Left = &TreeNode{Val: 7}
	// root.Left.Left.Right = &TreeNode{Val: 2}
	// root.Right.Left = &TreeNode{Val: 13}
	// root.Right.Right = &TreeNode{Val: 4}
	// root.Right.Right.Left = &TreeNode{Val: 5}
	// root.Right.Right.Right = &TreeNode{Val: 1}
	// log.Println(hasPathSum(root, 26))
	// log.Println(isBalanced(root))
	// log.Println(pathSum(root, 22))

	// jobs := make(chan int, 100)
	// results := make(chan int, 100)

	// for w := 1; w <= 3; w++ {
	// 	go worker(w, jobs, results)
	// }

	// for j := 1; j <= 9; j++ {
	// 	jobs <- j
	// }
	// close(jobs)

	// for v := range results {
	// 	log.Println(v)
	// }
	// for a := 1; a <= 9; a++ {
	// 	soet := <-results
	// 	log.Println(soet)
	// }
	// campaignName := "toi"
	// if len(campaignName) > 5 {
	// 	campaignName = campaignName[:5] + "..."
	// }
	// log.Println(campaignName)
	// log.Println(merge([][]int{
	// 	[]int{1, 3},
	// 	[]int{2, 6},
	// 	[]int{8, 10},
	// 	[]int{15, 18},
	// }))
	// log.Println(largestSumOfAverages([]int{9, 1, 2, 5, 3, 9}, 4))
	// log.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
	// log.Println(getRow(3))
	// queue := make(chan int64)
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go func(ch chan<- int64) {
	// 	for i := 1; i < 100; i++ {
	// 		ch <- int64(i)
	// 	}
	// 	wg.Done()
	// }(queue)
	// go func(ch chan<- int64) {
	// 	for i := 1; i < 100; i++ {
	// 		ch <- int64(i)
	// 	}
	// 	wg.Done()
	// }(queue)

	// go func() {
	// 	wg.Wait()
	// 	close(queue)
	// }()

	// for elem := range queue {
	// 	wg.Add(1)
	// 	elem := elem
	// 	go func() {
	// 		fmt.Println(elem)
	// 		time.Sleep(time.Second)
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()
	// t := time.Now()
	// year, month, day := t.Date()
	// asiaBangkokLocation, _ := time.LoadLocation("Asia/Bangkok")
	// t2 := time.Date(year, month, day, 0, 0, 0, 0, asiaBangkokLocation)
	// log.Println(t2.Unix())

	// log.Println(time.Now().Unix())

	// log.Println(countElements([]int{1, 2, 3}))

	// log.Println(checkPossibility([]int{3, 4, 2, 3}))
	// log.Println(numEquivDominoPairs([][]int{[]int{1, 2}, []int{2, 1}, []int{3, 4}, []int{5, 6}}))
	// type Customer struct {
	// 	ID             int
	// 	Name, LastName string
	// 	Age            int
	// 	TaxNumber      string
	// }
	// // Create the input channel
	// ch := make(chan rxgo.Item)
	// // Data producer
	// go func(ch chan<- rxgo.Item) {
	// 	for i := 1; i < 100; i++ {
	// 		ch <- rxgo.Item{
	// 			V: Customer{
	// 				ID:       i,
	// 				Name:     "Quang " + strconv.Itoa(i),
	// 				LastName: "Nguyen",
	// 				Age:      i,
	// 			},
	// 		}
	// 	}
	// 	close(ch)
	// }(ch)

	// // Create an Observable
	// observable := rxgo.FromEventSource(ch)

	// getTaxNumber := func(c Customer) (string, error) {
	// 	time.Sleep(1 * time.Second)
	// 	return "QUANGQUADEPTRAI", nil
	// }

	// <-observable.
	// 	Filter(func(item interface{}) bool {
	// 		// filter
	// 		customer := item.(Customer)
	// 		return customer.Age > 18
	// 	}).
	// 	Map(func(_ context.Context, item interface{}) (interface{}, error) {
	// 		// enrich it
	// 		customer := item.(Customer)
	// 		taxNumber, err := getTaxNumber(customer)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		customer.TaxNumber = taxNumber
	// 		return customer, nil
	// 	},
	// 		rxgo.WithPool(5),
	// 		// Serialize the items emitted by their Customer.ID
	// 		rxgo.Serialize(func(item interface{}) int {
	// 			customer := item.(Customer)
	// 			return customer.ID
	// 		}), rxgo.WithBufferedChannel(1)).
	// 	ForEach(func(v interface{}) {
	// 		fmt.Printf("received: %v\n", v)
	// 	}, func(err error) {
	// 		fmt.Printf("error: %e\n", err)
	// 	}, func() {
	// 		fmt.Println("observable is closed")
	// 	})
	// log.Println(backspaceCompare("ab#c", "ad#c"))
	// root := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 		},
	// 	},
	// }
	// log.Println(splitListToParts(root, 5))
	// year, month, _ := time.Now().Date()
	// asiaBangkokLocation, _ := time.LoadLocation("Asia/Bangkok")
	// startMonthTime := time.Date(year, month, 1, 0, 0, 0, 0, asiaBangkokLocation)
	// endMonthTime := time.Date(year, month+1, 1, 0, 0, 0, 0, asiaBangkokLocation).Add(-1)
	// log.Println(startMonthTime.UTC(), endMonthTime.UTC())
	// nums := []int{0, 1, 1, 1, 0, 0, 1, 0, 0, 0}
	// nums := []int{0, 1, 1}
	// nums := []int{0, 1}
	// log.Println(findMaxLength(nums))

	// nums := []int{2, 3}
	// log.Println(productExceptSelf(nums))

	// log.Println(checkValidString("(*))"))
	// log.Println(checkValidString("()"))
	// log.Println(checkValidString("(*)"))
	// tvailableFrom := int64(1587319308)
	// availableFrom := time.Unix(tvailableFrom, 0)
	// log.Println(availableFrom)
	// log.Println(availableFrom.Hour()*60*60 + availableFrom.Minute()*60 + availableFrom.Second())
	//log.Println(leftMostColumnWithOne([]int{0, 1, 1, 1}))

	// root := &TreeNode{Val: 10}
	// root.Left = &TreeNode{Val: 5}
	// root.Left.Left = &TreeNode{Val: 3}
	// root.Left.Right = &TreeNode{Val: 7}
	// root.Right = &TreeNode{Val: 15}
	// root.Right.Right = &TreeNode{Val: 18}
	// log.Print(rangeSumBST(root, 7, 15))

	// leetcode.MinRemoveToMakeValid("a)b(c)d")
	// leetcode.MinRemoveToMakeValid("lee(t(c)o)de)")
	// leetcode.MinRemoveToMakeValid("a)b(c)d")
	// leetcode.EvalRPN([]string{"2", "1", "+", "3", "*"})
	// leetcode.EvalRPN([]string{"4", "13", "5", "/", "+"})
	// leetcode.EvalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	res := 0
	recur(root, L, R, &res)
	return res
}

func recur(root *TreeNode, L int, R int, curSum *int) {
	if root == nil {
		return
	}
	if root.Val <= R && root.Val >= L {
		*curSum += root.Val
	} else if root.Val < L {
		recur(root.Right, L, R, curSum)
		return
	} else if root.Val > R {
		recur(root.Left, L, R, curSum)
		return
	}
	recur(root.Left, L, R, curSum)
	recur(root.Right, L, R, curSum)
}

func leftMostColumnWithOne(nums []int) int {
	left, right := 0, len(nums)-1
	mid := (left + right) / 2
	val := 0
	resp := -1
	for left <= right {
		mid = (left + right) / 2
		val = nums[mid]
		if left == right {
			break
		}
		if val == 0 {
			left = mid + 1
		}
		if val == 1 {
			right = mid
		}
	}
	if val == 1 {
		resp = mid
	}
	return resp
}

func checkValidString(s string) bool {
	if s == "" {
		return true
	}
	stack := list.New()
	stack.PushFront(s[0])

	return true
}

func productExceptSelf(nums []int) []int {
	outArr := make([]int, len(nums))
	outArr[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 1; i-- {
		outArr[i] = nums[i] * outArr[i+1]
	}
	pLeft := 1
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			outArr[i] = pLeft
		} else {
			outArr[i] = pLeft * outArr[i+1]
		}
		pLeft *= nums[i]
	}
	return outArr
}

func findMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := map[int][]int{}
	count := 0
	best := 0
	for k, v := range nums {
		if v == 0 {
			count--
		}
		if v == 1 {
			count++
		}
		m[count] = append(m[count], k)
		if count == 0 {
			best = k + 1
		}
	}
	for _, v := range m {
		if len(v) > 1 {
			if v[len(v)-1]-v[0] > best {
				best = v[len(v)-1] - v[0]
			}
		}
	}
	return best
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	lenRoot := lengthRoot(root)
	tempDiv := int(lenRoot / k)
	tempRem := lenRoot % k
	p := root
	countPartition := 0
	target := 0
	var resp []*ListNode
	q := p
	for len(resp) < k {
		if countPartition == 0 {
			if tempRem > 0 {
				tempRem--
				target = tempDiv + 1
			} else {
				target = tempDiv
			}
		}
		countPartition++
		if countPartition == target {
			countPartition = 0
			resp = append(resp, q)
			ptr := p
			p = p.Next
			q = p
			ptr.Next = nil
			continue
		}
		if p != nil {
			p = p.Next
		} else {
			resp = append(resp, nil)
		}
	}
	return resp
}

func lengthRoot(root *ListNode) int {
	count := 0
	p := root
	for p != nil {
		count++
		p = p.Next
	}
	return count
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if pq[i].priority == pq[j].priority {
		return pq[i].value < pq[j].value
	}
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func topKFrequent(words []string, k int) []string {
	frequentMap := map[string]int{}
	// Preprocess data
	lenUniqueWords := 0
	for _, word := range words {
		if frequentMap[word] == 0 {
			lenUniqueWords++
		}
		frequentMap[word]++
	}

	pq := make(PriorityQueue, lenUniqueWords)
	i := 0
	for value, priority := range frequentMap {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}

	heap.Init(&pq)

	var resp []string
	// Take the k items out; they arrive in decreasing priority order.
	for j := 0; j < k; j++ {
		item := heap.Pop(&pq).(*Item)
		resp = append(resp, item.value)
	}

	return resp
}

func backspaceCompare(s1, s2 string) bool {
	return preprocessBackspace(s1) == preprocessBackspace(s2)
}

func preprocessBackspace(s2 string) string {
	for i := 0; i < len(s2); i++ {
		if s2[i] == '#' {
			if i == 0 {
				s2 = s2[i+1:]
			} else {
				s2 = s2[:i-1] + s2[i+1:]
			}

			i = -1
		}
	}
	return s2
}

// func numEquivDominoPairs(dominoes [][]int) int {
// 	if len(dominoes) == 0 {
// 		return 0
// 	}
// 	m := map[int]map[int]struct{}{}
// 	res := 0
// 	for i := 0; i < len(dominoes); i++ {
// 		for j := i + 1; j < len(dominoes); j++ {
// 			if (dominoes[i][0] == dominoes[j][0] && dominoes[i][1] == dominoes[j][1]) ||
// 				(dominoes[i][1] == dominoes[j][0] && dominoes[i][0] == dominoes[j][1]) {
// 				res++
// 			}
// 		}
// 	}

// 	return res
// }

func getMaximumGold(grid [][]int) int {
	maxSoFar := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 {
				m := map[string]bool{}
				m[getKey(i, j)] = true
				maxSoFar = int(math.Max(float64(maxSoFar), float64(dfsGetGold(grid, m, i, j, len(grid), len(grid[i]), grid[i][j]))))
			}
		}
	}
	return maxSoFar
}

func getKey(i, j int) string {
	return strconv.Itoa(i) + strconv.Itoa(j)
}

func dfsGetGold(grid [][]int, visited map[string]bool, i, j, row, col, sum int) int {
	toVisit := [][]int{}
	maxSoFar := sum
	// up
	if i > 0 {
		toVisit = append(toVisit, []int{i - 1, j})
	}

	//down
	if i < row-1 {
		toVisit = append(toVisit, []int{i + 1, j})
	}

	// left
	if j > 0 {
		toVisit = append(toVisit, []int{i, j - 1})
	}

	// right
	if j < col-1 {
		toVisit = append(toVisit, []int{i, j + 1})
	}

	for _, arr := range toVisit {
		if grid[arr[0]][arr[1]] > 0 {
			if !visited[getKey(arr[0], arr[1])] {
				curSum := grid[arr[0]][arr[1]] + sum
				visited[getKey(arr[0], arr[1])] = true
				maxSoFar = int(math.Max(float64(maxSoFar), float64(dfsGetGold(grid, visited, arr[0], arr[1], row, col, curSum))))
				visited[getKey(arr[0], arr[1])] = false
			}
		}
	}
	return maxSoFar
}

func checkPossibility(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	var count int
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			count++
		}
		if count > 1 {
			return false
		}
	}

	return true
}

func generate(numRows int) [][]int {
	dp := [][]int{}
	for i := 0; i < numRows; i++ {
		dp = append(dp, []int{})
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				dp[i] = append(dp[i], 1)
			} else {
				dp[i] = append(dp[i], dp[i-1][j]+dp[i-1][j-1])
			}
		}
	}
	return dp
}

func calcElem(i, j int) int {
	if i == 1 || j == i {
		return 1
	}
	return calcElem(i-1, j-1) + calcElem(i-1, j)
}

func countElements(arr []int) int {
	m := map[int]struct{}{}
	for _, v := range arr {
		m[v] = struct{}{}
	}
	count := 0
	for _, v := range arr {
		if _, ok := m[v+1]; ok {
			count++
		}
	}

	return count
}

type Store struct {
	StoreCode string
}

func getRow(rowIndex int) []int {
	res := []int{}
	m := map[string]int{}
	for i := 0; i <= rowIndex; i++ {
		res = append(res, calcPascal(rowIndex, i, m))
		log.Println(m)
	}
	return res
}

func calcPascal(i, j int, m map[string]int) int {
	if j == 0 || j == i {
		return 1
	}
	if m[strconv.Itoa(i)+strconv.Itoa(j)] == 0 {
		m[strconv.Itoa(i)+strconv.Itoa(j)] = calcPascal(i-1, j-1, m) + calcPascal(i-1, j, m)
	}
	return m[strconv.Itoa(i)+strconv.Itoa(j)]
}

func reverseList(head *ListNode) *ListNode {
	p := head
	for p.Next != nil {
		p = p.Next
	}
	helperReverse(head)
	return p
}

func helperReverse(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	p := helperReverse(head.Next)
	p.Next = head
	head.Next = nil
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p, q, k := head, head.Next, head.Next.Next
	swapPairs(k)
	q.Next = p
	p.Next = swapPairs(k)
	return head
}

func maxProfit2(prices []int) int {

	return prices[0]
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	for i < m {
		if nums1[i] < nums2[j] {
			i++
		}

		j++
	}
}

// func largestSumOfAverages(A []int, K int) float64 {
// 	var mean float64

// 	for i := len(A); i >= K; i-- {

// 	}
// 	return mean
// }

const (
	Shift = 6    // 2 ^ n = n of 64
	Mask  = 0x3f // N = 6, i.e. 2 ^ n - 1 = 63, i.e. 0x3f
)

func index(n int) int {
	return n >> Shift
}

//Use the value of a flag in the scene relative to the flag bit
func posVal(n int) uint64 {
	return 1 << uint(n&Mask)
}

// func merge(intervals [][]int) [][]int {
// 	resp := [][]int{}
// 	sort.Slice(intervals, func(i, j int) bool {
// 		if intervals[i][0] == intervals[j][0] {
// 			return intervals[i][1] > intervals[j][1]
// 		}
// 		return intervals[i][0] > intervals[j][0]
// 	})
// 	log.Println(intervals)
// 	i := len(intervals) - 1
// 	for i > 0 {

// 	}
// 	return resp
// }
func workJob() {
	var wg sync.WaitGroup
	workerJobs := make(chan int, 10)
	for i := 0; i < 1; i++ {
		go worker2(&wg, workerJobs)
	}
	wg.Add(10)
	for j := 1; j <= 10; j++ {
		workerJobs <- j
	}
	close(workerJobs)
	wg.Wait()
}

func worker2(wg *sync.WaitGroup, workerJobs <-chan int) {
	for j := range workerJobs {
		executingJob(wg, j)
	}
}

func executingJob(wg *sync.WaitGroup, j int) {
	defer wg.Done()
	log.Println(j)
	time.Sleep(1 * time.Second)
	wg.Add(1)
	go func() {
		log.Println("excuting go func in ", j)
		defer wg.Done()
		defer func() {
			log.Println("success ", j)
		}()
		time.Sleep(2 * time.Second)
	}()
}

func print(s string, i int) {
	fmt.Println(s, i)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

const charsetlist = "ABCDEFGHJKLMNPQRSTUVWXYZ123456789"
const charsetsingle = "ABCDEFGHJKLMNPQRSTUVWXYZ"

func SingleRandomCharacter(n int) (string, error) {
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = charsetsingle[b%byte(len(charsetsingle))]
	}
	return string(bytes), nil
}

func ListRandomCharacter(n int) (string, error) {
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = charsetlist[b%byte(len(charsetlist))]
	}
	return string(bytes), nil
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// func pathSum(root *TreeNode, sum int) [][]int {
// 	if root.Left == nil && root.Right == nil {
// 		if root.Val == sum {
// 			return [][]int{sum}
// 		}
// 	}
// 	resp := [][]int{}
// 	respL := pathSum(root.Left, sum-root.Val)
// 	respR := pathSum(root.Right, sum-root.Val)
// 	if len(respL) > 0 {
// 		resp = append(resp)
// 	}
// 	if len(respR) > 0 {

// 	}
// 	return [][]int{}
// }

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalanced(root.Left) && isBalanced(root.Right) && math.Abs(float64(dfsMaxHeight(root.Left, 1, math.MinInt16)-dfsMaxHeight(root.Right, 1, math.MinInt16))) < 2
}

func dfsMaxHeight(root *TreeNode, cur, max int) int {
	if root == nil {
		return getMaxInt(cur-1, max)
	}
	if root.Left == nil && root.Right == nil {
		return getMaxInt(cur, max)
	}
	return getMaxInt(dfsMaxHeight(root.Left, cur+1, max), dfsMaxHeight(root.Right, cur+1, max))
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	res := []int{}
	res2 := []int{}
	dfsGetLeaves(root1, &res)
	dfsGetLeaves(root2, &res2)
	if len(res) != len(res2) {
		return false
	}
	for i := range res {
		if res[i] != res2[i] {
			return false
		}
	}
	return true
}

func dfsGetLeaves(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*res = append(*res, root.Val)
	}
	dfsGetLeaves(root.Left, res)
	dfsGetLeaves(root.Right, res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func nthUglyNumber(n int, a int, b int, c int) int {
	dp := make([]int, n)
	dp[0] = 1
	indexA, indexB, indexC := 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i] = getMinInt(dp[indexA]*a, dp[indexB]*b, dp[indexC]*c)
		if dp[i] == dp[indexA]*a {
			indexA++
		}
		if dp[i] == dp[indexB]*b {
			indexB++
		}
		if dp[i] == dp[indexC]*c {
			indexC++
		}
	}
	return dp[n-1]
}

func getMinInt(args ...int) int {
	min := math.MaxInt32
	for _, x := range args {
		if x < min {
			min = x
		}
	}
	return min
}

func sortColors(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i] > nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}
}

func isCyclicUndirectedDFS(adj [][]int) bool {
	visited := make(map[int]bool)
	queue := list.New()
	queue.PushFront(0)
	for queue.Len() > 0 {
		e := queue.Back()
		node := e.Value.(int)
		count := 0
		for _, neighbor := range adj[node] {
			if !visited[neighbor] {
				queue.PushFront(neighbor)
			} else {
				count++
			}
		}
		if count > 1 {
			return true
		}
		visited[node] = true
		queue.Remove(e)
	}
	return false
}

func isCyclicUndirectedBFS(adj [][]int) bool {
	visited := make(map[int]bool)
	queue := list.New()
	queue.PushFront(0)
	for queue.Len() > 0 {
		e := queue.Back()
		node := e.Value.(int)
		count := 0
		for _, neighbor := range adj[node] {
			if !visited[neighbor] {
				queue.PushFront(neighbor)
			} else {
				count++
			}
		}
		if count > 1 {
			return true
		}
		visited[node] = true
		queue.Remove(e)
	}
	return false
}

func addEdge(adj [][]int, u, v int) {
	adj[u] = append(adj[u], v)
	adj[v] = append(adj[v], u)
}

func groupAnagrams(strs []string) [][]string {

	return [][]string{}
}

func findAnagrams(s string, p string) []int {
	res := []int{}
	m := make(map[rune]int)
	for _, c := range p {
		m[c]++
	}
	i, j, counter := 0, 0, len(p)
	for j < len(s) {
		rj := rune(s[j])
		ri := rune(s[i])
		size := j - i + 1
		if m[rj] > 0 {
			counter--
		}
		m[rj]--
		if counter == 0 {
			res = append(res, i)
		}
		if size == len(p) {
			m[ri]++
			if m[ri] > 0 {
				counter++
			}
			i++
		}
		j++
	}
	return res
}

func rob(nums []int) int {
	return getMaxInt(robUtil(nums[1:]), robUtil(nums[:len(nums)-1]))
}

func robUtil(nums []int) int {
	dp := make([]int, len(nums))
	maxSoFar := math.MinInt32
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = nums[0]
			continue
		}
		if i == 1 {
			dp[i] = getMaxInt(nums[0], nums[1])
			continue
		}
		dp[i] = getMaxInt(dp[i-1], dp[i-2]+nums[i])
		maxSoFar = getMaxInt(dp[i], maxSoFar)
	}
	if len(nums) > 0 {
		return dp[len(nums)-1]
	}
	return 0
}

func getMaxInt(args ...int) int {
	max := math.MinInt32
	for _, val := range args {
		if val > max {
			max = val
		}
	}
	return max
}

// func maxSumContigousSub(nums []int) int {
// 	maxSoFar := math.MinInt32
// 	curSum := 0
// 	i, j := 0, 0
// 	for j < len(nums) {
// 		curSum += nums[j]
// 		j++
// 		if curSum > maxSoFar {
// 			maxSoFar = curSum
// 		}
// 		curSum -= nums[i]
// 		i++
// 	}
// }

// Sliding windows version
func maxProfit(prices []int) int {
	j := 0
	minPrice := math.MaxInt32
	maxProf := 0
	for j < len(prices) {
		if j <= 0 {
			minPrice = prices[j]
		}
		profit := prices[j] - minPrice
		if profit < 0 {
			minPrice = prices[j]
		}
		if profit > maxProf {
			maxProf = profit
		}
		j++
	}
	return maxProf
}
