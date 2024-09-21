package main

import (
	__array "leetcode/01_array"
	_2_linkedList "leetcode/02_linkedList"
	__hashmap "leetcode/03_hashmap"
	__strings "leetcode/04_strings"
	__stack_queue "leetcode/05_stack_queue"
	__backtracking "leetcode/07_backtracking"
	__greedy "leetcode/08_greedy"
	__dynamicProgramming "leetcode/09_dynamicProgramming"
	__monostack "leetcode/10_monostack"
	"leetcode/assessment"
)

func main() {
	// arrayTest()
	// linkedListTest()
	// hashmapTest()
	// stringTest()
	// stackQueueTest()
	// binaryTreeTest()
	// backtrackingTest()
	greedyTest()
	// dynamicProgrammingTest()
	// monostackTest()
	// graphTest()
	assessmentTest()
}

func assessmentTest() {
	// assessment.Echo()
	//assessment.FindMode()
	//assessment.ThousandSeparator()
	//assessment.GetNoZeroIntegers()
	//assessment.MinAreaRect()
	//assessment.CoinChange()
	//assessment.BitSetTest()
	assessment.Misc()
	assessment.AssignTasks()
	assessment.CountRangeSum()
	assessment.NumIslands()
	assessment.CountPrimes()
	assessment.Temp()
}

func graphTest() {

}

func monostackTest() {
	__monostack.Trap()
}

func dynamicProgrammingTest() {
	__dynamicProgramming.MinDistance()
	__dynamicProgramming.Dji()
	__dynamicProgramming.Trap()
}

func greedyTest() {
	__greedy.MaxProfit()
	__greedy.FindContentChildren()
}

func backtrackingTest() {
	__backtracking.SolveNQueens()
}

func binaryTreeTest() {

}

func stackQueueTest() {
	__stack_queue.MyQueueTest()
	__stack_queue.MyStackTest()
}

func stringTest() {
	__strings.ReverseString()
	__strings.ReverseStr()
	__strings.ReplaceNumber()
	__strings.ReverseWords()
	__strings.RightRotation()
	__strings.StrStr()
	__strings.RepeatedSubstringPattern()
}

func hashmapTest() {
	__hashmap.IsAnagram()
	__hashmap.Intersection()
	__hashmap.IsHappy()
	__hashmap.TwoSum()
	__hashmap.FourSumCount()
	__hashmap.CanConstruct()
	__hashmap.ThreeSum()
	__hashmap.FourSum()
}

func linkedListTest() {
	_2_linkedList.RemoveElements()
	_2_linkedList.LinkedList()
	_2_linkedList.ReverseList()
	_2_linkedList.SwapPairs()
	_2_linkedList.RemoveNthFromEnd()
	_2_linkedList.GetIntersectionNode()
	_2_linkedList.DetectCycle()
}

func arrayTest() {
	__array.Search()
	__array.RemoveElement()
	__array.SortedSquares()
	__array.MinSubArrayLen()
	__array.GenerateMatrix()
	__array.RangeSum()
}
